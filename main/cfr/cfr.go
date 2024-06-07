// Over multiple iterations I need to add up my regrets
package cfr

// Action profile.
const (
	ROCK        = 0
	PAPER       = 1
	SCISSORS    = 2
	NUM_ACTIONS = 3 // This is the number of actions a player can play in a given game
)

type Actions interface {
	Strategy() []int
	Action(strategy []byte) int // This should return a players Strategy
}

// Player is a structure that takes the player form.
type Player struct {
	strategy    [NUM_ACTIONS]int
	regretSum   [NUM_ACTIONS]int
	strategySum [NUM_ACTIONS]int
}

// Return a new Player object pointer
func NewPlay() *Player {
	return &Player{}
}

// Strategy
func (p *Player) Strategy() []int {
	normalizedStrategy := 0
	for a := 0; a < NUM_ACTIONS; a++ {
		if p.strategy[a] > 0 {
			p.strategy[a] = p.regretSum[a]
		} else {
			p.strategy[a] = 0
		}
		normalizedStrategy += p.strategy[a]
	}

	for a := 0; a < NUM_ACTIONS; a++ {
	}
	return p.strategy[:]
}

// Action
func (p *Player) Action(strategy []int) {
}

// PlayerPayOffs is a function that should map player action to player payoff
func (p *Player) PlayerPayOffs(p1, p2 *Player) {
}
