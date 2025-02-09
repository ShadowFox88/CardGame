package main

import (
	"testing"
)

func TestBestCard(t *testing.T) {
	tests := []struct {
		name     string
		card1    Card
		card2    Card
		expected Card
	}{
		{
			name:     "Same color, card1 higher",
			card1:    Card{Colour: "Red", Value: 5},
			card2:    Card{Colour: "Red", Value: 3},
			expected: Card{Colour: "Red", Value: 5},
		},
		{
			name:     "Same color, card2 higher",
			card1:    Card{Colour: "Yellow", Value: 2}, 
			card2:    Card{Colour: "Yellow", Value: 4},
			expected: Card{Colour: "Yellow", Value: 4},
		},
		{
			name:     "Red beats Black",
			card1:    Card{Colour: "Red", Value: 1},
			card2:    Card{Colour: "Black", Value: 9},
			expected: Card{Colour: "Red", Value: 1},
		},
		{
			name:     "Yellow beats Red",
			card1:    Card{Colour: "Yellow", Value: 2},
			card2:    Card{Colour: "Red", Value: 8}, 
			expected: Card{Colour: "Yellow", Value: 2},
		},
		{
			name:     "Black beats Yellow",
			card1:    Card{Colour: "Black", Value: 3},
			card2:    Card{Colour: "Yellow", Value: 7},
			expected: Card{Colour: "Black", Value: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BestCard(tt.card1, tt.card2)
			if got != tt.expected {
				t.Errorf("BestCard() = %v, want %v", got, tt.expected)
			}
		})
	}
}