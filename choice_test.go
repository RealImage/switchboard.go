package switchboard_test

import (
	"testing"

	"github.com/RealImage/switchboard.go"
	"github.com/stretchr/testify/assert"
)

type testSupply struct {
}

func (ts testSupply) Estimate(demand switchboard.Demand, choicesMade []switchboard.Choice) (switchboard.Choice, error) {
	return switchboard.NewChoice(nil, nil, 0.123, nil), nil
}

type testDemand struct {
}

func TestChoice(t *testing.T) {
	supply := testSupply{}
	demand := testDemand{}
	cost := 3.14
	choice := switchboard.NewChoice(supply, demand, cost, map[string]interface{}{"abc": 123})
	assert.Equal(t, supply, choice.Supply())
	assert.Equal(t, demand, choice.Demand())
	assert.Equal(t, cost, choice.Cost())
	assert.Equal(t, 123, choice.Get("abc"))
	assert.Nil(t, choice.Get("blah"))

}
