// Copyright (c) 2019, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi3d

import (
	"fmt"
	"log"

	"github.com/goki/gi"
	"github.com/goki/gi/mat32"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/oswin/gpu"
	"github.com/goki/ki/ints"
)

// todo: need to be able to trigger update based just on updating colors (e.g. netview render)

// MeshName is a mesh name -- provides an automatic gui chooser for meshes
type MeshName string

// Mesh holds the mesh-based shape used for rendering an object.
// Only indexed triangle meshes are supported.
// All Mesh's must define Vtx, Norm, TexUV (stored interleaved), and Idx components.
// Per-vertex Color is optional, and is appended to the vertex buffer non-interleaved if present.
type Mesh interface {
	// AsMeshBase returns the MeshBase for this Mesh
	AsMeshBase() *MeshBase

	// Reset resets all of the vector and index data for this mesh (to start fresh)
	Reset()

	// Make makes the shape mesh (defined for specific shape types)
	Make()

	// ComputeNorms automatically computes the normals from existing vertex data
	ComputeNorms()

	// AddPlane adds everything to render a plane with the given parameters.
	// waxis, haxis = width, height axis, wdir, hdir are the directions for width and height dimensions.
	// wsegs, hsegs = number of segments to create in each dimension -- more finely subdividing a plane
	// allows for higher-quality lighting and texture rendering (minimum of 1 will be enforced).
	// offset is the distance to place the plane along the orthogonal axis.
	// if clr is non-Nil then it will be added
	AddPlane(waxis, haxis mat32.Components, wdir, hdir int, width, height, offset float32, wsegs, hsegs int, clr gi.Color)

	// Validate checks if all the vertex data is valid
	// any errors are logged
	Validate() error

	// MakeVectors compiles the existing mesh data into the Vectors for GPU rendering
	// Must be called with relevant context active
	MakeVectors(sc *Scene) error

	// Activate activates the mesh Vectors on the GPU
	// Must be called with relevant context active
	Activate(sc *Scene)

	// TransferAll transfer all buffer data to GPU (vectors and indexes)
	// Activate must have just been called
	TransferAll()

	// TransferVectors transfer vectors buffer data to GPU (if vector data has changed)
	// Activate must have just been called
	TransferVectors()

	// TransferIndexes transfer vectors buffer data to GPU (if index data has changed)
	// Activate must have just been called
	TransferIndexes()
}

// MeshBase provides the core implementation of Mesh interface
type MeshBase struct {
	Name    string
	Dynamic bool           `desc:"if true, this mesh changes frequently -- otherwise considered to be static"`
	Vtx     mat32.ArrayF32 `desc:"verticies for triangle shapes that make up the mesh -- all mesh structures must use indexed triangle meshes"`
	Norm    mat32.ArrayF32 `desc:"computed normals for each vertex"`
	TexUV   mat32.ArrayF32 `desc:"texture U,V coordinates for mapping textures onto vertexes"`
	Idx     mat32.ArrayU32 `desc:"indexes that sequentially in groups of 3 define the actual triangle faces"`
	Color   mat32.ArrayF32 `desc:"if per-vertex color material type is used for this mesh, then these are the per-vertex colors -- may not be defined in which case per-vertex materials are not possible for such meshes"`
	Buff    gpu.BufferMgr  `view:"-" desc:"buffer holding computed verticies, normals, indices, etc for rendering"`
}

func (ms *MeshBase) Reset() {
	ms.Vtx = nil
	ms.Norm = nil
	ms.TexUV = nil
	ms.Idx = nil
	ms.Color = nil
}

// Validate checks if all the vertex data is valid
// any errors are logged
func (ms *MeshBase) Validate() error {
	vln := len(ms.Vtx)
	if vln == 0 {
		err := fmt.Errorf("gi3d.Mesh: %v has no verticies", ms.Name)
		log.Println(err)
		return err
	}
	nln := len(ms.Norm)
	if nln != vln {
		err := vmt.Errorf("gi3d.Mesh: %v number of Norms: %d != Vtx: %d", ms.Name, nln, vln)
		log.Println(err)
		return err
	}
	tln := len(ms.TexUV)
	if tln != vln {
		err := fmt.Errorf("gi3d.Mesh: %v number of TexUV: %d != Vtx: %d", ms.Name, tln, vln)
		log.Println(err)
		return err
	}
	cln := len(ms.Color)
	if clr == 0 {
		return nil
	}
	if cln != vln {
		err := fmt.Errorf("gi3d.Mesh: %v number of Colors: %d != Vtx: %d", ms.Name, cln, vln)
		log.Println(err)
		return err
	}
	return nil
}

// MakeVectors compiles the existing mesh data into the Vectors for GPU rendering
// Must be called with relevant context active
func (ms *MeshBase) MakeVectors(sc *Scene) error {
	err := ms.Validate()
	if err != nil {
		return err
	}
	oswin.TheApp.RunOnMain(func() {
		ms.MakeVectorsImpl(sc)
	})
}

