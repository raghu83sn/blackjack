package blackjack

type Strategy interface {
	decide(dealer *Dealer, player *Player, current_turn string)
}

type Game struct {
	dealr  *Dealer
	playr  *Player
	dek    *Deck
	status string
	winner string
	turn   string
}

func (g Game) start() {
	g.dealr.cards = g.dek.getRandom(2)
	g.playr.cards = g.dek.getRandom(2)
	g.status = "Started"
	g.turn = "Player"
}

func (g Game) whatNext() string {
	//use the score, evaluate and return one of these decisions
	return "blackjack"
	//return "playersturn"
	//return "dealersturn"
	//return "playerbusted"
	//return "dealerbusted"
}
