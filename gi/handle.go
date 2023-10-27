// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"goki.dev/colors"
	"goki.dev/cursors"
	"goki.dev/girl/abilities"
	"goki.dev/girl/states"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/goosi/events"
	"goki.dev/mat32/v2"
)

// Handle represents a draggable handle that can be
// used to control the size of an element.
type Handle struct {
	Frame

	// dimension along which the handle slides (opposite of the dimension it is longest on)
	Dim mat32.Dims

	Min float32
	Max float32
	Pos float32
}

func (hl *Handle) OnInit() {
	hl.HandleEvents()
	hl.HandleStyles()
}

func (hl *Handle) HandleStyles() {
	hl.Style(func(s *styles.Style) {
		s.SetAbilities(true, abilities.Focusable, abilities.Hoverable, abilities.Slideable)

		s.Border.Radius = styles.BorderRadiusFull
		s.BackgroundColor.SetSolid(colors.Scheme.OutlineVariant)

		if hl.Dim == mat32.X {
			s.SetFixedWidth(units.Dp(4))
			s.SetFixedHeight(units.Em(3))
		} else {
			s.SetFixedWidth(units.Em(3))
			s.SetFixedHeight(units.Dp(4))
		}

		if !hl.IsReadOnly() {
			s.Cursor = cursors.Grab
			switch {
			case s.Is(states.Sliding):
				s.Cursor = cursors.Grabbing
			case s.Is(states.Active):
				s.Cursor = cursors.Grabbing
			}
		}
	})
}

func (hl *Handle) HandleEvents() {
	hl.On(events.SlideMove, func(e events.Event) {
		hl.Pos = mat32.NewVec2FmPoint(e.Pos()).Dim(hl.Dim)
		hl.Send(events.Change, e)
	})
}

func (hl *Handle) Value() float32 {
	return hl.Pos / (hl.Max - hl.Min)
}
