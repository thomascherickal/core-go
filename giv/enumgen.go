// Code generated by "goki generate"; DO NOT EDIT.

package giv

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
	"goki.dev/gi/v2/gi"
)

var _SliceViewFlagsValues = []SliceViewFlags{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

// SliceViewFlagsN is the highest valid value
// for type SliceViewFlags, plus one.
const SliceViewFlagsN SliceViewFlags = 19

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _SliceViewFlagsNoOp() {
	var x [1]struct{}
	_ = x[SliceViewConfiged-(8)]
	_ = x[SliceViewNoAdd-(9)]
	_ = x[SliceViewNoDelete-(10)]
	_ = x[SliceViewShowViewCtxtMenu-(11)]
	_ = x[SliceViewIsArray-(12)]
	_ = x[SliceViewShowIndex-(13)]
	_ = x[SliceViewReadOnlyKeyNav-(14)]
	_ = x[SliceViewSelectMode-(15)]
	_ = x[SliceViewReadOnlyMultiSel-(16)]
	_ = x[SliceViewInFocusGrab-(17)]
	_ = x[SliceViewInFullRebuild-(18)]
}

var _SliceViewFlagsNameToValueMap = map[string]SliceViewFlags{
	`Configed`:         8,
	`configed`:         8,
	`NoAdd`:            9,
	`noadd`:            9,
	`NoDelete`:         10,
	`nodelete`:         10,
	`ShowViewCtxtMenu`: 11,
	`showviewctxtmenu`: 11,
	`IsArray`:          12,
	`isarray`:          12,
	`ShowIndex`:        13,
	`showindex`:        13,
	`ReadOnlyKeyNav`:   14,
	`readonlykeynav`:   14,
	`SelectMode`:       15,
	`selectmode`:       15,
	`ReadOnlyMultiSel`: 16,
	`readonlymultisel`: 16,
	`InFocusGrab`:      17,
	`infocusgrab`:      17,
	`InFullRebuild`:    18,
	`infullrebuild`:    18,
}

var _SliceViewFlagsDescMap = map[SliceViewFlags]string{
	8:  `indicates that the widgets have been configured and`,
	9:  `if true, user cannot add elements to the slice`,
	10: `if true, user cannot delete elements from the slice`,
	11: `if the type we&#39;re viewing has its own CtxtMenu property defined, should we also still show the view&#39;s standard context menu?`,
	12: `whether the slice is actually an array -- no modifications -- set by SetSlice`,
	13: `whether to show index or not`,
	14: `support key navigation when ReadOnly (default true) -- no focus really plausible in ReadOnly case, so it uses a low-pri capture of up / down events`,
	15: `editing-mode select rows mode`,
	16: `if view is ReadOnly, default selection mode is to choose one row only -- if this is true, standard multiple selection logic with modifier keys is instead supported`,
	17: `guard for recursive focus grabbing`,
	18: `guard for recursive rebuild`,
}

var _SliceViewFlagsMap = map[SliceViewFlags]string{
	8:  `Configed`,
	9:  `NoAdd`,
	10: `NoDelete`,
	11: `ShowViewCtxtMenu`,
	12: `IsArray`,
	13: `ShowIndex`,
	14: `ReadOnlyKeyNav`,
	15: `SelectMode`,
	16: `ReadOnlyMultiSel`,
	17: `InFocusGrab`,
	18: `InFullRebuild`,
}

// String returns the string representation
// of this SliceViewFlags value.
func (i SliceViewFlags) String() string {
	str := ""
	for _, ie := range gi.WidgetFlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _SliceViewFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this SliceViewFlags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i SliceViewFlags) BitIndexString() string {
	if str, ok := _SliceViewFlagsMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).BitIndexString()
}

// SetString sets the SliceViewFlags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *SliceViewFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the SliceViewFlags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *SliceViewFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _SliceViewFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _SliceViewFlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*gi.WidgetFlags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the SliceViewFlags value as an int64.
func (i SliceViewFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the SliceViewFlags value from an int64.
func (i *SliceViewFlags) SetInt64(in int64) {
	*i = SliceViewFlags(in)
}

// Desc returns the description of the SliceViewFlags value.
func (i SliceViewFlags) Desc() string {
	if str, ok := _SliceViewFlagsDescMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).Desc()
}

// SliceViewFlagsValues returns all possible values
// for the type SliceViewFlags.
func SliceViewFlagsValues() []SliceViewFlags {
	es := gi.WidgetFlagsValues()
	res := make([]SliceViewFlags, len(es))
	for i, e := range es {
		res[i] = SliceViewFlags(e)
	}
	res = append(res, _SliceViewFlagsValues...)
	return res
}

