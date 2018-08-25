package switchboard_test

import (
	"testing"

	"github.com/RealImage/switchboard.go"
	"github.com/stretchr/testify/assert"
)

func TestBoardCreation(t *testing.T) {
	supplies := []switchboard.Supply{testSupply{factor: 1}, testSupply{factor: 1.1}}
	demands := []switchboard.Demand{testDemand{cost: 1}, testDemand{cost: 2}}
	board := switchboard.NewBoard(supplies, demands)
	assert.Equal(t, supplies, board.Supplies())
	assert.Equal(t, demands, board.Demands())
	assert.InDelta(t, 0, board.Cost(), 0.01)
	assert.Zero(t, len(board.ChoicesMade()))

	bestBoard := board.Explore(func(board switchboard.Board) bool {
		return true
	})
	assert.Equal(t, 2, len(bestBoard.ChoicesMade()))

}
