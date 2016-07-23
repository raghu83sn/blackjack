package blackjack

import (
	"fmt"
	"testing"
)

func TestGameStartGameWithValidCards(t *testing.T) {
	dealr := Dealer{}
	playr := Player{}
	dek := Deck{each_card_count: 4}
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

func TestGamePlayerGettingBlackJackAtStart(t *testing.T) {
	dealr := Dealer{}
	playr := Player{}
	dek := Deck{}
	gme := Game{dealr: &dealr, playr: &playr, dek: &dek}

	for i := 0; i < 100; i++ {
		gme.start()
		what_next_str = gme.whatNext()
		if what_next_str == "blackjack" {
			if gme.playr.cards[0].name == "J" || gme.playr.cards[0].name == "Q" || gme.playr.cards[0].name == "K" && gme.playr.cards[1].name == "A" {
				fmt.Println("Black jack correctly evaluated")
			} else if gme.playr.cards[1].name == "J" || gme.playr.cards[1].name == "Q" || gme.playr.cards[1].name == "K" && gme.playr.cards[0].name == "A" {
				fmt.Println("Black jack correctly evaluated")
			} else {
				t.Errorf("Black jack incorrectly evaluated")
			}

			if gme.winner != "Player" && gme.status != "Ended" {
				t.Errorf("winner and/or status not set correctly")
			}

			break
		}
	}

}

func TestGamePlayerGettingBustedWithHit(t *testing.T) {
	dealr := Dealer{}
	playr := Player{}
	dek := Deck{}
	gme := Game{dealr: &dealr, playr: &playr, dek: &dek}

	for i := 0; i < 100; i++ {
		gme.start()
		what_next_str = gme.whatNext()
		if what_next_str == "playersturn" {
			player_keep_trying = true
			for player_keep_trying {
				if gme.playr.whatNext() == "hit" {
					append(gme.playr.cards, gme.dek.getRandom(1))
					what_next_str = gme.whatNext()
					if what_next_str == "playerbusted" {

						cal_scores = gme.playr.getScore()

						if len(cal_scores) == 1 && cal_scores[0] > 21 {
							fmt.Println("Player Busted Evaluation Done correctly")
						} else if len(cal_scores) > 1 && allValuesGreaterThan(cal_scores, 21) {
							fmt.Println("Player Busted Evaluation Done correctly")
						} else {
							t.Errorf("Player Busted Evaluation Not Done correctly")
						}
						if gme.winner != "Dealer" && gme.status != "Ended" {
							t.Errorf("winner and/or status not set correctly")
						}
						player_keep_trying = false
						break

					} else if what_next_str == "playersturn" {
						continue
					} else {
						player_keep_trying = false
						break
					}

				}
			}
			if !player_keep_trying {
				break
			}

		}
	}

}

// return true only if all values are greater than 21, even if any one is less than or equal to 21 then , user not busted
// this is checked when the there are one or more A in the cards
func allValuesGreaterThan(src []int, elem int) bool {
	all_values_greater = true

	for _, each_value := range src {
		if each_value <= elem {
			all_values_greater = false
			break
		}
	}

	return all_values_greater
}
