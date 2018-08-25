package switchboard_test

import (
	"testing"

	"github.com/RealImage/switchboard.go"
	"github.com/stretchr/testify/assert"
)

func TestBoardCreation(t *testing.T) {
	supplies := []switchboard.Supply{testSupply{}, testSupply{}}
	demands := []switchboard.Demand{testDemand{}}
	board := switchboard.NewBoard(supplies, demands)
	assert.Equal(t, supplies, board.Supplies())
	assert.Equal(t, demands, board.Demands())
	assert.InDelta(t, 0, board.Cost(), 0.01)
	assert.Zero(t, len(board.ChoicesMade()))
}
