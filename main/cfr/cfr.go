// Over multiple iterations I need to add up my regrets
package cfr

// Action profile.
const (
	ROCK        = 0
	PAPER       = 1
	SCISSORS    = 2
	NUM_ACTIONS = 3 // This is the number of actions a player can play in a given game
)

// Action
type Actions interface {
	Strategy() []int
	Action(strategy []int) int // This should return a players Strategy
}

// Player is a structure that takes the player form.
type Player struct {
	PlayerName  string
	strategy    [NUM_ACTIONS]int
	regretSum   [NUM_ACTIONS]int
	strategySum [NUM_ACTIONS]int
}

// Return a new Player object pointer
func NewPlay(name string) *Player {
	return &Player{PlayerName: name}
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
// in order for me to know the payoff for p1 in need to know what move p1 made and what move p2 made.
func (p *Player) PlayerPayOffs(p1, p2 *Player) {
	// p1.Play(ROCK) and p2.Play(PAPER) -1 1
	// p1.Play(ROCK) and p2.Play(SCISSORS) 1 -1
	// p1.Play(ROCK) and p2.Play(ROCK) 0 0

	// p1.Play(PAPER) and p2.Play(PAPER)
	// p1.Play(PAPER) and p2.Play(PAPER)
	// p1.Play(PAPER) and p2.Play(PAPER)

	// p1.Play(ROCK) and p2.Play(PAPER)
	// p1.Play(ROCK) and p2.Play(PAPER)
	// p1.Play(ROCK) and p2.Play(PAPER)
}
