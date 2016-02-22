package stage

import (
	"github.com/YuriSS/memory_game/app/cards"
)

type Scene struct {
	Deck     card.Deck
	New_Play struct {
		Index int `json:"index"`
	}
}

func (scene *Scene) New_Scene() {
	scene.Deck.New_Turn()
}

func (scene *Scene) Playing() string {
	var image string = scene.Deck.Get_Card(scene.New_Play.Index)
	return image
}
