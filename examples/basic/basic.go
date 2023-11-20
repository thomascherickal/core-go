// Copyright (c) 2020, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"goki.dev/colors"
	"goki.dev/girl/paint"
	"goki.dev/girl/styles"
	"goki.dev/mat32/v2"
)

func main() {
	paint.FontLibrary.InitFontPaths(paint.FontPaths...)

	imgsz := image.Point{320, 240}
	szrec := image.Rectangle{Max: imgsz}
	img := image.NewRGBA(szrec)

	rs := &paint.State{}
	pc := &paint.Paint{}

	pc.Defaults()               // zeros are not good defaults for paint
	pc.SetUnitContextExt(imgsz) // initialize units

	rs.Init(imgsz.X, imgsz.Y, img)
	rs.PushBounds(szrec)
	rs.Lock()

	// first, draw a frame around the entire image
	pc.StrokeStyle.SetColor(colors.Black)
	pc.FillStyle.SetColor(colors.White)
	pc.StrokeStyle.Width.SetDot(1) // use dots directly to render in literal pixels
	pc.DrawRectangle(rs, 0, 0, float32(imgsz.X), float32(imgsz.Y))
	pc.FillStrokeClear(rs) // actually render path that has been setup

	// next draw a rounded rectangle
	pc.FillStyle.SetColor(nil)
	pc.StrokeStyle.Width.SetDot(10)
	pc.DrawRoundedRectangle(rs, 20, 20, 150, 100, styles.NewSideFloats(6))
	pc.FillStrokeClear(rs)

	// use units-based styling instead of dots:
	pc.StrokeStyle.SetColor(colors.Blue)
	pc.StrokeStyle.Width.Ew(2) // percent of total image (width)
	pc.ToDots()                // convert pct -> dots based on units context
	// fmt.Printf("pct dots: %g\n", pc.StrokeStyle.Width.Dots) // 6.4
	pc.DrawRoundedRectangle(rs, 40, 40, 150, 100, styles.NewSideFloats(6))
	pc.FillStrokeClear(rs)

	// Text rendering
	tsty := &styles.Text{}
	tsty.Defaults()
	fsty := &styles.FontRender{}
	fsty.Defaults()

	// experiment!
	tsty.Align = styles.Center

	txt := &paint.Text{}
	txt.SetHTML("This is <b>HTML</b> formatted <i>text</i>", fsty, tsty, &pc.UnContext, nil)

	// the last, size arg provides constraints for layout to fit within -- uses width mainly
	tsz := txt.LayoutStdLR(tsty, fsty, &pc.UnContext, mat32.Vec2{100, 40})
	fmt.Printf("text size: %v\n", tsz)

	txt.Render(rs, mat32.Vec2{60, 50})

	rs.Unlock()

	file, err := os.Create("image.png")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
