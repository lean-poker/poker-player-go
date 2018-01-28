package main

// Player describes a player's properties in the game state JSON
type Player struct {
	// ID of the player (same as the index)
	ID int `json:"id"`

	// Name specified in the tournament config
	Name string `json:"name"`

	// Status of the player:
	//   - active: the player can make bets, and win the current pot
	//   - folded: the player folded, and gave up interest in
	//       the current pot. They can return in the next round.
	//   - out: the player lost all chips, and is out of this sit'n'go
	Status string `json:"status"`

	// Version identifier returned by the player
	Version string `json:"version"`

	// Amount of chips still available for the player.
	// (Not including the chips the player bet in this round)
	Stack int `json:"stack"`

	// The amount of chips the player put into the pot
	Bet int `json:"bet"`

	// The cards of the player. This is only visible for your own player
	// except after showdown, when cards revealed are also included.
	HoleCards []Card `json:"hole_cards"`
}
