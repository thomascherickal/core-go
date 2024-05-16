// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"cogentcore.org/core/base/update"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/tree"
)

// ConfigItem represents a Widget configuration element,
// with one New function to create and configure a new child Widget,
// and an Update function to update the state of that child Widget.
// Contains a list of Config for its children, which is all sorted
// out as items are added.
type ConfigItem struct {
	// Path is the / delimited path to the element
	Path string

	// New returns a new Widget of the correct type for this element,
	// fully configured and ready for use.
	New func() Widget

	// Update updates the widget based on current state, so that it
	// propertly renders the correct information.
	Update func(w Widget)

	// Config for Children elements.
	Children Config
}

// Config is an ordered list of [ConfigItem]s,
// ordered in progressive hierarchical order
// so that parents are listed before any children.
// The Key of the map is the path of the element.
type Config []*ConfigItem

// Configure adds a new config item to the given [Config] for a widget at the
// given forward-slash-separated path with the given function for constructing
// the widget and the given optional function for updating the widget. If the
// given path is blank, it is automatically set to a unique name based on the
// filepath and line number of the calling function.
func Configure[T Widget](c *Config, path string, new func() T, update ...func(w T)) {
	if len(update) == 0 {
		c.Add(path, func() Widget { return new() }, nil)
	} else {
		u := update[0]
		c.Add(path, func() Widget { return new() }, func(w Widget) { u(w.(T)) })
	}
}

// ConfigButton adds a new config item to the given [Config] for
// a Button, with up to 2 optional functions:
// the first is called only during initial Config of a new Button;
// the second is called during Update.
func ConfigButton(c *Config, funcs ...func(w *Button)) {
	path := ConfigCallerPath(2)
	n := len(funcs)
	switch n {
	case 0:
		c.Add(path, func() Widget { return NewButton() }, nil)
	case 1:
		c.Add(path, func() Widget {
			w := NewButton()
			funcs[0](w)
			return w
		}, nil)
	case 2:
		c.Add(path, func() Widget {
			w := NewButton()
			funcs[0](w)
			return w
		}, func(w Widget) { funcs[1](w.(*Button)) })
	}
}

// ConfigSeparator adds a new config item to the given [Config] for
// a Separator, with no further configuration.
func ConfigSeparator(c *Config) {
	path := ConfigCallerPath(2)
	c.Add(path, func() Widget { return NewSeparator() }, nil)
}

// ConfigCallerPath returns the dir-filename of runtime.Caller(level),
// with all / . replaced to -, which is suitable as a unique name
// for a ConfigItem path.  This is used with level 1 by default for
// blank names in [Config.Add], but wrappers around that may need to
// generate these paths with a relevant level if they wrap the Configure
// call.
func ConfigCallerPath(level int) string {
	_, file, line, _ := runtime.Caller(level)
	dir, fn := filepath.Split(file)
	d1 := ""
	if len(dir) > 1 {
		_, d1 = filepath.Split(dir[:len(dir)-1])
	}
	path := fn + "-" + d1
	// need to get rid of slashes and dots for path name
	path = strings.ReplaceAll(strings.ReplaceAll(path, "/", "-"), ".", "-") + "-" + strconv.Itoa(line)
	return path
}

// Add adds a new config item for a widget at the given forward-slash-separated
// path with the given function for constructing the widget and the given function
// for updating the widget. This should be called on the root level Config
// list. Any items with nested paths are added to Children lists as
// appropriate. If the given path is blank, it is automatically set to
// a unique name based on the filepath and line number of the calling function.
// Consider using the [Configure] global generic function for
// better type safety and increased convenience.
func (c *Config) Add(path string, new func() Widget, update func(w Widget)) {
	if path == "" {
		path = ConfigCallerPath(1)
	}
	itm := &ConfigItem{Path: path, New: new, Update: update}
	plist := strings.Split(path, "/")
	if len(plist) == 1 {
		*c = append(*c, itm)
		return
	}
	next := c.FindMakeChild(plist[0])
	next.AddSubItem(plist[1:], itm)
}

