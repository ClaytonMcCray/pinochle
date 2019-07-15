package pinochle

import (
	"testing"
)

func TestCompareCards(t *testing.T) {
	oneOne := Card{faceValues[0], suits[0]}
	oneTwo := Card{faceValues[0], suits[0]}
	if !CompareCards(oneOne, oneTwo) {
		t.Errorf("%v should equal %v, but failed", oneOne, oneTwo)
	}

	twoOne := Card{faceValues[2], suits[3]}
	twoTwo := Card{faceValues[0], suits[3]}
	if CompareCards(twoOne, twoTwo) {
		t.Errorf("%v should not equal %v, but does", twoOne, twoTwo)
	}

	threeOne := Card{faceValues[1], suits[0]}
	threeTwo := Card{faceValues[1], suits[2]}
	if CompareCards(threeOne, threeTwo) {
		t.Errorf("%v should not equal %v, but does", threeOne, threeTwo)
	}
}

func TestHandContains(t *testing.T) {
	playerOne := Human{}
	playerTwo := Computer{}

	m := InitializeMatch(&playerOne, &playerTwo, false, 1000)
	m.Deal()

	if !m.handContains(m.PlayerTwoHand(), Card{"9", "H"}) {
		t.Errorf("%v contains %v", m.PlayerTwoHand(), Card{"9", "H"})
	}

	if !m.handContains(m.PlayerTwoHand(), Card{"J", "C"}) {
		t.Errorf("%v contains %v", m.PlayerTwoHand(), Card{"J", "C"})
	}

	if m.handContains(m.PlayerTwoHand(), Card{"Q", "H"}) {
		t.Errorf("%v does not contain %v, but playerOne does.", m.PlayerTwoHand(), Card{"Q", "H"})
	}

	if m.handContains(m.PlayerTwoHand(), Card{"K", "D"}) {
		t.Errorf("%v does not contain %v", m.PlayerTwoHand(), Card{"K", "D"})
	}
}

func TestBuildDeckShuffle(t *testing.T) {
	playerOne := Human{}
	playerTwo := Computer{}

	shuffled := InitializeMatch(&playerOne, &playerTwo, true, 100)
	notShuffled := InitializeMatch(&playerOne, &playerTwo, false, 100)

	samePositionCount := 0
	for i := range shuffled.deck.stack {
		if CompareCards(shuffled.deck.stack[i], notShuffled.deck.stack[i]) {
			samePositionCount++
		}
	}

	if samePositionCount > len(shuffled.deck.stack)/4 {
		t.Errorf("greater than a quarter of the deck was unshuffled")
	}
}

func TestDeal(t *testing.T) {
	playerOne := Human{}
	playerTwo := Computer{}

	match := InitializeMatch(&playerOne, &playerTwo, true, 100)

	match.Deal()
	if !(len(match.PlayerOneHand()) == 12 && len(match.PlayerTwoHand()) == 12) {
		t.Errorf("playerOne hand: %v \n playerTwo hand: %v \n Stack: %v", match.PlayerOneHand(), match.PlayerTwoHand(), match.deck.stack)
	}

	if len(match.deck.stack) != 23 {
		t.Errorf("playerOne hand: %v \n playerTwo hand: %v \n Stack: %v", match.PlayerOneHand(), match.PlayerTwoHand(), match.deck.stack)
	}

	if CompareCards(match.deck.trump, DummyCard) {
		t.Error("trump is DummyCard")
	}
}

func TestBuildDeckGeneration(t *testing.T) {

	d := buildDeck(false)
	if len(d.stack) != 48 {
		t.Errorf("deck.stack is the wrong length: %v", d.stack)
	}

	for _, slowIter := range d.stack {
		cardCount := 0
		for _, fastIter := range d.stack {
			if CompareCards(slowIter, fastIter) {
				cardCount++
			}
		}

		if cardCount != 2 {
			t.Errorf("unshuffled: %v wrong number of cards: %v", slowIter, d.stack)
		}
	}

	d = buildDeck(true)
	if len(d.stack) != 48 {
		t.Errorf("deck.stack is the wrong length: %v", d.stack)
	}

	for _, slowIter := range d.stack {
		cardCount := 0
		for _, fastIter := range d.stack {
			if CompareCards(slowIter, fastIter) {
				cardCount++
			}
		}

		if cardCount != 2 {
			t.Errorf("shuffled: %v wrong number of cards: %v", slowIter, d.stack)
		}
	}
}
