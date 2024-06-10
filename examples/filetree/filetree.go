// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate core generate

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/filetree"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/keymap"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/texteditor"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/views"
)

// FileBrowse is a simple file browser / viewer / editor with a file tree and
// one or more editor windows.  It is based on an early version of the Gide
// IDE framework, and remains simple to test / demo the file tree component.
type FileBrowse struct {
	core.Frame

	// root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjectFilename
	ProjectRoot core.Filename

	// filename of the currently active texteditor
	ActiveFilename core.Filename

	// has the root changed?  we receive update signals from root for changes
	Changed bool `json:"-"`

	// all the files in the project directory and subdirectories
	Files *filetree.Tree

	// number of texteditors available for editing files (default 2) -- configurable with n-text-views property
	NTextEditors int `xml:"n-text-views"`

	// index of the currently active texteditor -- new files will be viewed in other views if available
	ActiveTextEditorIndex int `json:"-"`
}

func (fb *FileBrowse) Defaults() {
	fb.NTextEditors = 2
}

// todo: rewrite with direct config, as a better example

func (fb *FileBrowse) Init() {
	fb.Defaults()
	fb.Styler(func(s *styles.Style) {
		s.Direction = styles.Column
		s.Grow.Set(1, 1)
		s.Margin.Set(units.Dp(8))
	})
	fb.OnWidgetAdded(func(w core.Widget) { // TODO(config)
		switch w.PathFrom(fb) {
		case "title":
			title := w.(*core.Text)
			title.Type = core.TextHeadlineSmall
			w.Styler(func(s *styles.Style) {
				s.Justify.Content = styles.Center
			})
		}
		if w.Parent().PathFrom(fb) == "splits" {
			if w.AsTree().IndexInParent() == 0 {
				w.Styler(func(s *styles.Style) {
					s.Grow.Set(1, 1)
				})
			} else {
				w.Styler(func(s *styles.Style) {
					s.Grow.Set(1, 1)
					s.Min.X.Ch(20)
					s.Min.Y.Ch(10)
					s.Text.WhiteSpace = styles.WhiteSpacePreWrap
					s.Text.TabSize = 4
				})
			}
		}
	})
}

// UpdateFiles updates the list of files saved in project
func (fb *FileBrowse) UpdateFiles() { //types:add
	if fb.Files == nil {
		return
	}
	fb.Files.UpdateAll()
}

// IsEmpty returns true if given FileBrowse project is empty -- has not been set to a valid path
func (fb *FileBrowse) IsEmpty() bool {
	return fb.ProjectRoot == ""
}

// OpenPath opens a new browser viewer at given path, which can either be a
// specific file or a directory containing multiple files of interest -- opens
// in current FileBrowse object if it is empty, or otherwise opens a new
// window.
func (fb *FileBrowse) OpenPath(path core.Filename) { //types:add
	if !fb.IsEmpty() {
		NewFileBrowser(string(path))
		return
	}
	fb.Defaults()
	root, pnm, fnm, ok := ProjectPathParse(string(path))
	if !ok {
		return
	}
	fb.ProjectRoot = core.Filename(root)
	fb.SetName(pnm)
	fb.UpdateProject()
	fb.Files.OpenPath(root)
	// win := fb.ParentRenderWindow()
	// if win != nil {
	// 	winm := "browser-" + pnm
	// 	win.SetName(winm)
	// 	win.SetTitle(winm)
	// }
	if fnm != "" {
		fb.ViewFile(fnm)
	}
	fb.UpdateFiles()
}

// UpdateProject does full update to current proj
func (fb *FileBrowse) UpdateProject() {
	fb.StandardPlan()
	fb.SetTitle(fmt.Sprintf("FileBrowse of: %v", fb.ProjectRoot)) // todo: get rid of title
	fb.UpdateFiles()
	fb.ConfigSplits()
}

// ProjectPathParse parses given project path into a root directory (which could
// be the path or just the directory portion of the path, depending in whether
// the path is a directory or not), and a bool if all is good (otherwise error
// message has been reported). projnm is always the last directory of the path.
func ProjectPathParse(path string) (root, projnm, fnm string, ok bool) {
	if path == "" {
		return "", "blank", "", false
	}
	info, err := os.Lstat(path)
	if err != nil {
		emsg := fmt.Errorf("ProjectPathParse: Cannot open at given path: %q: Error: %v", path, err)
		log.Println(emsg)
		return
	}
	dir, fn := filepath.Split(path)
	pathIsDir := info.IsDir()
	if pathIsDir {
		root = path
	} else {
		root = dir
		fnm = fn
	}
	_, projnm = filepath.Split(root)
	ok = true
	return
}

