package main

import (
	"fmt"
	"sort"
)

var player_1 Player
var player_2 Player

var main_deck Deck = Deck{}



func main() {
	var p1,p2,u1,u2 string

	fmt.Println("Player 1, Please Enter The Password:")
	fmt.Scanln(&p1)

	if !authenticate(p1) {
		panic("Incorrect Password")
	}

	fmt.Println("Player 1, Please Enter Your Username:")
	fmt.Scanln(&u1)

	if len(u1) == 0 {
		panic("Username cannot be empty")
	}

	fmt.Println("Player 2, Please Enter The Password:")
	fmt.Scanln(&p2)

	if !authenticate(p2) {
		panic("Incorrect Password")
	}

	fmt.Println("Player 2, Please Enter Your Username:")
	fmt.Scanln(&u2)

	if len(u2) == 0 {
		panic("Username cannot be empty")
	}

	player_1 = Player{u1, Deck{}}
	player_2 = Player{u2, Deck{}}

	fmt.Println()

	main_deck.Init()

	for {
		if len(main_deck.Cards) == 0 {
			break // I know you once told me off for this, saying this was bad practice and lazy, but golang doesn't have while loops
		}

		fmt.Println("Round Start")

		fmt.Printf("%s's Turn, Press Enter to Draw", player_1.Username)
		fmt.Scanln()
		
		player_1_card := main_deck.Draw()
		fmt.Printf("%s drew %s \n\n", player_1.Username, player_1_card.toString())

		fmt.Printf("%s's Turn, Press Enter to Draw\n", player_2.Username)
		fmt.Scanln()

		player_2_card := main_deck.Draw()
		fmt.Printf("%s drew %s \n\n", player_2.Username, player_2_card.toString())

		winning_card := BestCard(player_1_card, player_2_card)

		if winning_card == player_1_card {
			player_1.Deck.Cards = append(player_1.Deck.Cards, player_1_card, player_2_card)
			fmt.Println(player_1.Username, "won the round")
		} else {
			player_2.Deck.Cards = append(player_2.Deck.Cards, player_1_card, player_2_card)
			fmt.Println(player_2.Username, "won the round")
		}
		fmt.Println()
	}

	if len(player_1.Deck.Cards) == len(player_2.Deck.Cards) {
		fmt.Println("It's a draw!")
	} else if len(player_1.Deck.Cards) > len(player_2.Deck.Cards) {
		fmt.Println(player_1.Username, "won the game!")
		fmt.Printf("Their Cards:")
		for _, card := range player_1.Deck.Cards {
			fmt.Printf(" %s,", card.toString())
		}
	} else {
		fmt.Println(player_2.Username, "won the game!")
		fmt.Printf("Their Cards:")
		for _, card := range player_1.Deck.Cards {
			fmt.Printf(" %s", card.toString())
		}
	}
	fmt.Println()

	previousWinners := GetPlayers()

	if len(previousWinners) == 0 {
		fmt.Println("No previous winners")
	} else {
		fmt.Println("Previous Winners:")
	}
	for _, player := range previousWinners {
		fmt.Println(player.Username, "with", player.Score, "cards")
	}

	previousWinners = append(previousWinners, PlayerStore{player_1.Username, len(player_1.Deck.Cards)})
	previousWinners = append(previousWinners, PlayerStore{player_2.Username, len(player_2.Deck.Cards)})

	sort.Slice(previousWinners, func(i, j int) bool {
		return previousWinners[i].Score > previousWinners[j].Score
	})

	if len(previousWinners) > 5 {
		previousWinners = previousWinners[:5]
	}

	WritePlayers(previousWinners)
}
