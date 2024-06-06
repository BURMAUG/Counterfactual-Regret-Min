package main

import "fmt"

// Action profile.
const (
	ROCK        = 0
	PAPER       = 1
	SCISSORS    = 2
	NUM_ACTIONS = 3 // This is the number of actions a player can play in a given game
)

type Actions interface {
	Strategy() int
	Action(strategy []byte) int // This should return a players Strategy
}

// Player is a structure that takes the player form.
type Player struct {
}

// Return a new Player object pointer
func NewPlay() *Player {
	return &Player{}
}

// Strategy
func (p *Player) Strategy() int {
	return ROCK
}

// Action
func (p *Player) Action(strategy []byte) {

}

// PlayerPayOffs is a function that should map player action to player payoff
func (p *Player) PlayerPayOffs(p1, p2 *Player) {
}

func main() {
	fmt.Print("Burmau Garba")
}
