// Code generated by "goki generate ./..."; DO NOT EDIT.

package base

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goosi/driver/base.App",
	ShortName: "base.App",
	IDName:    "app",
	Doc:       "App contains the data and logic common to all implementations of [goosi.App].",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"This", &gti.Field{Name: "This", Type: "goki.dev/goosi.App", LocalType: "goosi.App", Doc: "This is the App as a [goosi.App] interface, which preserves the actual identity\nof the app when calling interface methods in the base App.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"Mu", &gti.Field{Name: "Mu", Type: "sync.Mutex", LocalType: "sync.Mutex", Doc: "Mu is the main mutex protecting access to app operations, including [App.RunOnMain] functions.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"MainQueue", &gti.Field{Name: "MainQueue", Type: "chan goki.dev/goosi/driver/base.FuncRun", LocalType: "chan FuncRun", Doc: "MainQueue is the queue of functions to call on the main loop. To add to it, use [App.RunOnMain].", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"MainDone", &gti.Field{Name: "MainDone", Type: "chan struct{}", LocalType: "chan struct{}", Doc: "MainDone is a channel on which is a signal is sent when the main loop of the app should be terminated.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"Nm", &gti.Field{Name: "Nm", Type: "string", LocalType: "string", Doc: "Nm is the name of the app.", Directives: gti.Directives{}, Tag: "label:\"Name\""}},
		{"Abt", &gti.Field{Name: "Abt", Type: "string", LocalType: "string", Doc: "Abt is the about information for the app.", Directives: gti.Directives{}, Tag: "label:\"About\""}},
		{"OpenFls", &gti.Field{Name: "OpenFls", Type: "[]string", LocalType: "[]string", Doc: "OpenFls are files that have been set by the operating system to open at startup.", Directives: gti.Directives{}, Tag: "label:\"Open files\""}},
		{"Quitting", &gti.Field{Name: "Quitting", Type: "bool", LocalType: "bool", Doc: "Quitting is whether the app is quitting and thus closing all of the windows", Directives: gti.Directives{}, Tag: ""}},
		{"QuitReqFunc", &gti.Field{Name: "QuitReqFunc", Type: "func()", LocalType: "func()", Doc: "QuitReqFunc is a function to call when a quit is requested", Directives: gti.Directives{}, Tag: ""}},
		{"QuitCleanFunc", &gti.Field{Name: "QuitCleanFunc", Type: "func()", LocalType: "func()", Doc: "QuitCleanFunc is a function to call when the app is about to quit", Directives: gti.Directives{}, Tag: ""}},
		{"Dark", &gti.Field{Name: "Dark", Type: "bool", LocalType: "bool", Doc: "Dark is whether the system color theme is dark (as opposed to light)", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goosi/driver/base.AppMulti",
	ShortName: "base.AppMulti",
	IDName:    "app-multi",
	Doc:       "AppMulti contains the data and logic common to all implementations of [goosi.App]\non multi-window platforms (desktop), as opposed to single-window\nplatforms (mobile, web, and offscreen), for which you should use [AppSingle]. An AppMulti is associated\nwith a corresponding type of [goosi.Window]. The [goosi.Window]\ntype should embed [WindowMulti].",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Windows", &gti.Field{Name: "Windows", Type: "[]W", LocalType: "[]W", Doc: "Windows are the windows associated with the app", Directives: gti.Directives{}, Tag: ""}},
		{"Screens", &gti.Field{Name: "Screens", Type: "[]*goki.dev/goosi.Screen", LocalType: "[]*goosi.Screen", Doc: "Screens are the screens associated with the app", Directives: gti.Directives{}, Tag: ""}},
		{"AllScreens", &gti.Field{Name: "AllScreens", Type: "[]*goki.dev/goosi.Screen", LocalType: "[]*goosi.Screen", Doc: "AllScreens is a unique list of all screens ever seen, from which\ninformation can be got if something is missing in [AppMulti.Screens]", Directives: gti.Directives{}, Tag: ""}},
		{"CtxWindow", &gti.Field{Name: "CtxWindow", Type: "W", LocalType: "W", Doc: "CtxWindow is a dynamically set context window used for some operations", Directives: gti.Directives{}, Tag: "label:\"Context window\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"App", &gti.Field{Name: "App", Type: "goki.dev/goosi/driver/base.App", LocalType: "App", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goosi/driver/base.AppSingle",
	ShortName: "base.AppSingle",
	IDName:    "app-single",
	Doc:       "AppSingle contains the data and logic common to all implementations of [goosi.App]\non single-window platforms (mobile, web, and offscreen), as opposed to multi-window\nplatforms (desktop), for which you should use [AppMulti]. An AppSingle is associated\nwith a corresponding type of [goosi.Drawer] and [goosi.Window]. The [goosi.Window]\ntype should embed [WindowSingle].",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Drawer", &gti.Field{Name: "Drawer", Type: "D", LocalType: "D", Doc: "Drawer is the single [goosi.Drawer] used for the app.", Directives: gti.Directives{}, Tag: ""}},
		{"Win", &gti.Field{Name: "Win", Type: "W", LocalType: "W", Doc: "Win is the single [goosi.Window] associated with the app.", Directives: gti.Directives{}, Tag: "label:\"Window\""}},
		{"Scrn", &gti.Field{Name: "Scrn", Type: "*goki.dev/goosi.Screen", LocalType: "*goosi.Screen", Doc: "Scrn is the single [goosi.Screen] associated with the app.", Directives: gti.Directives{}, Tag: "label:\"Screen\""}},
		{"Insts", &gti.Field{Name: "Insts", Type: "goki.dev/girl/styles.SideFloats", LocalType: "styles.SideFloats", Doc: "Insts are the size of any insets on the sides of the screen.", Directives: gti.Directives{}, Tag: "label:\"Insets\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"App", &gti.Field{Name: "App", Type: "goki.dev/goosi/driver/base.App", LocalType: "App", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goosi/driver/base.Window",
	ShortName: "base.Window",
	IDName:    "window",
	Doc:       "Window contains the data and logic common to all implementations of [goosi.Window].\nA Window is associated with a corresponding [goosi.App] type.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"This", &gti.Field{Name: "This", Type: "goki.dev/goosi.Window", LocalType: "goosi.Window", Doc: "This is the Window as a [goosi.Window] interface, which preserves the actual identity\nof the window when calling interface methods in the base Window.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"App", &gti.Field{Name: "App", Type: "A", LocalType: "A", Doc: "App is the [goosi.App] associated with the window.", Directives: gti.Directives{}, Tag: ""}},
		{"Mu", &gti.Field{Name: "Mu", Type: "sync.Mutex", LocalType: "sync.Mutex", Doc: "Mu is the main mutex protecting access to window operations, including [Window.RunOnWin] functions.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"RunQueue", &gti.Field{Name: "RunQueue", Type: "chan goki.dev/goosi/driver/base.FuncRun", LocalType: "chan FuncRun", Doc: "RunQueue is the queue of functions to call on the window loop. To add to it, use [Window.RunOnWin].", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"WinClose", &gti.Field{Name: "WinClose", Type: "chan struct{}", LocalType: "chan struct{}", Doc: "WinClose is a channel on which a single is sent to indicate that the\nwindow should close.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"CloseReqFunc", &gti.Field{Name: "CloseReqFunc", Type: "func(win goki.dev/goosi.Window)", LocalType: "func(win goosi.Window)", Doc: "CloseReqFunc is the function to call on a close request", Directives: gti.Directives{}, Tag: ""}},
		{"CloseCleanFunc", &gti.Field{Name: "CloseCleanFunc", Type: "func(win goki.dev/goosi.Window)", LocalType: "func(win goosi.Window)", Doc: "CloseCleanFunc is the function to call to close the window", Directives: gti.Directives{}, Tag: ""}},
		{"Nm", &gti.Field{Name: "Nm", Type: "string", LocalType: "string", Doc: "Nm is the name of the window", Directives: gti.Directives{}, Tag: "label:\"Name\""}},
		{"Titl", &gti.Field{Name: "Titl", Type: "string", LocalType: "string", Doc: "Titl is the title of the window", Directives: gti.Directives{}, Tag: "label:\"Title\""}},
		{"Flgs", &gti.Field{Name: "Flgs", Type: "goki.dev/goosi.WindowFlags", LocalType: "goosi.WindowFlags", Doc: "Flgs contains the flags associated with the window", Directives: gti.Directives{}, Tag: "label:\"Flags\""}},
		{"FPS", &gti.Field{Name: "FPS", Type: "int", LocalType: "int", Doc: "FPS is the FPS (frames per second) for rendering the window", Directives: gti.Directives{}, Tag: ""}},
		{"EvMgr", &gti.Field{Name: "EvMgr", Type: "goki.dev/goosi/events.Mgr", LocalType: "events.Mgr", Doc: "EvMgr is the event manager for the window", Directives: gti.Directives{}, Tag: "label:\"Event manger\""}},
		{"DestroyGPUFunc", &gti.Field{Name: "DestroyGPUFunc", Type: "func()", LocalType: "func()", Doc: "DestroyGPUFunc should be set to a function that will destroy GPU resources\nin the main thread prior to destroying the drawer\nand the surface; otherwise it is difficult to\nensure that the proper ordering of destruction applies.", Directives: gti.Directives{}, Tag: ""}},
		{"CursorEnabled", &gti.Field{Name: "CursorEnabled", Type: "bool", LocalType: "bool", Doc: "CursorEnabled is whether the cursor is currently enabled", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Deque", &gti.Field{Name: "Deque", Type: "goki.dev/goosi/events.Deque", LocalType: "events.Deque", Doc: "", Directives: gti.Directives{}, Tag: "view:\"-\""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goosi/driver/base.WindowMulti",
	ShortName: "base.WindowMulti",
	IDName:    "window-multi",
	Doc:       "WindowMulti contains the data and logic common to all implementations of [goosi.Window]\non multi-window platforms (desktop), as opposed to single-window\nplatforms (mobile, web, and offscreen), for which you should use [WindowSingle].\nA WindowMulti is associated with a corresponding [goosi.App] type.\nThe [goosi.App] type should embed [AppMulti].",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Draw", &gti.Field{Name: "Draw", Type: "D", LocalType: "D", Doc: "Draw is the [goosi.Drawer] used for this window.", Directives: gti.Directives{}, Tag: "label:\"Drawer\""}},
		{"Pos", &gti.Field{Name: "Pos", Type: "image.Point", LocalType: "image.Point", Doc: "Pos is the position of the window", Directives: gti.Directives{}, Tag: "label:\"Position\""}},
		{"WnSize", &gti.Field{Name: "WnSize", Type: "image.Point", LocalType: "image.Point", Doc: "WnSize is the size of the window in window manager coordinates", Directives: gti.Directives{}, Tag: "label:\"Window manager size\""}},
		{"PixSize", &gti.Field{Name: "PixSize", Type: "image.Point", LocalType: "image.Point", Doc: "PixSize is the pixel size of the window in raw display dots", Directives: gti.Directives{}, Tag: "label:\"Pixel size\""}},
		{"DevicePixelRatio", &gti.Field{Name: "DevicePixelRatio", Type: "float32", LocalType: "float32", Doc: "DevicePixelRatio is a factor that scales the screen's\n\"natural\" pixel coordinates into actual device pixels.\nOn OS-X, it is backingScaleFactor = 2.0 on \"retina\"", Directives: gti.Directives{}, Tag: ""}},
		{"PhysDPI", &gti.Field{Name: "PhysDPI", Type: "float32", LocalType: "float32", Doc: "PhysicalDPI is the physical dots per inch of the screen,\nfor generating true-to-physical-size output.\nIt is computed as 25.4 * (PixSize.X / PhysicalSize.X)\nwhere 25.4 is the number of mm per inch.", Directives: gti.Directives{}, Tag: "label:\"Physical DPI\""}},
		{"LogDPI", &gti.Field{Name: "LogDPI", Type: "float32", LocalType: "float32", Doc: "LogicalDPI is the logical dots per inch of the screen,\nwhich is used for all rendering.\nIt is: transient zoom factor * screen-specific multiplier * PhysicalDPI", Directives: gti.Directives{}, Tag: "label:\"Logical DPI\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Window", &gti.Field{Name: "Window", Type: "goki.dev/goosi/driver/base.Window", LocalType: "Window[A]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/goosi/driver/base.WindowSingle",
	ShortName: "base.WindowSingle",
	IDName:    "window-single",
	Doc:       "WindowSingle contains the data and logic common to all implementations of [goosi.Window]\non single-window platforms (mobile, web, and offscreen), as opposed to multi-window\nplatforms (desktop), for which you should use [WindowSingle].\nA WindowSingle is associated with a corresponding [AppSingler] type.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Window", &gti.Field{Name: "Window", Type: "goki.dev/goosi/driver/base.Window", LocalType: "Window[A]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
