package window

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"strings"
)

const (
	FPS        = 10 //帧率  time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	WIDTH800   = 800
	HEIGHT600  = 600
	WIDTH1024  = 1024
	HEIGHT768  = 768
	Width1600  = 1600
	Height900  = 900
	WIDTH1920  = 1920
	HEIGHT1080 = 1080

	vertexShaderSource = `
		#version 400
		
		in vec3 vp;
		void main(){
			gl_Position = vec4(vp,1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 400

		out vec4 frag_colour;
		void main(){
			frag_colour = vec4(1,1,1,1.0);
		}
	` + "\x00"
)

func InitGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(Width1600, Height900, "MyTouhouGame", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	return window
}

func InitOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGl version ", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		repeat := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(repeat))

		return 0, fmt.Errorf("failed to compile %v:%v", source, repeat)
	}
	return shader, nil
}
