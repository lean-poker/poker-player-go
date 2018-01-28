package main

const VERSION = "Default Go folding player"

type PokerPlayer struct{}

func NewPokerPlayer() *PokerPlayer {
	return &PokerPlayer{}
}

func (p *PokerPlayer) BetRequest(state *Game) int {
	return 0
}

func (p *PokerPlayer) Showdown(state *Game) {

}

func (p *PokerPlayer) Version() string {
	return VERSION
}
