// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colors

import (
	"image/color"

	"goki.dev/colors/matcolor"
)

// Palette contains the main, global [MatPalette]. It can
// be used by end-user code for accessing tonal palette values,
// although [Scheme] is a more typical way to access the color
// scheme values. It defaults to a palette based around a
// primary color of Google Blue (#4285f4)
var Palette = matcolor.NewPalette(matcolor.KeyFromPrimary(color.RGBA{66, 133, 244, 255}))

// Schemes are the main global Material Design 3 color schemes.
// They should not be used for accessing the current color scheme;
// see [Scheme] for that. Instead, they should be set if you want
// to define your own custom color scheme for your app.
var Schemes = matcolor.NewSchemes(Palette)

// Scheme is the main currently active global Material Design 3
// color scheme. It is the main way that end-user code should
// access the color scheme; ideally, almost all color values should
// be set to something in here. For more specific tones of colors,
// see [Palette]. For setting the color scheme of your app, see
// [Schemes].
var Scheme = &Schemes.Light
