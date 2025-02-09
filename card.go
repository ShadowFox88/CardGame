package main

import (
	"strconv"
)

type Card struct {
	Colour string
	Value int
}

func (c *Card) toString() string {
	return c.Colour + " " + strconv.Itoa(c.Value)
}

func BestCard(card1 Card, card2 Card) Card {
	if card1.Colour == card2.Colour {
		if card1.Value > card2.Value {
			return card1
		} else {
			return card2
		}
	}

	if (card1.Colour == "Red" && card2.Colour == "Black") ||  // Red Wins
	   (card1.Colour == "Yellow" && card2.Colour == "Red") ||  // Yellow Wins
	   (card1.Colour == "Black" && card2.Colour == "Yellow") { // Black Wins
		return card1
	}

	return card2
}

