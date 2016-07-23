package blackjack

type Dealer struct {
	cards []Card
	score int
}

func (d Dealer) getScore() []int {
	//return array of int , if A is in the deck we return 2 values in array, else we return one value
	return []int{0}
}

func (p Player) whatNext() string {
	// return "hit"  based on wha tis appropriate
	// return "stand"
	return "hit"
}
