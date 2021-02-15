package window

import "gohou/character"

type playBox struct {
	character character.Character
	enemies   []character.Enemy
}

func (playBoard playBox) refresh() {
}

type bonusBoard struct {
	currentBonus int //当前分数
	maxBonus     int //最高分数
	life         int //生命
	spell        int //符卡
}

func (bonusBoard bonusBoard) refresh() {

}
