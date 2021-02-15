package window

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type container interface {
	refresh()
}

type GameWindow struct {
	window     *glfw.Window
	playBoard  playBox
	bonusBoard bonusBoard
}

func (gameWindow GameWindow) refresh() {

}

func refreshWindow() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
