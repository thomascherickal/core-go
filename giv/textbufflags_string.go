// Code generated by "stringer -type=TextBufFlags"; DO NOT EDIT.

package giv

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextBufAutoSaving-29]
	_ = x[TextBufMarkingUp-30]
	_ = x[TextBufChanged-31]
	_ = x[TextBufFileModOk-32]
	_ = x[TextBufFlagsN-33]
}

const _TextBufFlags_name = "TextBufAutoSavingTextBufMarkingUpTextBufChangedTextBufFileModOkTextBufFlagsN"

var _TextBufFlags_index = [...]uint8{0, 17, 33, 47, 63, 76}

func (i TextBufFlags) String() string {
	i -= 29
	if i < 0 || i >= TextBufFlags(len(_TextBufFlags_index)-1) {
		return "TextBufFlags(" + strconv.FormatInt(int64(i+29), 10) + ")"
	}
	return _TextBufFlags_name[_TextBufFlags_index[i]:_TextBufFlags_index[i+1]]
}
