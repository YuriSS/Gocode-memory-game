package card

import (
	"math/rand"
	"time"
)

var card_images = [4]string{
	"/static/images/card_1.jpg",
	"/static/images/card_2.jpg",
	"/static/images/card_3.jpg",
	"/static/images/card_4.jpg",
}

type Deck struct {
	cards_order [8]struct {
		card  string
		index int
	}
	lefts int
}

var source = rand.NewSource(time.Now().UnixNano())

func (d *Deck) New_Turn() {
	suffle(d)
	d.lefts = cap(d.cards_order)
}

func suffle(d *Deck) {
	for i, image := range card_images {
		var times int = 0
		for {
			rand_val := random_number(cap(d.cards_order))
			if d.cards_order[rand_val].card == "" {
				d.cards_order[rand_val].card = image
				d.cards_order[rand_val].index = i
				if times++; times >= 2 {
					break
				}
			}
		}
	}
}

func (d *Deck) Get_Card(index int) string {
	return d.cards_order[index].card
}

func (d *Deck) Has_Any_Left() bool {
	if d.lefts == 0 {
		return false
	}
	return true
}

func (d *Deck) Compare_Cards(card_1, card_2 int) bool {
	if d.cards_order[card_1].index == d.cards_order[card_2].index {
		return true
	}
	return false
}

func random_number(max int) int {
	random := rand.New(source)
	result := random.Intn(max)
	return result
}
