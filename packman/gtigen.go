// Code generated by "goki generate ./..."; DO NOT EDIT.

package packman

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.Build",
	Doc:  "Build builds an executable for the package\nat the config path for the config platforms.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.Install",
	Doc:  "Install installs the config package by looking for it in the list\nof supported packages. If the config ID is a filepath, it installs\nthe package at that filepath on the local system. Install uses the\nsame config info as build.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.Log",
	Doc:  "Log prints the logs from your app running on Android to the terminal.\nAndroid is the only supported platform for log; use the -debug flag on\nrun for other platforms.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.Release",
	Doc:  "Release releases the config project\nby calling [ReleaseApp] if it is an app\nand [ReleaseLibrary] if it is a library.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.Run",
	Doc:  "Run builds and runs the config package. It also displays the logs generated\nby the app. It uses the same config info as build.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.Serve",
	Doc:  "Serve builds the package into static web files and then\nserves them on localhost at the config port.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.GetVersion",
	Doc:  "GetVersion prints the version of the project.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.SetVersion",
	Doc:  "SetVersion updates the config and version file of the config project based\non the config version and commits and pushes the changes.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})

var _ = gti.AddFunc(&gti.Func{
	Name: "goki.dev/goki/packman.UpdateVersion",
	Doc:  "UpdateVersion updates the version of the project by one patch version.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"c", &gti.Field{Name: "c", Type: "*goki.dev/goki/config.Config", LocalType: "*config.Config", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
})
