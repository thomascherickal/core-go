// Code generated by "core generate -webcore content"; DO NOT EDIT.

package main

import (
	"errors"
	"fmt"
	"maps"
	"strings"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/units"
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
		gi.NewButton(parent).SetText("Download")
	},
	"widgets/buttons-1": func(parent gi.Widget) {
		gi.NewButton(parent).SetIcon(icons.Download)
	},
	"widgets/buttons-2": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Download").SetIcon(icons.Download)
	},
	"widgets/buttons-3": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Send").SetIcon(icons.Send).OnClick(func(e events.Event) {
			gi.MessageSnackbar(parent, "Message sent")
		})
	},
	"widgets/buttons-4": func(parent gi.Widget) {
		gi.NewButton(parent).SetText("Share").SetIcon(icons.Share).SetMenu(func(m *gi.Scene) {
			gi.NewButton(m).SetText("Copy link")
			gi.NewButton(m).SetText("Send message")
		})
	},
	"widgets/buttons-5": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonFilled).SetText("Filled")
	},
	"widgets/buttons-6": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonTonal).SetText("Tonal")
	},
	"widgets/buttons-7": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonElevated).SetText("Elevated")
	},
	"widgets/buttons-8": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonOutlined).SetText("Outlined")
	},
	"widgets/buttons-9": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonText).SetText("Text")
	},
	"widgets/buttons-10": func(parent gi.Widget) {
		gi.NewButton(parent).SetType(gi.ButtonAction).SetText("Action")
	},
	"widgets/choosers-0": func(parent gi.Widget) {
		gi.NewChooser(parent).SetStrings("macOS", "Windows", "Linux")
	},
	"widgets/choosers-1": func(parent gi.Widget) {
		gi.NewChooser(parent).SetItems(
			gi.ChooserItem{Value: "Computer", Icon: icons.Computer, Tooltip: "Use a computer"},
			gi.ChooserItem{Value: "Phone", Icon: icons.Smartphone, Tooltip: "Use a phone"},
		)
	},
	"widgets/choosers-2": func(parent gi.Widget) {
		gi.NewChooser(parent).SetPlaceholder("Choose a platform").SetStrings("macOS", "Windows", "Linux")
	},
	"widgets/choosers-3": func(parent gi.Widget) {
		gi.NewChooser(parent).SetStrings("Apple", "Orange", "Strawberry").SetCurrentValue("Orange")
	},
	"widgets/choosers-4": func(parent gi.Widget) {
		gi.NewChooser(parent).SetType(gi.ChooserOutlined).SetStrings("Apple", "Orange", "Strawberry")
	},
	"widgets/choosers-5": func(parent gi.Widget) {
		gi.NewChooser(parent).SetIcon(icons.Sort).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-6": func(parent gi.Widget) {
		gi.NewChooser(parent).SetEditable(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-7": func(parent gi.Widget) {
		gi.NewChooser(parent).SetAllowNew(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-8": func(parent gi.Widget) {
		gi.NewChooser(parent).SetEditable(true).SetAllowNew(true).SetStrings("Newest", "Oldest", "Popular")
	},
	"widgets/choosers-9": func(parent gi.Widget) {
		ch := gi.NewChooser(parent).SetStrings("Newest", "Oldest", "Popular")
		ch.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("Sorting by %v", ch.CurrentItem.Value))
		})
	},
	"widgets/dialogs-0": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Message")
		bt.OnClick(func(e events.Event) {
			gi.MessageDialog(bt, "Something happened", "Message")
		})
	},
	"widgets/dialogs-1": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Error")
		bt.OnClick(func(e events.Event) {
			gi.ErrorDialog(bt, errors.New("invalid encoding format"), "Error loading file")
		})
	},
	"widgets/dialogs-2": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Confirm")
		bt.OnClick(func(e events.Event) {
			d := gi.NewBody().AddTitle("Confirm").AddText("Send message?")
			d.AddBottomBar(func(pw gi.Widget) {
				d.AddCancel(pw).OnClick(func(e events.Event) {
					gi.MessageSnackbar(bt, "Dialog canceled")
				})
				d.AddOk(pw).OnClick(func(e events.Event) {
					gi.MessageSnackbar(bt, "Dialog accepted")
				})
			})
			d.NewDialog(bt).Run()
		})
	},
	"widgets/frames-0": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-1": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Background = colors.C(colors.Scheme.Warn.Container)
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-2": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.C(colors.Scheme.Outline))
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-3": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Border.Radius = styles.BorderRadiusLarge
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.C(colors.Scheme.Outline))
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/frames-4": func(parent gi.Widget) {
		fr := gi.NewFrame(parent)
		fr.Style(func(s *styles.Style) {
			s.Grow.Set(0, 0)
			s.Border.Width.Set(units.Dp(4))
			s.Border.Color.Set(colors.C(colors.Scheme.Outline))
		})
		gi.NewButton(fr).SetText("First")
		gi.NewButton(fr).SetText("Second")
		gi.NewButton(fr).SetText("Third")
	},
	"widgets/labels-0": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Hello, world!")
	},
	"widgets/labels-1": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("This is a very long sentence that demonstrates how label content will overflow onto multiple lines when the size of the label text exceeds the size of its surrounding container; labels are a customizable widget that Cogent Core provides, allowing you to display many kinds of text")
	},
	"widgets/labels-2": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText(`<b>You</b> can use <i>HTML</i> <u>formatting</u> inside of <b><i><u>Cogent Core</u></i></b> labels, including <span style="color:red;background-color:yellow">custom styling</span> and <a href="https://example.com">links</a>`)
	},
	"widgets/labels-3": func(parent gi.Widget) {
		gi.NewLabel(parent).SetType(gi.LabelHeadlineMedium).SetText("Hello, world!")
	},
	"widgets/labels-4": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Hello,\n\tworld!").Style(func(s *styles.Style) {
			s.Font.Size.Dp(21)
			s.Font.Style = styles.Italic
			s.Text.WhiteSpace = styles.WhiteSpacePre
			s.Color = colors.C(colors.Scheme.Success.Base)
			s.Font.Family = string(gi.AppearanceSettings.MonoFont)
		})
	},
	"widgets/layouts-0": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-1": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Direction = styles.Column
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-2": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Gap.Set(units.Em(2))
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-3": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Max.X.Em(10)
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-4": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Overflow.X = styles.OverflowAuto
			s.Max.X.Em(10)
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-5": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Wrap = true
			s.Max.X.Em(10)
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
	},
	"widgets/layouts-6": func(parent gi.Widget) {
		ly := gi.NewLayout(parent)
		ly.Style(func(s *styles.Style) {
			s.Display = styles.Grid
			s.Columns = 2
		})
		gi.NewButton(ly).SetText("First")
		gi.NewButton(ly).SetText("Second")
		gi.NewButton(ly).SetText("Third")
		gi.NewButton(ly).SetText("Fourth")
	},
	"widgets/meters-0": func(parent gi.Widget) {
		gi.NewMeter(parent)
	},
	"widgets/meters-1": func(parent gi.Widget) {
		gi.NewMeter(parent).SetValue(0.7)
	},
	"widgets/meters-2": func(parent gi.Widget) {
		gi.NewMeter(parent).SetMin(5.7).SetMax(18).SetValue(10.2)
	},
	"widgets/meters-3": func(parent gi.Widget) {
		gi.NewMeter(parent).Style(func(s *styles.Style) {
			s.Direction = styles.Column
		})
	},
	"widgets/meters-4": func(parent gi.Widget) {
		gi.NewMeter(parent).SetType(gi.MeterCircle)
	},
	"widgets/meters-5": func(parent gi.Widget) {
		gi.NewMeter(parent).SetType(gi.MeterSemicircle)
	},
	"widgets/meters-6": func(parent gi.Widget) {
		gi.NewMeter(parent).SetType(gi.MeterCircle).SetText("50%")
	},
	"widgets/meters-7": func(parent gi.Widget) {
		gi.NewMeter(parent).SetType(gi.MeterSemicircle).SetText("50%")
	},
	"widgets/sliders-0": func(parent gi.Widget) {
		gi.NewSlider(parent)
	},
	"widgets/sliders-1": func(parent gi.Widget) {
		gi.NewSlider(parent).SetValue(0.7)
	},
	"widgets/sliders-2": func(parent gi.Widget) {
		gi.NewSlider(parent).SetMin(5.7).SetMax(18).SetValue(10.2)
	},
	"widgets/sliders-3": func(parent gi.Widget) {
		gi.NewSlider(parent).SetStep(0.2)
	},
	"widgets/sliders-4": func(parent gi.Widget) {
		gi.NewSlider(parent).SetStep(0.2).SetEnforceStep(true)
	},
	"widgets/sliders-5": func(parent gi.Widget) {
		gi.NewSlider(parent).SetIcon(icons.DeployedCode.Fill())
	},
	"widgets/sliders-6": func(parent gi.Widget) {
		sr := gi.NewSlider(parent)
		sr.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("OnChange: %v", sr.Value))
		})
	},
	"widgets/sliders-7": func(parent gi.Widget) {
		sr := gi.NewSlider(parent)
		sr.OnInput(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("OnInput: %v", sr.Value))
		})
	},
	"widgets/snackbars-0": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Message")
		bt.OnClick(func(e events.Event) {
			gi.MessageSnackbar(bt, "New messages loaded")
		})
	},
	"widgets/snackbars-1": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Error")
		bt.OnClick(func(e events.Event) {
			gi.ErrorSnackbar(bt, errors.New("file not found"), "Error loading page")
		})
	},
	"widgets/snackbars-2": func(parent gi.Widget) {
		bt := gi.NewButton(parent).SetText("Custom")
		bt.OnClick(func(e events.Event) {
			gi.NewBody().AddSnackbarText("Files updated").
				AddSnackbarButton("Refresh", func(e events.Event) {
					gi.MessageSnackbar(bt, "Refreshed files")
				}).AddSnackbarIcon(icons.Close).NewSnackbar(bt).Run()
		})
	},
	"widgets/spinners-0": func(parent gi.Widget) {
		gi.NewSpinner(parent)
	},
	"widgets/spinners-1": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetValue(12.7)
	},
	"widgets/spinners-2": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetMin(-0.5).SetMax(2.7)
	},
	"widgets/spinners-3": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetStep(6)
	},
	"widgets/spinners-4": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetStep(4).SetEnforceStep(true)
	},
	"widgets/spinners-5": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetType(gi.TextFieldOutlined)
	},
	"widgets/spinners-6": func(parent gi.Widget) {
		gi.NewSpinner(parent).SetFormat("%X").SetStep(1).SetValue(44)
	},
	"widgets/spinners-7": func(parent gi.Widget) {
		sp := gi.NewSpinner(parent)
		sp.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, fmt.Sprintf("Value changed to %g", sp.Value))
		})
	},
	"widgets/splits-0": func(parent gi.Widget) {
		sp := gi.NewSplits(parent)
		gi.NewLabel(sp).SetText("First")
		gi.NewLabel(sp).SetText("Second")
	},
	"widgets/splits-1": func(parent gi.Widget) {
		sp := gi.NewSplits(parent)
		gi.NewLabel(sp).SetText("First")
		gi.NewLabel(sp).SetText("Second")
		gi.NewLabel(sp).SetText("Third")
		gi.NewLabel(sp).SetText("Fourth")
	},
	"widgets/switches-0": func(parent gi.Widget) {
		gi.NewSwitch(parent)
	},
	"widgets/switches-1": func(parent gi.Widget) {
		gi.NewSwitch(parent).SetText("Remember me")
	},
	"widgets/switches-2": func(parent gi.Widget) {
		gi.NewSwitch(parent).SetType(gi.SwitchCheckbox).SetText("Remember me")
	},
	"widgets/switches-3": func(parent gi.Widget) {
		gi.NewSwitch(parent).SetType(gi.SwitchRadioButton).SetText("Remember me")
	},
	"widgets/switches-4": func(parent gi.Widget) {
		sw := gi.NewSwitch(parent).SetText("Remember me")
		sw.OnChange(func(e events.Event) {
			gi.MessageSnackbar(sw, fmt.Sprintf("Switch is %v", sw.IsChecked()))
		})
	},
	"widgets/switches-5": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-6": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetItems(
			gi.SwitchItem{Label: "Go", Tooltip: "Elegant, fast, and easy-to-use"},
			gi.SwitchItem{Label: "Python", Tooltip: "Slow and duck-typed"},
			gi.SwitchItem{Label: "C++", Tooltip: "Hard to use and slow to compile"},
		)
	},
	"widgets/switches-7": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetMutex(true).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-8": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetType(gi.SwitchChip).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-9": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetType(gi.SwitchCheckbox).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-10": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetType(gi.SwitchRadioButton).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-11": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetType(gi.SwitchSegmentedButton).SetStrings("Go", "Python", "C++")
	},
	"widgets/switches-12": func(parent gi.Widget) {
		gi.NewSwitches(parent).SetStrings("Go", "Python", "C++").Style(func(s *styles.Style) {
			s.Direction = styles.Column
		})
	},
	"widgets/switches-13": func(parent gi.Widget) {
		sw := gi.NewSwitches(parent).SetStrings("Go", "Python", "C++")
		sw.OnChange(func(e events.Event) {
			gi.MessageSnackbar(sw, fmt.Sprintf("Currently selected: %v", sw.SelectedItems()))
		})
	},
	"widgets/text-fields-0": func(parent gi.Widget) {
		gi.NewTextField(parent)
	},
	"widgets/text-fields-1": func(parent gi.Widget) {
		gi.NewLabel(parent).SetText("Name:")
		gi.NewTextField(parent).SetPlaceholder("Jane Doe")
	},
	"widgets/text-fields-2": func(parent gi.Widget) {
		gi.NewTextField(parent).SetText("Hello, world!")
	},
	"widgets/text-fields-3": func(parent gi.Widget) {
		gi.NewTextField(parent).SetText("This is a long sentence that demonstrates how text field content can overflow onto multiple lines")
	},
	"widgets/text-fields-4": func(parent gi.Widget) {
		gi.NewTextField(parent).SetType(gi.TextFieldOutlined)
	},
	"widgets/text-fields-5": func(parent gi.Widget) {
		gi.NewTextField(parent).SetTypePassword()
	},
	"widgets/text-fields-6": func(parent gi.Widget) {
		gi.NewTextField(parent).AddClearButton()
	},
	"widgets/text-fields-7": func(parent gi.Widget) {
		gi.NewTextField(parent).SetLeadingIcon(icons.Euro).SetTrailingIcon(icons.OpenInNew, func(e events.Event) {
			gi.MessageSnackbar(parent, "Opening shopping cart")
		})
	},
	"widgets/text-fields-8": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.SetValidator(func() error {
			if !strings.Contains(tf.Text(), "Go") {
				return errors.New("Must contain Go")
			}
			return nil
		})
	},
	"widgets/text-fields-9": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.OnChange(func(e events.Event) {
			gi.MessageSnackbar(parent, "OnChange: "+tf.Text())
		})
	},
	"widgets/text-fields-10": func(parent gi.Widget) {
		tf := gi.NewTextField(parent)
		tf.OnInput(func(e events.Event) {
			gi.MessageSnackbar(parent, "OnInput: "+tf.Text())
		})
	},
}
