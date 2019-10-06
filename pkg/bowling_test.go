package bowling

import (
	"fmt"
	"testing"
)

func Test_TwoGutterBalls(t *testing.T) {
	expectedScore := 0
	bowling := NewBowling()

	bowling.Throw(0)
	bowling.Throw(0)

	score := bowling.GetScore()
	if score != expectedScore {
		t.Fatal(fmt.Sprintf("Expected score to be %d, was: %d", expectedScore, score))
	}
}

func Test_OneGutterBallAndOneGoodThrow(t *testing.T) {
	expectedScore := 3
	bowling := NewBowling()

	bowling.Throw(3)
	bowling.Throw(0)

	score := bowling.GetScore()
	if score != 3 {
		t.Fatal(fmt.Sprintf("Expected score to be %d, was: %d", expectedScore, score))
	}
}

func Test_TwoGoodThrows(t *testing.T) {
	expectedScore := 8
	bowling := NewBowling()

	bowling.Throw(3)
	bowling.Throw(5)

	score := bowling.GetScore()
	if score != expectedScore {
		t.Fatal(fmt.Sprintf("Expected score to be %d, was: %d", expectedScore, score))
	}
}
