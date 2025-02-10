package main

type Player struct {
	Username string `json:"username"`
	Deck Deck `json:"deck"`
}

type PlayerStore struct {
	Username string `json:"username"`
	Score int `json:"score"`
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
