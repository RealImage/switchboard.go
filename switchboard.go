package switchboard

// Supply needs to be implemented by any type acting as a supplier.
// Must be safe for concurrent use by multiple goroutines.
type Supply interface {
	Estimate(demand Demand, choicesMade []Choice) Choice
}

// Demand is an empty interface that denotes a demand.
type Demand interface {
}

// Choice represents the fulfillment of a Demand with a particular Supply with
// its corresponding cost.
type Choice struct {
	Supply     Supply
	Demand     Demand
	Cost       float64
	Attributes map[string]interface{}
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
