// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"fmt"
	"reflect"
	"sort"
	"unicode/utf8"

	"github.com/goki/gi/gist"
	"github.com/goki/gi/icons"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/oswin/key"
	"github.com/goki/gi/units"
	"github.com/goki/ki/ints"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
)

// ComboBox is for selecting items from a dropdown list, with an optional
// edit TextField for typing directly.
// The items can be of any type, including enum values -- they are converted
// to strings for the display.  If the items are IconName type, then they
// are displayed using icons instead.
type ComboBox struct {
	ButtonBase
	Editable  bool      `xml:"editable" desc:"provide a text field for editing the value, or just a button for selecting items?  Set the editable property"`
	CurVal    any       `json:"-" xml:"-" desc:"current selected value"`
	CurIndex  int       `json:"-" xml:"-" desc:"current index in list of possible items"`
	Items     []any     `json:"-" xml:"-" desc:"items available for selection"`
	ItemsMenu Menu      `json:"-" xml:"-" desc:"the menu of actions for selecting items -- automatically generated from Items"`
	ComboSig  ki.Signal `copy:"-" json:"-" xml:"-" view:"-" desc:"signal for combo box, when a new value has been selected -- the signal type is the index of the selected item, and the data is the value"`
	MaxLength int       `desc:"maximum label length (in runes)"`
}

var KiT_ComboBox = kit.Types.AddType(&ComboBox{}, ComboBoxProps)

// AddNewComboBox adds a new button to given parent node, with given name.
func AddNewComboBox(parent ki.Ki, name string) *ComboBox {
	return parent.AddNewChild(KiT_ComboBox, name).(*ComboBox)
}

func (cb *ComboBox) CopyFieldsFrom(frm any) {
	fr := frm.(*ComboBox)
	cb.ButtonBase.CopyFieldsFrom(&fr.ButtonBase)
	cb.Editable = fr.Editable
	cb.CurVal = fr.CurVal
	cb.CurIndex = fr.CurIndex
	cb.Items = fr.Items
	cb.ItemsMenu.CopyFrom(&fr.ItemsMenu)
	cb.MaxLength = fr.MaxLength
}

func (cb *ComboBox) Disconnect() {
	cb.ButtonBase.Disconnect()
	cb.ComboSig.DisconnectAll()
}

// DefaultStyle implements the [DefaultStyler] interface
func (cb *ComboBox) DefaultStyle() {
	cs := CurrentColorScheme()
	s := &cb.Style

	s.Border.Style.Set(gist.BorderNone)
	s.Border.Radius.Set(units.Px(4))
	s.Layout.Padding.Set(units.Px(4))
	s.Layout.Margin.Set(units.Px(4))
	s.Text.Align = gist.AlignCenter
	s.Font.BgColor.SetColor(cs.Background)
	s.Font.Color.SetColor(cs.Font)
}

var ComboBoxProps = ki.Props{
	"EnumType:Flag":    KiT_ButtonFlags,
	"border-width":     units.Px(1),
	"border-radius":    units.Px(4),
	"border-color":     &Prefs.Colors.Border,
	"padding":          units.Px(4),
	"margin":           units.Px(4),
	"text-align":       gist.AlignCenter,
	"background-color": &Prefs.Colors.Control,
	"color":            &Prefs.Colors.Font,
	"#icon": ki.Props{
		"width":   units.Em(1),
		"height":  units.Em(1),
		"margin":  units.Px(0),
		"padding": units.Px(0),
		"fill":    &Prefs.Colors.Icon,
		"stroke":  &Prefs.Colors.Font,
	},
	"#label": ki.Props{
		"margin":  units.Px(0),
		"padding": units.Px(0),
	},
	"#text": ki.Props{
		"margin":    units.Px(1),
		"padding":   units.Px(1),
		"max-width": -1,
		"width":     units.Ch(12),
	},
	"#indicator": ki.Props{
		"width":          units.Ex(1.5),
		"height":         units.Ex(1.5),
		"margin":         units.Px(0),
		"padding":        units.Px(0),
		"vertical-align": gist.AlignBottom,
		"fill":           &Prefs.Colors.Icon,
		"stroke":         &Prefs.Colors.Font,
	},
	"#ind-stretch": ki.Props{
		"width": units.Em(1),
	},
	ButtonSelectors[ButtonActive]: ki.Props{
		"background-color": "linear-gradient(lighter-0, highlight-10)",
	},
	ButtonSelectors[ButtonInactive]: ki.Props{
		"border-color": "highlight-50",
		"color":        "highlight-50",
	},
	ButtonSelectors[ButtonHover]: ki.Props{
		"background-color": "linear-gradient(highlight-10, highlight-10)",
	},
	ButtonSelectors[ButtonFocus]: ki.Props{
		"border-width":     units.Px(2),
		"background-color": "linear-gradient(samelight-50, highlight-10)",
	},
	ButtonSelectors[ButtonDown]: ki.Props{
		"color":            "highlight-90",
		"background-color": "linear-gradient(highlight-30, highlight-10)",
	},
	ButtonSelectors[ButtonSelected]: ki.Props{
		"background-color": "linear-gradient(pref(Select), highlight-10)",
		"color":            "highlight-90",
	},
}