// ChildPath returns this item's path + "/" + child
func (c *ConfigItem) ChildPath(child string) string {
	return c.Path + "/" + child
}

// ItemName returns this item's name from its path,
// as the last element in the path.
func (c *ConfigItem) ItemName() string {
	pi := strings.LastIndex(c.Path, "/")
	if pi < 0 {
		return c.Path
	}
	return c.Path[pi+1:]
}

// AddSubItem adds given sub item to this config item, based on the
// list of path elements, where the first path element should be
// immediate Children of this item, etc.
func (c *ConfigItem) AddSubItem(path []string, itm *ConfigItem) {
	fpath := c.ChildPath(path[0])
	child := c.Children.FindMakeChild(fpath)
	if len(path) == 1 {
		*child = *itm
		return
	}
	child.AddSubItem(path[1:], itm)
}

// FindChild finds item with given path string within this list.
// Does not search within any Children of items here.
// Returns nil if not found.
func (c *Config) FindChild(path string) *ConfigItem {
	for _, itm := range *c {
		if itm.Path == path {
			return itm
		}
	}
	return nil
}

// FindMakeChild finds item with given path element within this list.
// Does not search within any Children of items here.
// Makes a new item if not found.
func (c *Config) FindMakeChild(path string) *ConfigItem {
	for _, itm := range *c {
		if itm.Path == path {
			return itm
		}
	}
	itm := &ConfigItem{Path: path}
	*c = append(*c, itm)
	return itm
}

// String returns a newline separated list of paths for all the items,
// including the Children of items.
func (c *Config) String() string {
	str := ""
	for _, itm := range *c {
		str += itm.Path + "\n"
		str += itm.Children.String()
	}
	return str
}

// SplitParts splits out a config item with sub-name "parts" from remainder
// returning both (each of which could be nil).  parpath contains any parent
// path to add to the path
func (c *Config) SplitParts(parpath string) (parts *ConfigItem, children Config) {
	partnm := "parts"
	if parpath != "" {
		partnm = parpath + "/" + partnm
	}
	for _, itm := range *c {
		if itm.Path == partnm {
			parts = itm
		} else {
			children = append(children, itm)
		}
	}
	return
}

// ConfigWidget runs the Config on the given widget, ensuring that
// the widget has the specified parts and direct Children.
// The given parent path is used for recursion and should be blank
// when calling the function externally.
func (c *Config) ConfigWidget(w Widget, parentPath string) {
	wb := w.AsWidget()
	parts, children := c.SplitParts(parentPath)
	if parts != nil {
		if wb.Parts == nil {
			wparts := parts.New()
			wparts.SetName("parts")
			wb.Parts = wparts.(*Frame)
			tree.SetParent(wb.Parts, wb)
			parts.Children.ConfigWidget(wparts, parts.ChildPath("parts"))
		}
	}
	n := len(children)
	if n > 0 {
		wb.Kids, _ = update.Update(wb.Kids, n,
			func(i int) string { return children[i].ItemName() },
			func(name string, i int) tree.Node {
				child := children[i]
				if child.New == nil {
					fmt.Println("core.Config child.New is nil:", child.ItemName())
					return nil
				}
				cw := child.New()
				cw.SetName(name)
				// fmt.Println(name, cw, wb)
				tree.SetParent(cw, wb)
				if child.Update != nil { // do initial setting in case children might reference
					child.Update(cw)
				}
				return cw
			})
		for i, child := range children { // always config children even if not new
			if len(child.Children) > 0 {
				cw := wb.Child(i).(Widget)
				child.Children.ConfigWidget(cw, child.Path)
			}
		}
	}
	if parentPath == "" { // top level
		c.UpdateWidget(w, parentPath) // this is recursive
	}
}

