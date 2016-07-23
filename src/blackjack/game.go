package blackjack

import (
	"fmt"
	"rand"
)

type Card struct {
	name   string
	value1 int
	value2 int
}

type Dealer struct {
	cards []Card
}

type Player struct {
	cards []Card
}

type Deck struct {
	cards []Card
	each_card_count int
}

usedCards := map[string]int

func (dek Deck) getRandom(noOfCards int) []Card {

	cardsDeck = []Card{Card{name: "A", value1: 1, value2: 11}, Card{name: "2", value1: 2}, Card{name: "3", value1: 3}, Card{name: "4", value1: 4}, Card{name: "5", value1: 5}, Card{name: "6", value1: 6}, Card{name: "7", value1: 7}, Card{name: "8", value1: 8}, Card{name: "9", value1: 9}, Card{name: "J", value1: 10}, Card{name: "Q", value1: 10}, Card{name: "K", value1: 10}}
	each_rand_index = rand.Int31n(12)
	if(isCardinDeck(cardsDeck[each_rand_index])) {
		card1 := cardsDeck[each_rand_index]
	}
	each_rand_index = rand.Int31n(12)
	if(isCardinDeck(cardsDeck[each_rand_index])) {
		card2 := cardsDeck[each_rand_index]
	}

	//card1 := Card{name: "2", value: 2}
	//card2 := Card{name: "J", value: 10}
	return []Card{card1, card2}
}

func isCardinDeck(card Card) {
	return true;
	usedCards[card.name] = 
}

type Strategy interface {
	decide(dealer_cards []Card, player_cards []Card, current_turn string)
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
	fmt.Println(g.dealr.cards)
	g.playr.cards = g.dek.getRandom(2)
	fmt.Println(g.playr.cards)
	g.status = "Started"
	g.turn = "Player"
}
