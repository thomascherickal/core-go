// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is initially adapted from https://github.com/vulkan-go/demos
// Copyright © 2017 Maxim Kupriianov <max@kc.vc>, under the MIT License
// and https://bakedbits.dev/posts/vulkan-compute-example/

package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	vk "github.com/vulkan-go/vulkan"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/vgpu/vgpu"
	"github.com/xlab/closer"
)

func init() {
	// must lock main thread for gpu!  this also means that vulkan must be used
	// for gogi/oswin eventually if we want gui and compute
	runtime.LockOSThread()
}

var TheGPU *vgpu.GPU

func main() {
	glfw.Init()
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())
	vk.Init()
	defer closer.Close()

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	window, err := glfw.CreateWindow(1024, 768, "Draw Triangle", nil, nil)
	vgpu.IfPanic(err)

	// note: for graphics, require these instance extensions before init gpu!
	winext := window.GetRequiredInstanceExtensions()
	gp := vgpu.NewGPU()
	gp.AddInstanceExt(winext...)
	gp.Debug = true
	gp.Config("drawtri")
	TheGPU = gp

	// gp.PropsString(true) // print

	surfPtr, err := window.CreateWindowSurface(gp.Instance, nil)
	if err != nil {
		log.Println(err)
		return
	}
	sf := vgpu.NewSurface(gp, vk.SurfaceFromPointer(surfPtr))

	fmt.Printf("format: %#v\n", sf.Format)

	sy := gp.NewGraphicsSystem("drawtri", &sf.Device)
	pl := sy.NewPipeline("drawtri")
	sy.SetRenderPass(&sf.Format, vk.FormatUndefined)
	sf.SetRenderPass(&sy.RenderPass)
	pl.SetGraphicsDefaults()

	pl.AddShaderFile("trianglelit", vgpu.VertexShader, "trianglelit.spv")
	pl.AddShaderFile("vtxcolor", vgpu.FragmentShader, "vtxcolor.spv")

	inv := sy.Vars.Add("Vtx", vgpu.Float32Vec4, vgpu.Uniform, 0, vgpu.VertexShader)
	_ = inv

	sy.Config()
	sy.Mem.Config()

	destroy := func() {
		sy.Destroy()
		sf.Destroy()
		gp.Destroy()
		window.Destroy()
		glfw.Terminate()
	}

	frameCount := 0

	renderFrame := func() {
		fmt.Printf("frame: %d\n", frameCount)
		idx := sf.AcquireNextImage()
		cmd := pl.GraphicsCommand(sf.Frames[idx])
		sf.SubmitRender(cmd)
		sf.PresentImage(idx)
		frameCount++
	}

	exitC := make(chan struct{}, 2)

	fpsDelay := time.Second // / 60
	fpsTicker := time.NewTicker(fpsDelay)
	for {
		select {
		case <-exitC:
			fpsTicker.Stop()
			destroy()
			return
		case <-fpsTicker.C:
			if window.ShouldClose() {
				exitC <- struct{}{}
				continue
			}
			glfw.PollEvents()
			renderFrame()
		}
	}
}
