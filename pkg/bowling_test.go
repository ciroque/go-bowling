package bowling

import (
	"fmt"
	"testing"
)

func Test_TwoGutterBalls(t *testing.T) {
	bowling := NewBowling()

	bowling.Throw(0)
	bowling.Throw(0)

	score := bowling.GetScore()
	if score != 0 {
		t.Fatal(fmt.Sprintf("Expected score to be 0, was: %d", score))
	}
}