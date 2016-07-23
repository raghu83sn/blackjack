package blackjack

type Strategy interface {
	decide(dealer *Dealer, player *Player, current_turn string)
}