// ButtonWidget interface

func (cb *ComboBox) ButtonRelease() {
	if cb.IsInactive() {
		return
	}
	wasPressed := (cb.State == ButtonDown)
	cb.MakeItemsMenu()
	if len(cb.ItemsMenu) == 0 {
		return
	}
	updt := cb.UpdateStart()
	cb.SetButtonState(ButtonActive)
	cb.ButtonSig.Emit(cb.This(), int64(ButtonReleased), nil)
	if wasPressed {
		cb.ButtonSig.Emit(cb.This(), int64(ButtonClicked), nil)
	}
	cb.UpdateEnd(updt)
	cb.BBoxMu.RLock()
	pos := cb.WinBBox.Max
	if pos.X == 0 && pos.Y == 0 { // offscreen
		pos = cb.ObjBBox.Max
	}
	indic := cb.Parts.ChildByName("indicator", 3)
	if indic != nil {
		pos = KiToNode2DBase(indic).WinBBox.Min
		if pos.X == 0 && pos.Y == 0 {
			pos = KiToNode2DBase(indic).ObjBBox.Min
		}
	} else {
		pos.Y -= 10
		pos.X -= 10
	}
	cb.BBoxMu.RUnlock()
	PopupMenu(cb.ItemsMenu, pos.X, pos.Y, cb.Viewport, cb.Text)
}

// ConfigPartsIconText returns a standard config for creating parts, of icon
// and text left-to right in a row -- always makes text
func (cb *ComboBox) ConfigPartsIconText(config *kit.TypeAndNameList, icnm icons.Icon) (icIdx, txIdx int) {
	// todo: add some styles for button layout
	icIdx = -1
	txIdx = -1
	if TheIconMgr.IsValid(icnm) {
		icIdx = len(*config)
		config.Add(KiT_Icon, "icon")
		config.Add(KiT_Space, "space")
	}
	txIdx = len(*config)
	config.Add(KiT_TextField, "text")
	return
}

// ConfigPartsSetText sets part style props, using given props if not set in
// object props
func (cb *ComboBox) ConfigPartsSetText(txt string, txIdx, icIdx, indIdx int) {
	if txIdx >= 0 {
		tx := cb.Parts.Child(txIdx).(*TextField)
		tx.SetText(txt)
		if _, err := tx.PropTry("__comboInit"); err != nil {
			cb.StylePart(Node2D(tx))
			if icIdx >= 0 {
				cb.StylePart(cb.Parts.Child(txIdx - 1).(Node2D)) // also get the space
			}
			tx.SetProp("__comboInit", true)
			if cb.MaxLength > 0 {
				tx.SetMinPrefWidth(units.Ch(float32(cb.MaxLength)))
			}
			if indIdx > 0 {
				ispc := cb.Parts.Child(indIdx - 1).(Node2D)
				ispc.SetProp("max-width", 0)
			}
		}
	}
}

// ConfigPartsAddIndicatorSpace adds indicator with a space instead of a stretch
// for editable combobox, where textfield then takes up the rest of the space
func (bb *ButtonBase) ConfigPartsAddIndicatorSpace(config *kit.TypeAndNameList, defOn bool) int {
	needInd := (bb.HasMenu() || defOn) && bb.Indicator != "none"
	if !needInd {
		return -1
	}
	indIdx := -1
	config.Add(KiT_Space, "ind-stretch")
	indIdx = len(*config)
	config.Add(KiT_Icon, "indicator")
	return indIdx
}

