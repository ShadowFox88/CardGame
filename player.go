package main

type Player struct {
	Username string
	Deck     Deck
}

type PlayerStore struct {
	Username string
	Score    int
}

func ValidUsername(username string) bool {
	length := len(username)

	if length == 0 {
		return false
	} else if length >= 15 {
		return false
	}
	return true
}
