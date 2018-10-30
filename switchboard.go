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
func NewChoice(supply Supply, demand Demand, cost float64, attributes map[string]interface{}) (choice Choice) {
	choice.supply = supply
	choice.demand = demand
	choice.cost = cost
	choice.attributes = make(map[string]interface{})
	for k, v := range attributes {
		choice.attributes[k] = v
	}
	return
}

// Player evaluates a board and chooses a new board using a built-in strategy.
type Player func(board Board) (chosenBoard Board)

// GreedyPlayer is a fast player that just chooses the best available choice,
// resulting in a local minima.
func GreedyPlayer() Player {
	return func(board Board) Board {
		currentBestBoard := board
		for _, choice := range board.availableChoices() {
			candidateBoard := board.choose(choice)
			if board.comparator(currentBestBoard, candidateBoard) {
				currentBestBoard = candidateBoard
			}
		}
		return currentBestBoard
	}
}

// Explorer evaluates a board to determine whether its choices are worth
// exploring further. Must be safe for concurrent use by multiple goroutines.
type Explorer func(board Board) bool

// BruteForceExplorer forces the exploration of every possible board
func BruteForceExplorer() Explorer {
	return func(board Board) bool { return true }
}

// GoalTransformer allows transformation of choice scores to determine the goal
// of the exploration
type GoalTransformer func(cost float64) float64

// Maximize provides a GoalTransformer which helps find the board with the
// highest total cost
func Maximize() GoalTransformer {
	return func(cost float64) float64 { return cost }
}

// Minimize provides a GoalTransformer which helps find the board with the
// lowest total cost
func Minimize() GoalTransformer {
	return func(cost float64) float64 { return cost * -1 }
}

// BoardComparator returns true if board2 is better than board1.
type BoardComparator func(board1, board2 Board) bool

// PrioritizeWorkDone gives first priority to the number of choices made - it
// can be used when it's more important to fulfill as much demand as possible
// than to just meet the cost goals
func PrioritizeWorkDone(transformer GoalTransformer) BoardComparator {
	return func(b1, b2 Board) bool {
		if len(b1.ChoicesMade()) == len(b2.ChoicesMade()) {
			if transformer(b2.Cost()) > transformer(b1.Cost()) {
				return true
			}
		} else {
			if len(b2.ChoicesMade()) > len(b1.ChoicesMade()) {
				return true
			}
		}
		return false
	}
}

// PrioritizeTotalCost simply targets a board with the given cost goal
func PrioritizeTotalCost(transformer GoalTransformer) BoardComparator {
	return func(b1, b2 Board) bool {
		return transformer(b2.Cost()) > transformer(b1.Cost())
	}
}