//////////////////////////////////////////////////////////////////////////////////////
//   TextEditors

// ActiveTextEditor returns the currently active TextEditor
func (fb *FileBrowse) ActiveTextEditor() *texteditor.Editor {
	return fb.TextEditorByIndex(fb.ActiveTextEditorIndex)
}

// SetActiveTextEditor sets the given view index as the currently active
// TextEditor -- returns that texteditor
func (fb *FileBrowse) SetActiveTextEditor(idx int) *texteditor.Editor {
	if idx < 0 || idx >= fb.NTextEditors {
		log.Printf("FileBrowse SetActiveTextEditor: text view index out of range: %v\n", idx)
		return nil
	}
	fb.ActiveTextEditorIndex = idx
	av := fb.ActiveTextEditor()
	if av.Buffer != nil {
		fb.ActiveFilename = av.Buffer.Filename
	}
	av.SetFocusEvent()
	return av
}

// NextTextEditor returns the next text view available for viewing a file and
// its index -- if the active text view is empty, then it is used, otherwise
// it is the next one
func (fb *FileBrowse) NextTextEditor() (*texteditor.Editor, int) {
	av := fb.TextEditorByIndex(fb.ActiveTextEditorIndex)
	if av.Buffer == nil {
		return av, fb.ActiveTextEditorIndex
	}
	nxt := (fb.ActiveTextEditorIndex + 1) % fb.NTextEditors
	return fb.TextEditorByIndex(nxt), nxt
}

// SaveActiveView saves the contents of the currently active texteditor
func (fb *FileBrowse) SaveActiveView() { //types:add
	tv := fb.ActiveTextEditor()
	if tv.Buffer != nil {
		tv.Buffer.Save() // todo: errs..
		fb.UpdateFiles()
	}
}

// SaveActiveViewAs save with specified filename the contents of the
// currently active texteditor
func (fb *FileBrowse) SaveActiveViewAs(filename core.Filename) { //types:add
	tv := fb.ActiveTextEditor()
	if tv.Buffer != nil {
		tv.Buffer.SaveAs(filename)
	}
}

// ViewFileNode sets the next text view to view file in given node (opens
// buffer if not already opened)
func (fb *FileBrowse) ViewFileNode(fn *filetree.Node) {
	if _, err := fn.OpenBuf(); err == nil {
		nv, nidx := fb.NextTextEditor()
		if nv.Buffer != nil && nv.Buffer.IsNotSaved() { // todo: save current changes?
			fmt.Printf("Changes not saved in file: %v before switching view there to new file\n", nv.Buffer.Filename)
		}
		nv.SetBuffer(fn.Buffer)
		fn.Buffer.Hi.Style = "emacs" // todo prefs
		fb.SetActiveTextEditor(nidx)
		fb.UpdateFiles()
	}
}

// ViewFile sets the next text view to view given file name -- include as much
// of name as possible to disambiguate -- will use the first matching --
// returns false if not found
func (fb *FileBrowse) ViewFile(fnm string) bool {
	fn, ok := fb.Files.FindFile(fnm)
	if !ok {
		return false
	}
	fb.ViewFileNode(fn)
	return true
}

//////////////////////////////////////////////////////////////////////////////////////
//   GUI plans

// StandardFramePlan returns a Plan for configuring a standard Frame
// -- can modify as desired before calling Build on Frame using this
func (fb *FileBrowse) StandardFramePlan() tree.TypePlan {
	plan := tree.TypePlan{}
	plan.Add(core.TextType, "title")
	plan.Add(core.SplitsType, "splits")
	return plan
}

// StandardPlan configures a standard setup of the overall Frame.
// It returns whether any modifications were made.
func (fb *FileBrowse) StandardPlan() bool {
	plan := fb.StandardFramePlan()
	return tree.Update(fb, plan)
}

// SetTitle sets the optional title and updates the title text
func (fb *FileBrowse) SetTitle(title string) {
	t := fb.TitleWidget()
	if t != nil {
		t.Text = title
	}
}

// Title returns the title text widget.
func (fb *FileBrowse) TitleWidget() *core.Text {
	return fb.ChildByName("title", 0).(*core.Text)
}

// Splits returns the main Splits widget.
func (fb *FileBrowse) Splits() *core.Splits {
	return fb.ChildByName("splits", 2).(*core.Splits)
}

