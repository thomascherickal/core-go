// Code generated by "stringer -type=VirtualKeyboardTypes"; DO NOT EDIT.

package oswin

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DefaultKeyboard-0]
	_ = x[SingleLineKeyboard-1]
	_ = x[NumberKeyboard-2]
	_ = x[VirtualKeyboardTypesN-3]
}

const _VirtualKeyboardTypes_name = "DefaultKeyboardSingleLineKeyboardNumberKeyboardVirtualKeyboardTypesN"

var _VirtualKeyboardTypes_index = [...]uint8{0, 15, 33, 47, 68}

func (i VirtualKeyboardTypes) String() string {
	if i < 0 || i >= VirtualKeyboardTypes(len(_VirtualKeyboardTypes_index)-1) {
		return "VirtualKeyboardTypes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VirtualKeyboardTypes_name[_VirtualKeyboardTypes_index[i]:_VirtualKeyboardTypes_index[i+1]]
}

func (i *VirtualKeyboardTypes) FromString(s string) error {
	for j := 0; j < len(_VirtualKeyboardTypes_index)-1; j++ {
		if s == _VirtualKeyboardTypes_name[_VirtualKeyboardTypes_index[j]:_VirtualKeyboardTypes_index[j+1]] {
			*i = VirtualKeyboardTypes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: VirtualKeyboardTypes")
}

var _VirtualKeyboardTypes_descMap = map[VirtualKeyboardTypes]string{
	0: `DefaultKeyboard is the keyboard with default input style and &#34;return&#34; return key`,
	1: `SingleLineKeyboard is the keyboard with default input style and &#34;Done&#34; return key`,
	2: `NumberKeyboard is the keyboard with number input style and &#34;Done&#34; return key`,
	3: ``,
}

func (i VirtualKeyboardTypes) Desc() string {
	if str, ok := _VirtualKeyboardTypes_descMap[i]; ok {
		return str
	}
	return "VirtualKeyboardTypes(" + strconv.FormatInt(int64(i), 10) + ")"
}
