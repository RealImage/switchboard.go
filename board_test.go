package switchboard_test

import (
	"testing"

	"github.com/RealImage/switchboard.go"
	"github.com/stretchr/testify/assert"
)

func TestExploration(t *testing.T) {
	supplies := []switchboard.Supply{testSupply{factor: 1}, testSupply{factor: 1.1}}
	demands := []switchboard.Demand{testDemand{cost: 1}, testDemand{cost: 2}}

	board := switchboard.NewBoard(supplies, demands, switchboard.PrioritizeWorkDone(switchboard.Minimize()))
	assert.Equal(t, supplies, board.Supplies())
	assert.Equal(t, demands, board.Demands())
	assert.InDelta(t, 0, board.Cost(), 0.01)
	assert.Zero(t, len(board.ChoicesMade()))

	bestBoard := board.Explore(switchboard.BruteForceExplorer())
	assert.Equal(t, 2, len(bestBoard.ChoicesMade()))
	assert.InDelta(t, 3, bestBoard.Cost(), 0.01)

	board2 := switchboard.NewBoard(supplies, demands, switchboard.PrioritizeWorkDone(switchboard.Maximize()))
	bestBoard2 := board2.Explore(switchboard.BruteForceExplorer())
	assert.Equal(t, 2, len(bestBoard2.ChoicesMade()))
	assert.InDelta(t, 3.3, bestBoard2.Cost(), 0.01)
}
