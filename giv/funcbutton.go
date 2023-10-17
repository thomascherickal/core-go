// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package giv

import (
	"log/slog"
	"strings"

	"github.com/iancoleman/strcase"
	"goki.dev/gi/v2/gi"
	"goki.dev/glop/sentencecase"
	"goki.dev/gti"
	"goki.dev/icons"
	"goki.dev/ki/v2"
)

// NewFuncButton adds to the given parent a new [gi.Button] that is set up to call the given
// function when pressed, using a dialog to prompt the user for any arguments. Also, it sets
// various properties like the name, label, tooltip, and icon of the button based on the
// properties of the function, using reflect and gti. The given function must be registered
// with gti; add a `//gti:add` comment directive and run `goki generate` if you get errors.
func NewFuncButton(par ki.Ki, fun any) *gi.Button {
	fnm := gti.FuncName(fun)
	f := gti.FuncByName(fnm)
	if f == nil {
		slog.Error("programmer error: cannot use giv.NewFuncButton with a function that has not been added to gti", "function", fnm)
		return nil
	}
	// get name without package
	snm := f.Name
	li := strings.LastIndex(snm, ".")
	if li >= 0 {
		snm = snm[li+1:] // must also get rid of "."
	}
	bt := gi.NewButton(par, snm).SetText(sentencecase.Of(snm))
	bt.SetTooltip(f.Doc)
	// we default to the icon with the same name as
	// the function, if it exists
	ic := icons.Icon(strcase.ToSnake(snm))
	if ic.IsValid() {
		bt.SetIcon(ic)
	}
	return bt
}
