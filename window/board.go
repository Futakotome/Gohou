package window

import "gohou/character"

type playBoard struct {
	character character.Character
	enemies   []character.Enemy
}

func (playBoard playBoard) refresh() {

}

type bonusBoard struct {
	currentBonus int //当前分数
	maxBonus     int //最高分数
	life         int //生命
	skill        int //大招
}

func (bonusBoard bonusBoard) refresh() {

}
