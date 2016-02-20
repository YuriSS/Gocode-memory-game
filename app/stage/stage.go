package stage

import (
	"github.com/YuriSS/memory_game/app/cards"
)

type Scene struct {
	deck card.Deck
}

func (scene *Scene) New_Scene() {
	scene.deck.New_Turn()
}
