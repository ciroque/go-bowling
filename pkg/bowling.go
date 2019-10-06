package bowling

type Bowling struct {
	score int
}

func NewBowling() Bowling {
	return Bowling{}
}

func (b *Bowling) Throw(pins int) {
	b.score = b.score + pins
}

func (b *Bowling) GetScore() int {
	return b.score
}
