package main

import (
	"reflect"
	"testing"
)

func TestDeckCreate(t *testing.T) {
	deck := Deck{}
	deck.createCards()

	// Check we actually have 30 cards
	if len(deck.Cards) != 30 {
		t.Errorf("Expected deck length of 30, but got %d", len(deck.Cards))
	}

		// Check if all colors and numbers are present
		colorCount := make(map[string]int)
		numberCount := make(map[int]int)
	
		for _, card := range deck.Cards {
			colorCount[card.Colour]++
			numberCount[card.Value]++
		}
	
		if len(colorCount) != 3 {
			t.Error("Not all colors are present in deck")
		}
		if len(numberCount) != 10 {
			t.Error("Not all numbers are present in deck")
	}	
}

func TestDeckInit(t *testing.T) {
	deck := Deck{}
	deck.Init()

	// Check if deck has correct number of cards (3 colors * 10 numbers = 30 cards)
	if len(deck.Cards) != 30 {
		t.Errorf("Expected deck length of 30, but got %d", len(deck.Cards))
	}

	// Check if all colors and numbers are present
	colorCount := make(map[string]int)
	numberCount := make(map[int]int)

	for _, card := range deck.Cards {
		colorCount[card.Colour]++
		numberCount[card.Value]++
	}

	if len(colorCount) != 3 {
		t.Error("Not all colors are present in deck")
	}
	if len(numberCount) != 10 {
		t.Error("Not all numbers are present in deck")
	}

	test_deck := Deck{}
	test_deck.createCards()
	cards := test_deck.Cards

	// Check that they aren't all the same as the default list
	if reflect.DeepEqual(deck.Cards, cards) {
		t.Errorf("The deck is not shuffled. \nTest Deck: %v\nShuffled Deck: %v", cards, deck.Cards)
	}
}

func TestDeckDraw(t *testing.T) {
	deck := Deck{}
	deck.Init()
	
	initialLength := len(deck.Cards)
	drawnCard := deck.Draw()

	// Check if deck size decreased
	if len(deck.Cards) != initialLength-1 {
		t.Error("Deck size did not decrease after drawing")
	}

	// Check if drawn card is valid
	if drawnCard.Value < 1 || drawnCard.Value > 10 {
		t.Error("Drawn card has invalid number")
	}
	
	validColors := map[string]bool{"Red": true, "Yellow": true, "Black": true}
	if !validColors[drawnCard.Colour] {
		t.Error("Drawn card has invalid color")
	}
}