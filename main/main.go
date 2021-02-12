package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"gohou/window"
	"log"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	log.Println("初始化环境----->")
	log.Printf("glfw版本:%+v.%+v.%+v", glfw.VersionMajor, glfw.VersionMinor, glfw.VersionRevision)
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, glfw.VersionMajor)
	glfw.WindowHint(glfw.ContextVersionMinor, glfw.VersionMinor)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfwWindow, err := glfw.CreateWindow(window.Width1600, window.Height900, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}
	glfwWindow.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Printf("OpenGL版本:%+v", version)
	for !glfwWindow.ShouldClose() {
		glfwWindow.SwapBuffers()
		glfw.PollEvents()
	}

}
