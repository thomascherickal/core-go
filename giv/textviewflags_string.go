// Code generated by "stringer -type=TextViewFlags"; DO NOT EDIT.

package giv

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TextViewNeedsRefresh-29]
	_ = x[TextViewInReLayout-30]
	_ = x[TextViewRenderScrolls-31]
	_ = x[TextViewFocusActive-32]
	_ = x[TextViewHasLineNos-33]
	_ = x[TextViewLastWasTabAI-34]
	_ = x[TextViewLastWasUndo-35]
	_ = x[TextViewFlagsN-36]
}

const _TextViewFlags_name = "TextViewNeedsRefreshTextViewInReLayoutTextViewRenderScrollsTextViewFocusActiveTextViewHasLineNosTextViewLastWasTabAITextViewLastWasUndoTextViewFlagsN"

var _TextViewFlags_index = [...]uint8{0, 20, 38, 59, 78, 96, 116, 135, 149}

func (i TextViewFlags) String() string {
	i -= 29
	if i < 0 || i >= TextViewFlags(len(_TextViewFlags_index)-1) {
		return "TextViewFlags(" + strconv.FormatInt(int64(i+29), 10) + ")"
	}
	return _TextViewFlags_name[_TextViewFlags_index[i]:_TextViewFlags_index[i+1]]
}