func (cb *ComboBox) ConfigPartsIfNeeded() {
	if cb.Editable {
		cn := cb.Parts.ChildByName("text", 2)
		if !cb.PartsNeedUpdateIconLabel(cb.Icon, "") && cn != nil {
			return
		}
	} else {
		if !cb.PartsNeedUpdateIconLabel(cb.Icon, cb.Text) {
			return
		}
	}
	cb.This().(ButtonWidget).ConfigParts()
}

func (cb *ComboBox) ConfigParts() {
	if eb, err := cb.PropTry("editable"); err == nil {
		cb.Editable, _ = kit.ToBool(eb)
	}
	config := kit.TypeAndNameList{}
	var icIdx, lbIdx, txIdx, indIdx int
	if cb.Editable {
		lbIdx = -1
		icIdx, txIdx = cb.ConfigPartsIconText(&config, cb.Icon)
		cb.SetProp("no-focus", true)
		indIdx = cb.ConfigPartsAddIndicatorSpace(&config, true) // use space instead of stretch
	} else {
		txIdx = -1
		icIdx, lbIdx = cb.ConfigPartsIconLabel(&config, cb.Icon, cb.Text)
		indIdx = cb.ConfigPartsAddIndicator(&config, true) // default on
	}
	mods, updt := cb.Parts.ConfigChildren(config)
	cb.ConfigPartsSetIconLabel(cb.Icon, cb.Text, icIdx, lbIdx)
	cb.ConfigPartsIndicator(indIdx)
	if txIdx >= 0 {
		cb.ConfigPartsSetText(cb.Text, txIdx, icIdx, indIdx)
	}
	if cb.MaxLength > 0 && lbIdx >= 0 {
		lbl := cb.Parts.Child(lbIdx).(*Label)
		lbl.SetMinPrefWidth(units.Ch(float32(cb.MaxLength)))
	}
	if mods {
		cb.UpdateEnd(updt)
	}
}

// TextField returns the text field of an editable combobox, and false if not made
func (cb *ComboBox) TextField() (*TextField, bool) {
	tff := cb.Parts.ChildByName("text", 2)
	if tff == nil {
		return nil, false
	}
	return tff.(*TextField), true
}

// MakeItems makes sure the Items list is made, and if not, or reset is true,
// creates one with the given capacity
func (cb *ComboBox) MakeItems(reset bool, capacity int) {
	if cb.Items == nil || reset {
		cb.Items = make([]any, 0, capacity)
	}
}

// SortItems sorts the items according to their labels
func (cb *ComboBox) SortItems(ascending bool) {
	sort.Slice(cb.Items, func(i, j int) bool {
		if ascending {
			return ToLabel(cb.Items[i]) < ToLabel(cb.Items[j])
		} else {
			return ToLabel(cb.Items[i]) > ToLabel(cb.Items[j])
		}
	})
}

// SetToMaxLength gets the maximum label length so that the width of the
// button label is automatically set according to the max length of all items
// in the list -- if maxLen > 0 then it is used as an upper do-not-exceed
// length
func (cb *ComboBox) SetToMaxLength(maxLen int) {
	ml := 0
	for _, it := range cb.Items {
		ml = ints.MaxInt(ml, utf8.RuneCountInString(ToLabel(it)))
	}
	if maxLen > 0 {
		ml = ints.MinInt(ml, maxLen)
	}
	cb.MaxLength = ml
}

// ItemsFromTypes sets the Items list from a list of types -- see e.g.,
// AllImplementersOf or AllEmbedsOf in kit.TypeRegistry -- if setFirst then
// set current item to the first item in the list, sort sorts the list in
// ascending order, and maxLen if > 0 auto-sets the width of the button to the
// contents, with the given upper limit
func (cb *ComboBox) ItemsFromTypes(tl []reflect.Type, setFirst, sort bool, maxLen int) {
	sz := len(tl)
	if sz == 0 {
		return
	}
	cb.Items = make([]any, sz)
	for i, typ := range tl {
		cb.Items[i] = typ
	}
	if sort {
		cb.SortItems(true)
	}
	if maxLen > 0 {
		cb.SetToMaxLength(maxLen)
	}
	if setFirst {
		cb.SetCurIndex(0)
	}
}