// Values returns all possible values
// for the type SliceViewFlags.
func (i SliceViewFlags) Values() []enums.Enum {
	es := gi.WidgetFlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_SliceViewFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _SliceViewFlagsValues {
		res[i+les] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type SliceViewFlags.
func (i SliceViewFlags) IsValid() bool {
	_, ok := _SliceViewFlagsMap[i]
	if !ok {
		return gi.WidgetFlags(i).IsValid()
	}
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i SliceViewFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *SliceViewFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i SliceViewFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *SliceViewFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _TreeViewFlagsValues = []TreeViewFlags{8, 9}

// TreeViewFlagsN is the highest valid value
// for type TreeViewFlags, plus one.
const TreeViewFlagsN TreeViewFlags = 10

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _TreeViewFlagsNoOp() {
	var x [1]struct{}
	_ = x[TreeViewFlagClosed-(8)]
	_ = x[TreeViewFlagSelectMode-(9)]
}

var _TreeViewFlagsNameToValueMap = map[string]TreeViewFlags{
	`Closed`:     8,
	`closed`:     8,
	`SelectMode`: 9,
	`selectmode`: 9,
}

var _TreeViewFlagsDescMap = map[TreeViewFlags]string{
	8: `TreeViewFlagClosed means node is toggled closed (children not visible) Otherwise Open.`,
	9: `This flag on the Root node determines whether keyboard movements update selection or not.`,
}

var _TreeViewFlagsMap = map[TreeViewFlags]string{
	8: `Closed`,
	9: `SelectMode`,
}

// String returns the string representation
// of this TreeViewFlags value.
func (i TreeViewFlags) String() string {
	str := ""
	for _, ie := range gi.WidgetFlagsValues() {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	for _, ie := range _TreeViewFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this TreeViewFlags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i TreeViewFlags) BitIndexString() string {
	if str, ok := _TreeViewFlagsMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).BitIndexString()
}

// SetString sets the TreeViewFlags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *TreeViewFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the TreeViewFlags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *TreeViewFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _TreeViewFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _TreeViewFlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			err := (*gi.WidgetFlags)(i).SetStringOr(flg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Int64 returns the TreeViewFlags value as an int64.
func (i TreeViewFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the TreeViewFlags value from an int64.
func (i *TreeViewFlags) SetInt64(in int64) {
	*i = TreeViewFlags(in)
}

// Desc returns the description of the TreeViewFlags value.
func (i TreeViewFlags) Desc() string {
	if str, ok := _TreeViewFlagsDescMap[i]; ok {
		return str
	}
	return gi.WidgetFlags(i).Desc()
}

// TreeViewFlagsValues returns all possible values
// for the type TreeViewFlags.
func TreeViewFlagsValues() []TreeViewFlags {
	es := gi.WidgetFlagsValues()
	res := make([]TreeViewFlags, len(es))
	for i, e := range es {
		res[i] = TreeViewFlags(e)
	}
	res = append(res, _TreeViewFlagsValues...)
	return res
}

// Values returns all possible values
// for the type TreeViewFlags.
func (i TreeViewFlags) Values() []enums.Enum {
	es := gi.WidgetFlagsValues()
	les := len(es)
	res := make([]enums.Enum, les+len(_TreeViewFlagsValues))
	for i, d := range es {
		res[i] = d
	}
	for i, d := range _TreeViewFlagsValues {
		res[i+les] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type TreeViewFlags.
func (i TreeViewFlags) IsValid() bool {
	_, ok := _TreeViewFlagsMap[i]
	if !ok {
		return gi.WidgetFlags(i).IsValid()
	}
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i TreeViewFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *TreeViewFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i TreeViewFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *TreeViewFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _ValueFlagsValues = []ValueFlags{0, 1, 2, 3}

// ValueFlagsN is the highest valid value
// for type ValueFlags, plus one.
const ValueFlagsN ValueFlags = 4

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ValueFlagsNoOp() {
	var x [1]struct{}
	_ = x[ValueReadOnly-(0)]
	_ = x[ValueMapKey-(1)]
	_ = x[ValueHasSavedLabel-(2)]
	_ = x[ValueHasSavedDoc-(3)]
}

var _ValueFlagsNameToValueMap = map[string]ValueFlags{
	`ReadOnly`:      0,
	`readonly`:      0,
	`MapKey`:        1,
	`mapkey`:        1,
	`HasSavedLabel`: 2,
	`hassavedlabel`: 2,
	`HasSavedDoc`:   3,
	`hassaveddoc`:   3,
}

var _ValueFlagsDescMap = map[ValueFlags]string{
	0: `flagged after first configuration`,
	1: `for OwnKind = Map, this value represents the Key -- otherwise the Value`,
	2: `ValueHasSavedLabel is whether the value has a saved version of its label, which can be set either automatically or explicitly`,
	3: `ValueHasSavedDoc is whether the value has a saved version of its documentation, which can be set either automatically or explicitly`,
}

var _ValueFlagsMap = map[ValueFlags]string{
	0: `ReadOnly`,
	1: `MapKey`,
	2: `HasSavedLabel`,
	3: `HasSavedDoc`,
}

// String returns the string representation
// of this ValueFlags value.
func (i ValueFlags) String() string {
	str := ""
	for _, ie := range _ValueFlagsValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this ValueFlags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i ValueFlags) BitIndexString() string {
	if str, ok := _ValueFlagsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the ValueFlags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *ValueFlags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the ValueFlags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *ValueFlags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _ValueFlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _ValueFlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else if flg == "" {
			continue
		} else {
			return fmt.Errorf("%q is not a valid value for type ValueFlags", flg)
		}
	}
	return nil
}

// Int64 returns the ValueFlags value as an int64.
func (i ValueFlags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the ValueFlags value from an int64.
func (i *ValueFlags) SetInt64(in int64) {
	*i = ValueFlags(in)
}

// Desc returns the description of the ValueFlags value.
func (i ValueFlags) Desc() string {
	if str, ok := _ValueFlagsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ValueFlagsValues returns all possible values
// for the type ValueFlags.
func ValueFlagsValues() []ValueFlags {
	return _ValueFlagsValues
}

// Values returns all possible values
// for the type ValueFlags.
func (i ValueFlags) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ValueFlagsValues))
	for i, d := range _ValueFlagsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type ValueFlags.
func (i ValueFlags) IsValid() bool {
	_, ok := _ValueFlagsMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i ValueFlags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *ValueFlags) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i ValueFlags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *ValueFlags) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
