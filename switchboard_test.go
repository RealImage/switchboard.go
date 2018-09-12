package switchboard_test

import "github.com/RealImage/switchboard.go"

type testSupply struct {
	factor float64
}

func (ts testSupply) Estimate(demand switchboard.Demand, choicesMade []switchboard.Choice) (switchboard.Choice, error) {
	estimate := ts.factor * demand.(testDemand).cost
	return switchboard.NewChoice(ts, demand, estimate, nil), nil
}

type testDemand struct {
	cost float64
}