// UpdateWidget runs the [ConfigItem.Update] functions on the given widget,
// and recursively on all of its children as specified in the Config.
// It is called at the end of [Config.ConfigWidget].
// The given parent path is used for recursion and should be blank
// when calling the function externally.
func (c *Config) UpdateWidget(w Widget, parentPath string) {
	wb := w.AsWidget()
	parts, children := c.SplitParts(parentPath)
	if parts != nil {
		parts.Children.UpdateWidget(wb.Parts, parts.ChildPath("parts"))
	}
	n := len(children)
	if n == 0 {
		return
	}
	for i, child := range children {
		cw := wb.Child(i).(Widget)
		if child.Update != nil {
			child.Update(cw)
		}
		if len(child.Children) > 0 {
			child.Children.UpdateWidget(cw, child.Path)
		}
	}
}

// NewParts returns a new [Frame] that serves as the internal parts
// of a widget, which typically contain content that the widget automatically
// manages through its [Widget.Config] method.
func NewParts() *Frame {
	w := NewFrame()
	w.SetName("parts")
	w.SetFlag(true, tree.Field)
	w.Style(func(s *styles.Style) {
		s.Grow.Set(1, 1)
	})
	return w
}

// ConfigWidget is the base implementation of [Widget.ConfigWidget] that
// configures the widget by doing steps that apply to all widgets and then
// calling [Widget.Config] for widget-specific configuration steps.
func (wb *WidgetBase) ConfigWidget() {
	if wb.ValueUpdate != nil {
		wb.ValueUpdate()
	}
	c := Config{}
	wb.This().(Widget).Config(&c)
	if len(c) > 0 {
		c.ConfigWidget(wb.This().(Widget), "")
	}
}

// Config is the interface method called by [Widget.ConfigWidget] that
// should be defined for each [Widget] type, which actually does
// the configuration work.
func (wb *WidgetBase) Config(c *Config) {
	// this must be defined for each widget type
}

// ConfigTree calls [Widget.ConfigWidget] on every Widget in the tree from me.
func (wb *WidgetBase) ConfigTree() {
	if wb.This() == nil {
		return
	}
	// pr := profile.Start(wb.This().NodeType().ShortName())
	wb.WidgetWalkDown(func(wi Widget, wb *WidgetBase) bool {
		wi.ConfigWidget()
		return tree.Continue
	})
	// pr.End()
}

// Update does a general purpose update of the widget and everything
// below it by reconfiguring it, applying its styles, and indicating
// that it needs a new layout pass. It is the main way that end users
// should update widgets, and it should be called after making any
// changes to the core properties of a widget (for example, the text
// of [Text], the icon of a [Button], or the slice of a table view).
//
// If you are calling this in a separate goroutine outside of the main
// configuration, rendering, and event handling structure, you need to
// call [WidgetBase.AsyncLock] and [WidgetBase.AsyncUnlock] before and
// after this, respectively.
func (wb *WidgetBase) Update() { //types:add
	if wb == nil || wb.This() == nil {
		return
	}
	if DebugSettings.UpdateTrace {
		fmt.Println("\tDebugSettings.UpdateTrace Update:", wb)
	}
	wb.WidgetWalkDown(func(wi Widget, wb *WidgetBase) bool {
		wi.ConfigWidget()
		wi.ApplyStyle()
		return tree.Continue
	})
	wb.NeedsLayout()
}

//////////////////////////////////////////////////////////////////////////////
// 	ConfigFuncs

// ConfigFuncs is a stack of config functions, which take a Config
// and add to it.
type ConfigFuncs []func(c *Config)

// Add adds the given function for configuring a toolbar
func (cf *ConfigFuncs) Add(fun ...func(c *Config)) *ConfigFuncs {
	*cf = append(*cf, fun...)
	return cf
}

// Call calls all the functions for configuring given toolbar
func (cf *ConfigFuncs) Call(c *Config) {
	for _, fun := range *cf {
		fun(c)
	}
}

// IsEmpty returns true if there are no functions added
func (cf *ConfigFuncs) IsEmpty() bool {
	return len(*cf) == 0
}
