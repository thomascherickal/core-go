// Code generated by "enumgen"; DO NOT EDIT.

package ki

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

var _FlagsValues = []Flags{0, 1, 2, 3, 4, 5, 6, 7, 8}

// FlagsN is the highest valid value
// for type Flags, plus one.
const FlagsN Flags = 9

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _FlagsNoOp() {
	var x [1]struct{}
	_ = x[IsField-(0)]
	_ = x[Updating-(1)]
	_ = x[OnlySelfUpdate-(2)]
	_ = x[NodeDeleted-(3)]
	_ = x[NodeDestroyed-(4)]
	_ = x[ChildAdded-(5)]
	_ = x[ChildDeleted-(6)]
	_ = x[ChildrenDeleted-(7)]
	_ = x[ValUpdated-(8)]
}

var _FlagsNameToValueMap = map[string]Flags{
	`IsField`:         0,
	`isfield`:         0,
	`Updating`:        1,
	`updating`:        1,
	`OnlySelfUpdate`:  2,
	`onlyselfupdate`:  2,
	`NodeDeleted`:     3,
	`nodedeleted`:     3,
	`NodeDestroyed`:   4,
	`nodedestroyed`:   4,
	`ChildAdded`:      5,
	`childadded`:      5,
	`ChildDeleted`:    6,
	`childdeleted`:    6,
	`ChildrenDeleted`: 7,
	`childrendeleted`: 7,
	`ValUpdated`:      8,
	`valupdated`:      8,
}

var _FlagsDescMap = map[Flags]string{
	0: `IsField indicates a node is a field in its parent node, not a child in children.`,
	1: `Updating flag is set at UpdateStart and cleared if we were the first updater at UpdateEnd.`,
	2: `OnlySelfUpdate means that the UpdateStart / End logic only applies to this node in isolation, not to its children -- useful for a parent node that has a different functional role than its children.`,
	3: `NodeDeleted means this node has been deleted.`,
	4: `NodeDestroyed means this node has been destroyed -- do not trigger any more update signals on it.`,
	5: `ChildAdded means one or more new children were added to the node.`,
	6: `ChildDeleted means one or more children were deleted from the node.`,
	7: `ChildrenDeleted means all children were deleted.`,
	8: `ValUpdated means a value was updated (Field, Prop, any kind of value)`,
}

var _FlagsMap = map[Flags]string{
	0: `IsField`,
	1: `Updating`,
	2: `OnlySelfUpdate`,
	3: `NodeDeleted`,
	4: `NodeDestroyed`,
	5: `ChildAdded`,
	6: `ChildDeleted`,
	7: `ChildrenDeleted`,
	8: `ValUpdated`,
}

// String returns the string representation
// of this Flags value.
func (i Flags) String() string {
	str := ""
	for _, ie := range _FlagsValues {
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
// representation of this Flags value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i Flags) BitIndexString() string {
	if str, ok := _FlagsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Flags value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Flags) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the Flags value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *Flags) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _FlagsNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _FlagsNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type Flags")
		}
	}
	return nil
}

// Int64 returns the Flags value as an int64.
func (i Flags) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Flags value from an int64.
func (i *Flags) SetInt64(in int64) {
	*i = Flags(in)
}

// Desc returns the description of the Flags value.
func (i Flags) Desc() string {
	if str, ok := _FlagsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// FlagsValues returns all possible values
// for the type Flags.
func FlagsValues() []Flags {
	return _FlagsValues
}

// Values returns all possible values
// for the type Flags.
func (i Flags) Values() []enums.Enum {
	res := make([]enums.Enum, len(_FlagsValues))
	for i, d := range _FlagsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Flags.
func (i Flags) IsValid() bool {
	_, ok := _FlagsMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i Flags) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *Flags) SetFlag(on bool, f ...enums.BitFlag) {
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
func (i Flags) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Flags) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _NodeSignalsValues = []NodeSignals{0, 1, 2}

// NodeSignalsN is the highest valid value
// for type NodeSignals, plus one.
const NodeSignalsN NodeSignals = 3

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _NodeSignalsNoOp() {
	var x [1]struct{}
	_ = x[NodeSignalNil-(0)]
	_ = x[NodeSignalUpdated-(1)]
	_ = x[NodeSignalDeleting-(2)]
}

var _NodeSignalsNameToValueMap = map[string]NodeSignals{
	`NodeSignalNil`:      0,
	`nodesignalnil`:      0,
	`NodeSignalUpdated`:  1,
	`nodesignalupdated`:  1,
	`NodeSignalDeleting`: 2,
	`nodesignaldeleting`: 2,
}

var _NodeSignalsDescMap = map[NodeSignals]string{
	0: `NodeSignalNil is a nil signal value`,
	1: `NodeSignalUpdated indicates that the node was updated -- the node Flags accumulate the specific changes made since the last update signal -- these flags are sent in the signal data -- strongly recommend using that instead of the flags, which can be subsequently updated by the time a signal is processed`,
	2: `NodeSignalDeleting indicates that the node is being deleted from its parent children list -- this is not blocked by Updating status and is delivered immediately. No further notifications are sent -- assume it will be destroyed unless you hear from it again.`,
}

var _NodeSignalsMap = map[NodeSignals]string{
	0: `NodeSignalNil`,
	1: `NodeSignalUpdated`,
	2: `NodeSignalDeleting`,
}

// String returns the string representation
// of this NodeSignals value.
func (i NodeSignals) String() string {
	if str, ok := _NodeSignalsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the NodeSignals value from its
// string representation, and returns an
// error if the string is invalid.
func (i *NodeSignals) SetString(s string) error {
	if val, ok := _NodeSignalsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _NodeSignalsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type NodeSignals")
}

// Int64 returns the NodeSignals value as an int64.
func (i NodeSignals) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the NodeSignals value from an int64.
func (i *NodeSignals) SetInt64(in int64) {
	*i = NodeSignals(in)
}

// Desc returns the description of the NodeSignals value.
func (i NodeSignals) Desc() string {
	if str, ok := _NodeSignalsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// NodeSignalsValues returns all possible values
// for the type NodeSignals.
func NodeSignalsValues() []NodeSignals {
	return _NodeSignalsValues
}

// Values returns all possible values
// for the type NodeSignals.
func (i NodeSignals) Values() []enums.Enum {
	res := make([]enums.Enum, len(_NodeSignalsValues))
	for i, d := range _NodeSignalsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type NodeSignals.
func (i NodeSignals) IsValid() bool {
	_, ok := _NodeSignalsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i NodeSignals) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *NodeSignals) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
