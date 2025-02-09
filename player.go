package main

type Player struct {
	Username string `json:"username"`
	Deck Deck `json:"deck"`
}

type PlayerStore struct {
	Username string `json:"username"`
	Score int `json:"score"`
}