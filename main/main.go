package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"gohou/window"
	"runtime"
)

func init() {
}

func main() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
	gameWindow := window.InitGlfw()
	defer glfw.Terminate()
	prog := window.InitOpenGL()

	for !gameWindow.ShouldClose() {
		if err := draw(prog, gameWindow); err != nil {
			panic(err)
		}
	}
}

func draw(prog uint32, window *glfw.Window) error {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)
	glfw.PollEvents()
	window.SwapBuffers()
	return nil
}
