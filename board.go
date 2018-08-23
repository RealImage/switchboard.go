package switchboard

// Board represents a set of supplies and demands for which a universes of choices can be explored.
type Board struct {
	supplies []Supply
	demands  []Demand
}

// NewBoard constructs a new board with the given supplies and demands
func NewBoard(supplies []Supply, demands []Demand) (board Board) {
	board.supplies = append(board.supplies, supplies...)
	board.demands = append(board.demands, demands...)
	return
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

// Supplies returns the list of supplies associated with the board
func (board Board) Supplies() (supplies []Supply) {
	return append(supplies, board.supplies...)
}

// Demands returns the list of demands associated with the board
func (board Board) Demands() (demands []Demand) {
	return append(demands, board.demands...)
}
