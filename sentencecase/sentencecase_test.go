// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sentencecase

import (
	"testing"
)

func TestSentenceCase(t *testing.T) {
	AddProperNouns("Google")
	src := "ThisIsAStringInSentenceCaseThatIWroteInTheUSAWithTheHelpOfGoogle"
	want := "This is a string in sentence case that I wrote in the USA with the help of Google"
	have := Of(src)
	if have != want {
		t.Errorf("sentence case of \n%s\nwas\n%s\nbut wanted\n%s", src, have, want)
	}
}
