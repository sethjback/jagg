package guess

import "math/rand"

type Number struct {
	Rounds int
	Max    int
	Secret int
}

func NewNumber(rounds, max int) *Number {
	return &Number{
		Rounds: rounds,
		Max:    max,
		Secret: rand.Intn(max + 1),
	}
}

func (n *Number) Guess(g int) bool {
	n.Rounds -= 1
	return n.Secret == g
}

func (n *Number) RoundsLeft() int {
	return n.Rounds
}
