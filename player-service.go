package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const VERSION = "Default Go folding player"

func BetRequest(state *GameState) int {
	return 0
}

func Showdown(state *GameState) {

}

func Version() string {
	return VERSION
}

func main() {
	http.HandleFunc("/", handleRequest)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing PORT environment variable")
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Printf("Error parsing form data: %s", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	action := request.FormValue("action")
	log.Printf("Request method=%s url=%s action=%s from client=%s\n", request.Method, request.URL, action, request.RemoteAddr)
	switch action {
	case "check":
		fmt.Fprint(w, "")
		return
	case "bet_request":
		gameState := parseGameState(request.FormValue("game_state"))
		if gameState == nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		result := BetRequest(gameState)
		fmt.Fprintf(w, "%d", result)
		return
	case "showdown":
		gameState := parseGameState(request.FormValue("game_state"))
		if gameState == nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		Showdown(gameState)
		fmt.Fprint(w, "")
		return
	case "version":
		fmt.Fprint(w, Version())
		return
	default:
		http.Error(w, "Invalid action", 400)
	}
}

type GameState struct {

	// The small blind in the current round. The big blind is twice
	// the small blind
	SmallBlind int `json:"small_blind"`

	// The amount of the largest current bet from any one player
	CurrentBuyIn int `json:current_buy_in`

	// The size of the pot (sum of the player bets)
	Pot int `json:pot`

	// Minimum raise amount. To raise you have to return at least:
	//     current_buy_in - players[in_action][bet] + minimum_raise
	MinimumRaise int `json:minimum_raise`

	// The index of the player on the dealer button in this round
	// The first player is (dealer+1)%(players.length)
	Dealer int `json:dealer`

	// Number of orbits completed. (The number of times the dealer button
	// returned to the same player.)
	Orbits int `json:orbits`

	// The index of your player, in the players array
	InAction int `json:in_action`

	// An array of the players. The order stays the same during the
	// entire tournament
	Players []Player `json:players`

	// Finally the array of community cards.
	CommunityCards []Card `json:community_cards`
}

type Player struct {

	// Id of the player (same as the index)
	Id int `json:"id"`

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

type Card struct {

	// Rank of the card. Possible values are numbers 2-10 and J,Q,K,A
	Rank string `json:"rank"`

	// Suit of the card. Possible values are: clubs,spades,hearts,diamonds
	Suit string `json:"suit"`
}

func parseGameState(stateStr string) *GameState {
	stateBytes := []byte(stateStr)
	gameState := new(GameState)
	if err := json.Unmarshal(stateBytes, &gameState); err != nil {
		log.Printf("Error parsing game state: %s", err)
		return nil
	}
	return gameState
}
