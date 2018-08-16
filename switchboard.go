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
	return nil
}

// Demand returns the demand associated with the current choice
func (choice Choice) Demand() Demand {
	return nil
}

// Cost returns the cost associated with the current choice
func (choice Choice) Cost() float64 {
	return 0
}

// Get retrieves the associated value from the attribute map
func (choice Choice) Get(key string) interface{} {
	return nil
}

// NewChoice build a new choice from the given supply, demand, cost and attributes.
func NewChoice(supply Supply, demand Demand, cost float64, attributes map[string]interface{}) Choice {
	return Choice{}
}

// Explorer evaluates a board to determine whether its choices are worth exploring further.
// Must be safe for concurrent use by multiple goroutines.
type Explorer interface {
	ShouldExplore(board Board) bool
}

// Board represents a set of supplies and demands for which a universes of choices can be explored.
type Board struct {
	supplies []Supply
	demands  []Demand
}

// ChoicesMade returns a list of the choices made so far, in order
func (board Board) ChoicesMade() []Choice {
	return []Choice{}
}

// Cost returns the sum of the costs of all the choices made so far
func (board Board) Cost() float64 {
	return 0
}

// FindBestBoard uses the given explorer to discover the best sequence
// of choices among the universe of all possible choice sequences.
func (board Board) FindBestBoard(explorer Explorer) (bestBoard Board) {
	return Board{}
}
