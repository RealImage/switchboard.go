package switchboard

type Supply interface {
	Estimate(demand Demand, choicesMade []Choice) Choice
}

type Demand interface {
}

type Choice struct {
	Supply     Supply
	Demand     Demand
	Score      float64
	Attributes map[string]interface{}
}

type Explorer interface {
	ShouldExplore(board Board) bool
}

type Board struct {
}

func (b Board) Supplies() []Supply {
	return []Supply{}
}

func (b Board) Demands() []Demand {
	return []Demand{}
}

func (b Board) ChoicesMade() []Choice {
	return []Choice{}
}

func (b Board) Score() float64 {
	return 0
}

func (b Board) FindBestBoardUsing(explorer Explorer) (bestBoard Board) {
	return Board{}
}