func (ms *MeshBase) MakeVectorsImpl(sc *Scene) {
	var vbuf gpu.VectorsBuffer
	var ibuf gpu.IndexesBuffer
	if ms.Buff == nil {
		ms.Buff = gpu.TheGPU.NewBufferMgr()
		usg := pu.StaticDraw
		if ms.Dynamic {
			usg = gpu.DynamicDraw
		}
		vbuf = ms.Buff.AddVectorsBuffer(usg)
		ibuf = ms.Buff.AddIndexesBuffer(gpu.StaticDraw)
	} else {
		vbuf = ms.Buff.VectorsBuffer()
		ibuf = ms.Buff.IndexesBuffer()
	}
	nvec := 3
	hasColor := false
	if len(ms.Color) > 0 {
		hasColor = true
		nvec++
	}
	vtx := sc.Rends.Vectors[InVtxPos]
	nrm := sc.Rends.Vectors[InVtxNorm]
	tex := sc.Rends.Vectors[InVtxTexUV]
	clr := sc.Rends.Vectors[InVtxColor]
	if vbuf.NumVectors() != nvec {
		vbuf.DeleteAllVectors()
		vbuf.AddVectors(vtx, true) // interleave
		vbuf.AddVectors(nrm, true) // interleave
		vbuf.AddVectors(tex, true) // interleave
		if hasColor {
			vbuf.AddVectors(clr, false) // NO interleave
		}
	}
	vln := len(ms.Vtx)
	vbuf.SetLen(vln)
	vbuf.SetVecData(vtx, ms.Vtx)
	vbuf.SetVecData(nrm, ms.Norm)
	vbuf.SetVecData(tex, ms.TexUV)
	if hasColor {
		vbuf.SetVecData(clr, ms.Color)
	}
	iln := len(ms.Idx)
	ibuf.SetLen(iln)
	ibuf.Set(ms.Idx)
}

// Activate activates the mesh Vectors on the GPU
// Must be called with relevant context active
func (ms *MeshBase) Activate(sc *Scene) {
	if ms.Buff == nil {
		ms.MakeVectors(sc)
	}
	oswin.TheApp.RunOnMain(func() {
		ms.Buff.Activate()
	})
}

// TransferAll transfer all buffer data to GPU (vectors and indexes)
// Activate must have just been called
func (ms *MeshBase) TransferAll() {
	oswin.TheApp.RunOnMain(func() {
		ms.Buff.TransferAll()
	})
}

// TransferVectors transfer vectors buffer data to GPU (if vector data has changed)
// Activate must have just been called
func (ms *MeshBase) TransferVectors() {
	oswin.TheApp.RunOnMain(func() {
		ms.Buff.TransferVectors()
	})
}

// TransferIndexes transfer vectors buffer data to GPU (if index data has changed)
// Activate must have just been called
func (ms *MeshBase) TransferIndexes() {
	oswin.TheApp.RunOnMain(func() {
		ms.Buff.TransferIndexes()
	})
}

/////////////////////////////////////////////////////////////////////
//  Shape primitives

// AddPlane adds everything to render a plane with the given parameters.
// waxis, haxis = width, height axis, wdir, hdir are the directions for width and height dimensions.
// wsegs, hsegs = number of segments to create in each dimension -- more finely subdividing a plane
// allows for higher-quality lighting and texture rendering (minimum of 1 will be enforced).
// offset is the distance to place the plane along the orthogonal axis.
func (ms *MeshBase) AddPlane(waxis, haxis mat32.Components, wdir, hdir int, width, height, offset float32, wsegs, hsegs int, clr gi.Color) {
	idxSt := ms.Vtx.Len() / 3 // starting index based on what's there already

	w := mat32.Z
	if (waxis == mat32.X && haxis == mat32.Y) || (waxis == mat32.Y && haxis == mat32.X) {
		w = mat32.Z
	} else if (waxis == mat32.X && haxis == mat32.Z) || (waxis == mat32.Z && haxis == mat32.X) {
		w = mat32.Y
	} else if (waxis == mat32.Z && haxis == mat32.Y) || (waxis == mat32.Y && haxis == mat32.Z) {
		w = mat32.X
	}
	wsegs = ints.MaxInt(wsegs, 1)
	hsegs = ints.MaxInt(hsegs, 1)

	norm := mat32.Vec3{}
	if offset > 0 {
		norm.SetComponent(w, 1)
	} else {
		norm.SetComponent(w, -1)
	}

	wHalf := width / 2
	hHalf := height / 2
	wsegs1 := wsegs + 1
	hsegs1 := hsegs + 1
	segWidth := width / float32(wsegs)
	segHeight := height / float32(hsegs)

	// Generate the plane vertices, norms, and uv coordinates
	for iy := 0; iy < hsegs1; iy++ {
		for ix := 0; ix < wsegs1; ix++ {
			vtx := mat32.Vec3{}
			vtx.SetComponent(waxis, (float32(ix)*segWidth-wHalf)*float32(wdir))
			vtx.SetComponent(haxis, (float32(iy)*segHeight-hHalf)*float32(hdir))
			vtx.SetComponent(w, offset)
			ms.Vtx.AppendVec3(&vtx)
			ms.Norm.AppendVec3(&norm)
			ms.TexUV.Append(float32(ix)/float32(wsegs), float32(1)-(float32(iy)/float32(hsegs)))
			if !clr.IsNil() {
				cv := ColorToVec4f(clr)
				ms.Color.AppendVec4(&cv)
			}
		}
	}

	// Generate the indices
	for iy := 0; iy < hsegs; iy++ {
		for ix := 0; ix < wsegs; ix++ {
			a := ix + wsegs1*iy
			b := ix + wsegs1*(iy+1)
			c := (ix + 1) + wsegs1*(iy+1)
			d := (ix + 1) + wsegs1*iy
			ms.Idx.Append(uint32(a+idxSt), uint32(b+idxSt), uint32(d+idxSt), uint32(b+idxSt), uint32(c+idxSt), uint32(d+idxSt))
		}
	}
}
