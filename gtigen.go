// Code generated by "goki generate"; DO NOT EDIT.

package ki

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

// NodeType is the [gti.Type] for [Node]
var NodeType = gti.AddType(&gti.Type{
	Name:       "goki.dev/ki/v2.Node",
	Doc:        "The Node implements the Ki interface and provides the core functionality\nfor the GoKi tree -- use the Node as an embedded struct or as a struct\nfield -- the embedded version supports full JSON save / load.\n\nThe desc: key for fields is used by the GoGi GUI viewer for help / tooltip\ninfo -- add these to all your derived struct's fields.  See relevant docs\nfor other such tags controlling a wide range of GUI and other functionality\n-- Ki makes extensive use of such tags.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Nm", &gti.Field{Name: "Nm", Type: "string", Doc: "Ki.Name() user-supplied name of this node -- can be empty or non-unique", Directives: gti.Directives{}}},
		{"Flags", &gti.Field{Name: "Flags", Type: "Flags", Doc: "[tableview: -] bit flags for internal node state -- can extend this using enums package", Directives: gti.Directives{}}},
		{"Props", &gti.Field{Name: "Props", Type: "Props", Doc: "[tableview: -] Ki.Properties() property map for arbitrary extensible properties, including style properties", Directives: gti.Directives{}}},
		{"Par", &gti.Field{Name: "Par", Type: "Ki", Doc: "[view: -] [tableview: -] Ki.Parent() parent of this node -- set automatically when this node is added as a child of parent", Directives: gti.Directives{}}},
		{"Kids", &gti.Field{Name: "Kids", Type: "Slice", Doc: "[tableview: -] Ki.Children() list of children of this node -- all are set to have this node as their parent -- can reorder etc but generally use Ki Node methods to Add / Delete to ensure proper usage", Directives: gti.Directives{}}},
		{"NodeSig", &gti.Field{Name: "NodeSig", Type: "Signal", Doc: "[view: -] Ki.NodeSignal() signal for node structure / state changes -- emits NodeSignals signals -- can also extend to custom signals (see signal.go) but in general better to create a new Signal instead", Directives: gti.Directives{}}},
		{"Ths", &gti.Field{Name: "Ths", Type: "Ki", Doc: "[view: -] we need a pointer to ourselves as a Ki, which can always be used to extract the true underlying type of object when Node is embedded in other structs -- function receivers do not have this ability so this is necessary.  This is set to nil when deleted.  Typically use This() convenience accessor which protects against concurrent access.", Directives: gti.Directives{}}},
		{"index", &gti.Field{Name: "index", Type: "int", Doc: "[view: -] last value of our index -- used as a starting point for finding us in our parent next time -- is not guaranteed to be accurate!  use IndexInParent() method", Directives: gti.Directives{}}},
		{"depth", &gti.Field{Name: "depth", Type: "int", Doc: "[view: -] optional depth parameter of this node -- only valid during specific contexts, not generally -- e.g., used in FuncDownBreadthFirst function", Directives: gti.Directives{}}},
		{"fieldOffs", &gti.Field{Name: "fieldOffs", Type: "[]uintptr", Doc: "[view: -] cached version of the field offsets relative to base Node address -- used in generic field access.", Directives: gti.Directives{}}},
	}),
	Embeds:   ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &Node{},
})

// NewNode adds a new [Node] with
// the given name to the given parent.
func NewNode(par Ki, name string) *Node {
	return par.NewChild(NodeType, name).(*Node)
}

// Type returns the [*gti.Type] of [Node]
func (t *Node) KiType() *gti.Type {
	return NodeType
}

// New returns a new [*Node] value
func (t *Node) New() Ki {
	return &Node{}
}
