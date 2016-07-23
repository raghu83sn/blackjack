package blackjack

import (
	"fmt"
	"testing"
)

func TestGameStartGameWithValidCards(t *testing.T) {
	dealr := Dealer{}
	playr := Player{}
	dek := Deck{each_card_count:4}
	gme := Game{dealr: &dealr, playr: &playr, dek: &dek}

	gme.start()
	fmt.Println(gme)
	/*if gme.status != "Started" {
		t.Errorf("Error Starting Game : status")
	}

	if gme.turn != "Player" {
		t.Errorf("Error Starting Game : turn ")
	}*/

	if gme.dealr.cards == nil || len(gme.dealr.cards) != 2 {
		t.Errorf("Error Starting Game : dealr cards")
	}

	if gme.playr.cards == nil && len(gme.playr.cards) != 2 {
		t.Errorf("Error Starting Game : player cards")
	}
}

func TestGamePlayerGettingBlackJack(t *testing.T) {
	dealr := Dealer{}
	playr := Player{}
	dek := Deck{}
	gme := Game{dealr: &dealr, playr: &playr, dek: &dek}

	gme.start()
	fmt.Println(gme)
	/*if gme.status != "Started" {
		t.Errorf("Error Starting Game : status")
	}

	if gme.turn != "Player" {
		t.Errorf("Error Starting Game : turn ")
	}*/

	if gme.dealr.cards == nil || len(gme.dealr.cards) != 2 {
		t.Errorf("Error Starting Game : dealr cards")
	}

	if gme.playr.cards == nil && len(gme.playr.cards) != 2 {
		t.Errorf("Error Starting Game : player cards")
	}
}
