package blackjack

import "math/rand"

type Deck struct {
	cards           []Card
	each_card_count int // 4 if there is one deck, 8 for 2 deck , 12 for 3 ..........
}

//usedCards map[string]int

func (dek Deck) getRandom(noOfCards int) []Card {

	cardsDeck := []Card{Card{name: "A", value1: 1, value2: 11}, Card{name: "2", value1: 2}, Card{name: "3", value1: 3}, Card{name: "4", value1: 4}, Card{name: "5", value1: 5}, Card{name: "6", value1: 6}, Card{name: "7", value1: 7}, Card{name: "8", value1: 8}, Card{name: "9", value1: 9}, Card{name: "J", value1: 10}, Card{name: "Q", value1: 10}, Card{name: "K", value1: 10}}
	each_rand_index := rand.Int31n(12)
	card1 Card := nil
	card2 Card := nil
	if isCardinDeck(cardsDeck[each_rand_index]) {
		card1 = cardsDeck[each_rand_index]
	}
	each_rand_index = rand.Int31n(12)
	if isCardinDeck(cardsDeck[each_rand_index]) {
		card2 = cardsDeck[each_rand_index]
	}
	return []Card{card1, card2}
}

func isCardinDeck(card Card) {

	// add in a map if any card returned more than 4(for 1 deck), 8(2 deck) ......, then return false
	return true

}
