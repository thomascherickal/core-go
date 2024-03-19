// Code generated by "core generate -webcore content"; DO NOT EDIT.

package main

import (
	"fmt"
	"maps"

	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/webcore"
)

func init() {
	maps.Copy(webcore.Examples, WebcoreExamples)
}

// WebcoreExamples are the compiled webcore examples for this app.
var WebcoreExamples = map[string]func(parent gi.Widget){
	"getting-started/hello-world-0": func(parent gi.Widget) {
		b := parent
		gi.NewButton(b).SetText("Hello, World!")
	},
	"basics/widgets-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").SetIcon(icons.Add)
	},
	"basics/widgets-1": func(parent gi.Widget) {
		sw := gi.NewSwitch(parent).SetText("Switch me!")
		// Later...
		gi.MessageSnackbar(parent, sw.Text)
	},
	"basics/events-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Button clicked")
		})
	},
	"basics/events-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Click me!").OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprint("Button clicked at ", e.Pos()))
			e.SetHandled() // this event will not be handled by other event handlers now
		})
	},
	"basics/icons-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Send").SetIcon(icons.Send).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Message sent")
		})
	},
	"widgets/buttons-0": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Send").SetIcon(icons.Send).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Message sent")
		})
	},
}
