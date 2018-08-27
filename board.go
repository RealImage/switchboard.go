package switchboard

// Board represents a set of supplies and demands for which a universes of choices can be explored.
type Board struct {
	supplies []Supply
	demands  []Demand
	choices  []Choice
}

// NewBoard constructs a new board with the given supplies and demands
func NewBoard(supplies []Supply, demands []Demand) (board Board) {
	board.supplies = append(board.supplies, supplies...)
	board.demands = append(board.demands, demands...)
	return
}

// ChoicesMade returns a list of the choices made so far, in the order they were made
func (board Board) ChoicesMade() (choicesMade []Choice) {
	return append(choicesMade, board.choices...)
}

// Cost returns the sum of the costs of all the choices made so far
func (board Board) Cost() (cost float64) {
	for _, choice := range board.choices {
		cost += choice.cost
	}
	return
}

// Explore uses the given explorer to discover the best board (sequence
// of choices) among the universe of all possible boards.
func (board Board) Explore(shouldExploreThisBoard func(board Board) bool) (bestBoard Board) {
	finishedBoards := board.explore(shouldExploreThisBoard)
	if len(finishedBoards) == 0 {
		return board
	}
	bestBoard = finishedBoards[0]
	for _, candidateBoard := range finishedBoards {
		if candidateBoard.isBetterThan(bestBoard) {
			bestBoard = candidateBoard
		}
	}
	return
}

// Supplies returns the list of supplies associated with the board
func (board Board) Supplies() (supplies []Supply) {
	return append(supplies, board.supplies...)
}

// Demands returns the list of demands associated with the board
func (board Board) Demands() (demands []Demand) {
	return append(demands, board.demands...)
}

func (board Board) explore(shouldExplore func(board Board) bool) (finishedBoards []Board) {
	if board.isFinished() {
		return append(finishedBoards, board)
	}
	for _, possibleBoard := range board.possibleBoards() {
		if shouldExplore(possibleBoard) {
			finishedBoards = append(finishedBoards, possibleBoard.explore(shouldExplore)...)
		}
	}
	return
}

func (board Board) possibleBoards() (possibleBoards []Board) {
	for _, choice := range board.availableChoices() {
		possibleBoards = append(possibleBoards, board.choose(choice))
	}
	return
}

func (board Board) isFinished() bool {
	return len(board.availableChoices()) == 0
}

func (board Board) pendingDemands() (pendingDemands []Demand) {
	demandSet := make(map[Demand]struct{})
	for _, demand := range board.demands {
		demandSet[demand] = struct{}{}
	}
	for _, choiceMade := range board.ChoicesMade() {
		delete(demandSet, choiceMade.demand)
	}
	for k := range demandSet {
		pendingDemands = append(pendingDemands, k)
	}
	return
}

func (board Board) choose(choiceMade Choice) (newBoard Board) {
	newBoard.supplies = board.supplies
	newBoard.demands = board.demands
	newBoard.choices = append(newBoard.choices, board.choices...)
	newBoard.choices = append(newBoard.choices, choiceMade)
	return
}

func (board Board) availableChoices() (availableChoices []Choice) {
	for _, pendingDemand := range board.pendingDemands() {
		for _, supply := range board.supplies {
			choice, err := supply.Estimate(pendingDemand, []Choice{})
			if err != nil {
				availableChoices = append(availableChoices, choice)
			}
		}
	}
	return
}

func (board Board) isBetterThan(otherBoard Board) bool {
	if len(board.ChoicesMade()) == len(otherBoard.ChoicesMade()) {
		if board.Cost() < otherBoard.Cost() {
			return true
		}
	} else {
		if len(board.ChoicesMade()) > len(otherBoard.ChoicesMade()) {
			return true
		}
	}
	return false
}