// TextEditorByIndex returns the TextEditor by index, nil if not found
func (fb *FileBrowse) TextEditorByIndex(idx int) *texteditor.Editor {
	if idx < 0 || idx >= fb.NTextEditors {
		log.Printf("FileBrowse: text view index out of range: %v\n", idx)
		return nil
	}
	split := fb.Splits()
	stidx := 1 // 0 = file browser -- could be collapsed but always there.
	if split != nil {
		svk := split.Child(stidx + idx)
		return svk.(*texteditor.Editor)
	}
	return nil
}

func (fb *FileBrowse) MakeToolbar(p *core.Plan) { //types:add
	core.Add(p, func(w *views.FuncButton) {
		w.SetFunc(fb.UpdateFiles).SetIcon(icons.Refresh).SetShortcut("Command+U")
	})
	core.Add(p, func(w *views.FuncButton) {
		w.SetFunc(fb.OpenPath).SetKey(keymap.Open)
		w.Args[0].SetValue(fb.ActiveFilename)
		w.Args[0].SetTag(`ext:".json"`)
	})
	core.Add(p, func(w *views.FuncButton) {
		w.SetFunc(fb.SaveActiveView).SetKey(keymap.Save)
		w.Styler(func(s *styles.Style) {
			s.SetEnabled(fb.Changed && fb.ActiveFilename != "")
		})
	})
	core.Add(p, func(w *views.FuncButton) {
		w.SetFunc(fb.SaveActiveViewAs).SetKey(keymap.SaveAs)
		w.Args[0].SetValue(fb.ActiveFilename)
		w.Args[0].SetTag(`ext:".json"`)
	})
}

// SplitsPlan returns a Plan for configuring the Splits
func (fb *FileBrowse) SplitsPlan() tree.TypePlan {
	plan := tree.TypePlan{}
	plan.Add(core.FrameType, "filetree-fr")
	for i := 0; i < fb.NTextEditors; i++ {
		plan.Add(texteditor.EditorType, fmt.Sprintf("texteditor-%v", i))
	}
	// todo: tab view
	return plan
}

// ConfigSplits configures the Splits.
func (fb *FileBrowse) ConfigSplits() {
	split := fb.Splits()
	if split == nil {
		return
	}
	split.SetSplits(.2, .4, .4)

	plan := fb.SplitsPlan()
	if tree.Update(split, plan) {
		ftfr := split.Child(0).(*core.Frame)
		fb.Files = filetree.NewTree(ftfr)
		fb.Files.OnSelect(func(e events.Event) {
			e.SetHandled()
			if len(fb.Files.SelectedNodes) > 0 {
				sn, ok := fb.Files.SelectedNodes[0].This().(*filetree.Node)
				if ok {
					fb.FileNodeSelected(sn)
				}
			}
		})
		fb.Files.DoubleClickFun = func(e events.Event) {
			e.SetHandled()
			if len(fb.Files.SelectedNodes) > 0 {
				sn, ok := fb.Files.SelectedNodes[0].This().(*filetree.Node)
				if ok {
					fb.FileNodeOpened(sn)
				}
			}
		}
	}
}

func (fb *FileBrowse) FileNodeSelected(fn *filetree.Node) {
	fmt.Println("selected:", fn.FPath)
}

func (fb *FileBrowse) FileNodeOpened(fn *filetree.Node) {
	if fn.IsDir() {
		if !fn.HasChildren() {
			fn.OpenEmptyDir()
		} else {
			fn.ToggleClose()
		}
	} else {
		fb.ViewFileNode(fn)
		fn.UpdateNode()
	}
}

//////////////////////////////////////////////////////////////////////////////////////
//   Project window

// NewFileBrowser creates a new FileBrowse window with a new FileBrowse project for given
// path, returning the window and the path
func NewFileBrowser(path string) (*FileBrowse, *core.Stage) {
	_, projnm, _, _ := ProjectPathParse(path)

	b := core.NewBody("Browser: " + projnm)
	fb := NewFileBrowse(b)
	b.AddAppBar(fb.MakeToolbar)
	fb.OpenPath(core.Filename(path))
	return fb, b.RunWindow()
}

//////////////////////////////////////////////////////////////////////////////////////
//  main

func main() {
	var path string

	// process command args
	if len(os.Args) > 1 {
		flag.StringVar(&path, "path", "", "path to open -- can be to a directory or a filename within the directory")
		// todo: other args?
		flag.Parse()
		if path == "" {
			if flag.NArg() > 0 {
				path = flag.Arg(0)
			}
		}
	}
	if path == "" {
		path = "./"
	}
	if path != "" {
		path, _ = filepath.Abs(path)
	}
	fmt.Println("path:", path)
	_, st := NewFileBrowser(path)
	st.Wait()
}