// ItemsFromStringList sets the Items list from a list of string values -- if
// setFirst then set current item to the first item in the list, and maxLen if
// > 0 auto-sets the width of the button to the contents, with the given upper
// limit
func (cb *ComboBox) ItemsFromStringList(el []string, setFirst bool, maxLen int) {
	sz := len(el)
	if sz == 0 {
		return
	}
	cb.Items = make([]any, sz)
	for i, str := range el {
		cb.Items[i] = str
	}
	if maxLen > 0 {
		cb.SetToMaxLength(maxLen)
	}
	if setFirst {
		cb.SetCurIndex(0)
	}
}

// ItemsFromIconList sets the Items list from a list of icons.Icon values -- if
// setFirst then set current item to the first item in the list, and maxLen if
// > 0 auto-sets the width of the button to the contents, with the given upper
// limit
func (cb *ComboBox) ItemsFromIconList(el []icons.Icon, setFirst bool, maxLen int) {
	sz := len(el)
	if sz == 0 {
		return
	}
	cb.Items = make([]any, sz)
	for i, str := range el {
		cb.Items[i] = str
	}
	if maxLen > 0 {
		cb.SetToMaxLength(maxLen)
	}
	if setFirst {
		cb.SetCurIndex(0)
	}
}

// ItemsFromEnumList sets the Items list from a list of enum values (see
// kit.EnumRegistry) -- if setFirst then set current item to the first item in
// the list, and maxLen if > 0 auto-sets the width of the button to the
// contents, with the given upper limit
func (cb *ComboBox) ItemsFromEnumList(el []kit.EnumValue, setFirst bool, maxLen int) {
	sz := len(el)
	if sz == 0 {
		return
	}
	cb.Items = make([]any, sz)
	for i, enum := range el {
		cb.Items[i] = enum
	}
	if maxLen > 0 {
		cb.SetToMaxLength(maxLen)
	}
	if setFirst {
		cb.SetCurIndex(0)
	}
}

// ItemsFromEnum sets the Items list from an enum type, which must be
// registered on kit.EnumRegistry -- if setFirst then set current item to the
// first item in the list, and maxLen if > 0 auto-sets the width of the button
// to the contents, with the given upper limit -- see kit.EnumRegistry, and
// maxLen if > 0 auto-sets the width of the button to the contents, with the
// given upper limit
func (cb *ComboBox) ItemsFromEnum(enumtyp reflect.Type, setFirst bool, maxLen int) {
	cb.ItemsFromEnumList(kit.Enums.TypeValues(enumtyp, true), setFirst, maxLen)
}

// FindItem finds an item on list of items and returns its index
func (cb *ComboBox) FindItem(it any) int {
	if cb.Items == nil {
		return -1
	}
	for i, v := range cb.Items {
		if v == it {
			return i
		}
	}
	return -1
}

// SetCurVal sets the current value (CurVal) and the corresponding CurIndex
// for that item on the current Items list (adds to items list if not found)
// -- returns that index -- and sets the text to the string value of that
// value (using standard Stringer string conversion)
func (cb *ComboBox) SetCurVal(it any) int {
	cb.CurVal = it
	cb.CurIndex = cb.FindItem(it)
	if cb.CurIndex < 0 { // add to list if not found..
		cb.CurIndex = len(cb.Items)
		cb.Items = append(cb.Items, it)
	}
	cb.ShowCurVal()
	return cb.CurIndex
}

// SetCurIndex sets the current index (CurIndex) and the corresponding CurVal
// for that item on the current Items list (-1 if not found) -- returns value
// -- and sets the text to the string value of that value (using standard
// Stringer string conversion)
func (cb *ComboBox) SetCurIndex(idx int) any {
	cb.CurIndex = idx
	if idx < 0 || idx >= len(cb.Items) {
		cb.CurVal = nil
		cb.SetText(fmt.Sprintf("idx %v > len", idx))
	} else {
		cb.CurVal = cb.Items[idx]
		cb.ShowCurVal()
	}
	return cb.CurVal
}

// ShowCurVal updates the display to present the
// currently-selected value (CurVal)
func (cb *ComboBox) ShowCurVal() {
	if icnm, isic := cb.CurVal.(icons.Icon); isic {
		cb.SetIcon(icnm)
	} else {
		cb.SetText(ToLabel(cb.CurVal))
	}
}

// SelectItem selects a given item and updates the display to it
func (cb *ComboBox) SelectItem(idx int) {
	if cb.This() == nil {
		return
	}
	updt := cb.UpdateStart()
	cb.SetCurIndex(idx)
	cb.UpdateEnd(updt)
}

