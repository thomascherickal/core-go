// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package texteditor

/* TODO(config)
import (
	"fmt"

	"cogentcore.org/core/base/reflectx"
	"cogentcore.org/core/events"
	"cogentcore.org/core/views"
)

// Value represents a string with an [Editor].
type Value struct {
	views.ValueBase[*Editor]
}

func (v *Value) Config() {
	tb := NewBuffer()
	tb.OnChange(func(e events.Event) {
		v.SetValue(tb.String())
		fmt.Println(reflectx.UnderlyingPointer(v.Value).Interface())
	})
	v.Widget.SetBuffer(tb)
}

func (v *Value) Update() {
	npv := reflectx.NonPointerValue(v.Value)
	v.Widget.Buffer.SetText([]byte(npv.String()))
}
*/
