package switchboard

// Supply needs to be implemented by any type acting as a supplier.
// Must be safe for concurrent use by multiple goroutines.
type Supply interface {
	Estimate(demand Demand, choicesMade []Choice) (Choice, error)
}

// Demand is an empty interface that denotes a demand.
type Demand interface {
}

// Choice represents the fulfillment of a Demand with a particular Supply with
// its corresponding cost.
type Choice struct {
	supply     Supply
	demand     Demand
	cost       float64
	attributes map[string]interface{}
}

// Supply returns the supply associated with the current choice
func (choice Choice) Supply() Supply {
	return choice.supply
}

// Demand returns the demand associated with the current choice
func (choice Choice) Demand() Demand {
	return choice.demand
}

// Cost returns the cost associated with the current choice
func (choice Choice) Cost() float64 {
	return choice.cost
}

// Get retrieves the associated value from the attribute map
func (choice Choice) Get(key string) interface{} {
	return choice.attributes[key]
}

// NewChoice build a new choice from the given supply, demand, cost and attributes.
func NewChoice(supply Supply, demand Demand, cost float64, attributes map[string]interface{}) Choice {
	return Choice{
		supply:     supply,
		demand:     demand,
		cost:       cost,
		attributes: attributes,
	}
}

// Explorer evaluates a board to determine whether its choices are worth exploring further.
// Must be safe for concurrent use by multiple goroutines.
type Explorer interface {
	ShouldExplore(board Board) bool
}
