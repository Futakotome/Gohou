package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"gohou/window"
	"runtime"
)

var (
	triangle = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
)

func init() {
}

func main() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
	gameWindow := window.InitGlfw()
	defer glfw.Terminate()
	prog := window.InitOpenGL()

	vao := window.MakeVao(triangle)
	for !gameWindow.ShouldClose() {
		drawVao(vao, gameWindow, prog)
		//if err := drawVao(vao, gameWindow,prog); err != nil {
		//	panic(err)
		//}
	}
}

func drawVao(vao uint32, window *glfw.Window, prog uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

func draw(prog uint32, window *glfw.Window) error {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)
	glfw.PollEvents()
	window.SwapBuffers()
	return nil
}