// SelectItemAction selects a given item and emits the index as the ComboSig signal
// and the selected item as the data.
func (cb *ComboBox) SelectItemAction(idx int) {
	if cb.This() == nil {
		return
	}
	updt := cb.UpdateStart()
	cb.SelectItem(idx)
	cb.ComboSig.Emit(cb.This(), int64(cb.CurIndex), cb.CurVal)
	cb.UpdateEnd(updt)
}

// MakeItemsMenu makes menu of all the items
func (cb *ComboBox) MakeItemsMenu() {
	nitm := len(cb.Items)
	if cb.ItemsMenu == nil {
		cb.ItemsMenu = make(Menu, 0, nitm)
	}
	sz := len(cb.ItemsMenu)
	if nitm < sz {
		cb.ItemsMenu = cb.ItemsMenu[0:nitm]
	}
	if nitm == 0 {
		return
	}
	_, ics := cb.Items[0].(icons.Icon) // if true, we render as icons
	for i, it := range cb.Items {
		var ac *Action
		if sz > i {
			ac = cb.ItemsMenu[i].(*Action)
		} else {
			ac = &Action{}
			ki.InitNode(ac)
			cb.ItemsMenu = append(cb.ItemsMenu, ac.This().(Node2D))
		}
		nm := fmt.Sprintf("Item_%v", i)
		ac.SetName(nm)
		if ics {
			ac.Icon = it.(icons.Icon)
			ac.Tooltip = string(ac.Icon)
		} else {
			ac.Text = ToLabel(it)
			fmt.Printf("text %s type %T\n", ac.Text, it)
			if d, ok := it.(kit.Describer); ok {
				ac.Tooltip = d.Desc()
			} else if ev, ok := it.(kit.EnumValue); ok {
				ac.Tooltip = ev.Desc
			}
		}
		ac.Data = i // index is the data
		ac.SetSelectedState(i == cb.CurIndex)
		ac.SetAsMenu()
		ac.ActionSig.ConnectOnly(cb.This(), func(recv, send ki.Ki, sig int64, data any) {
			idx := data.(int)
			cbb := recv.(*ComboBox)
			cbb.SelectItemAction(idx)
		})
	}
}

func (cb *ComboBox) HasFocus2D() bool {
	if cb.IsInactive() {
		return false
	}
	return cb.ContainsFocus() // needed for getting key events
}

func (cb *ComboBox) ConnectEvents2D() {
	cb.ButtonEvents()
	cb.KeyChordEvent()
}

func (cb *ComboBox) KeyChordEvent() {
	cb.ConnectEvent(oswin.KeyChordEvent, HiPri, func(recv, send ki.Ki, sig int64, d any) {
		cbb := recv.(*ComboBox)
		if cbb.IsInactive() {
			return
		}
		kt := d.(*key.ChordEvent)
		if KeyEventTrace {
			fmt.Printf("ComboBox KeyChordEvent: %v\n", cbb.Path())
		}
		kf := KeyFun(kt.Chord())
		switch {
		case kf == KeyFunMoveUp:
			kt.SetProcessed()
			if len(cbb.Items) > 0 {
				idx := cbb.CurIndex - 1
				if idx < 0 {
					idx += len(cbb.Items)
				}
				cbb.SelectItemAction(idx)
			}
		case kf == KeyFunMoveDown:
			kt.SetProcessed()
			if len(cbb.Items) > 0 {
				idx := cbb.CurIndex + 1
				if idx >= len(cbb.Items) {
					idx -= len(cbb.Items)
				}
				cbb.SelectItemAction(idx)
			}
		case kf == KeyFunPageUp:
			kt.SetProcessed()
			if len(cbb.Items) > 10 {
				idx := cbb.CurIndex - 10
				for idx < 0 {
					idx += len(cbb.Items)
				}
				cbb.SelectItemAction(idx)
			}
		case kf == KeyFunPageDown:
			kt.SetProcessed()
			if len(cbb.Items) > 10 {
				idx := cbb.CurIndex + 10
				for idx >= len(cbb.Items) {
					idx -= len(cbb.Items)
				}
				cbb.SelectItemAction(idx)
			}
		case kf == KeyFunEnter || (!cbb.Editable && kt.Rune == ' '):
			if !(kt.Rune == ' ' && cbb.Viewport.IsCompleter()) {
				kt.SetProcessed()
				cbb.ButtonPress()
				cbb.ButtonRelease()
			}
		}
	})
}
