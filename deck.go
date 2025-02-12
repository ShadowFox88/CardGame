package main

import (
	"math/rand"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) Init() {
	d.createCards()
	d.shuffle()
}

func (d *Deck) createCards() {
	allColours := []string{"Red", "Yellow", "Black"}
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, colour := range allColours {
		for _, number := range allNumbers {
			d.Cards = append(d.Cards, Card{colour, number})
		}
	}
}

func (d *Deck) shuffle() {
    rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Draw() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}