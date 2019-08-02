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

	m := InitializeMatch(&playerOne, &playerTwo, 1000)
	m.dealerPlayerOne = true // assure that deal will be accurate
	m.NewGame(false)

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

	shuffled := InitializeMatch(&playerOne, &playerTwo, 100)
	shuffled.NewGame(true)
	notShuffled := InitializeMatch(&playerOne, &playerTwo, 100)
	notShuffled.NewGame(false)

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

	match := InitializeMatch(&playerOne, &playerTwo, 100)

	match.NewGame(true)
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

func TestDealingCards(t *testing.T) {
	playerOne := Human{}
	playerTwo := Computer{}
	match := InitializeMatch(&playerOne, &playerTwo, 100)

	// capture dealerPlayerOne == false
	for i := 0; i < 2; i++ {

		dealerVal := !match.dealerPlayerOne
		err := match.NewGame(true)
		if dealerVal != match.dealerPlayerOne {
			t.Error("dealerPlayerOne was not inverted")
		}

		if err != nil {
			t.Error(err.Error())
		}

		for trickPhase, lastCard := match.TrickPhase(); trickPhase; trickPhase, lastCard = match.TrickPhase() {
			// In an actual game, here you'd base which played first on match.PlayerOneTurn()
			errPlayPlayerOne := match.PlayerOnePlayed(match.playerOne.getHand()[0])
			errPlayPlayerTwo := match.PlayerTwoPlayed(DummyCard)
			// ------------------------------------------------------------------------------
			if errPlayPlayerOne != nil {
				t.Error("playerOne got an error while playing a card (trick phase) : " +
					errPlayPlayerOne.Error())
			}

			if errPlayPlayerTwo != nil {
				t.Error("playerTwo got an error while playing a card (trick phase) : " +
					errPlayPlayerTwo.Error())
			}

			var errDealPlayerOne, errDealPlayerTwo error
			if !lastCard {
				// In an actual game, whoever draws first should be based on match.PlayerOneWonTrick()
				errDealPlayerOne = match.DealToPlayerOne()
				errDealPlayerTwo = match.DealToPlayerTwo()
			} else {
				errDealPlayerOne = match.DealToPlayerOne()
				errDealPlayerTwo = match.DealTrumpToPlayerTwo()
				// ----------------------------------------------------------------------------------
			}

			if errDealPlayerOne != nil {
				t.Error("playerOne got an error while being dealt a card : " +
					errDealPlayerOne.Error())
			}
			if errDealPlayerTwo != nil {
				t.Error("playerTwo got an error while being dealt a card : " +
					errDealPlayerTwo.Error())
			}
		}

		for match.Playoff() {
			// Base of off match.PlayerOneWonTrick()
			errPlayPlayerOne := match.PlayerOnePlayed(match.playerOne.getHand()[0])
			errPlayPlayerTwo := match.PlayerTwoPlayed(DummyCard)
			// -------------------------------------
			if errPlayPlayerOne != nil {
				t.Error("playerOne got an error while playing a card (playoff) : " +
					errPlayPlayerOne.Error())
			}
			if errPlayPlayerTwo != nil {
				t.Error("playerTwo got an error while playing a card (playoff)")
			}
		}

		if len(match.deck.stack) != 0 {
			t.Errorf("match.stack still has cards: %v", match.deck.stack)
		}

		if !CompareCards(match.deck.trump, DummyCard) {
			t.Errorf("match.trump should be DummyCard but isn't: %v", match.deck.trump)
		}

		if match.playerOne.hasCards() {
			t.Errorf("playerOne has cards in hand but shouldn't: %v", match.PlayerOneHand())
		}

		if match.playerTwo.hasCards() {
			t.Errorf("playerTwo has cards in hand but shouldn't: %v", match.PlayerTwoHand())
		}

		if len(match.playerOne.getMelds()) != 0 {
			t.Errorf("playerOne has melds but shouldn't: %v", match.playerOne.getMelds())
		}

		if len(match.playerTwo.getMelds()) != 0 {
			t.Errorf("playerTwo has melds but shouldn't: %v", match.playerTwo.getMelds())
		}
	}
}

func TestDecideTrickWinner(t *testing.T) {
	pOne := Human{}
	pTwo := Computer{}

	match := InitializeMatch(&pOne, &pTwo, 100)
	match.NewGame(true)

	// test 1 -----------------------------------------
	match.deck.trump = Card{"K", "S"}
	match.playerOneLed = true
	match.storePlayerOneCard(Card{"9", "S"})
	match.storePlayerTwoCard(Card{"A", "H"})
	match.DecideTrickWinner()
	if !match.playerOneWonTrick {
		t.Errorf("playerOne lost trick, but should've won: %v", match.mostRecentlyPlayed)
	}

	if !match.playerOneLed {
		t.Error("playerOne should be trick leader but isn't")
	}
	// ------------------------------------------------

	// test 2 -----------------------------------------
	match.deck.trump = Card{"K", "S"}
	match.playerOneLed = true
	match.storePlayerOneCard(Card{"9", "S"})
	match.storePlayerTwoCard(Card{"9", "S"})
	match.DecideTrickWinner()
	if !match.playerOneWonTrick {
		t.Errorf("playerOne lost trick, but should've won: %v", match.mostRecentlyPlayed)
	}

	if !match.playerOneLed {
		t.Error("playerOne should be trick leader but isn't")
	}
	// ------------------------------------------------

	// test 3 -----------------------------------------
	match.deck.trump = Card{"9", "H"}
	match.playerOneLed = false
	match.storePlayerOneCard(Card{"9", "H"})
	match.storePlayerTwoCard(Card{"10", "H"})
	match.DecideTrickWinner()
	if match.playerOneWonTrick {
		t.Errorf("playerOne won trick, but shouldn't have: %v", match.mostRecentlyPlayed)
	}

	if match.playerOneLed {
		t.Error("playerOne is lead, but shouldn't be")
	}
	// --------------------------------------------------

	// test 4 -------------------------------------------
	match.deck.trump = Card{"A", "S"}
	match.playerOneLed = true
	match.storePlayerOneCard(Card{"10", "H"})
	match.storePlayerTwoCard(Card{"9", "C"})
	match.DecideTrickWinner()
	if !match.playerOneWonTrick {
		t.Errorf("playerOne didn't win trick, but should have: %v", match.mostRecentlyPlayed)
	}

	if !match.playerOneLed {
		t.Error("playerOne isn't lead, but should be")
	}
	// --------------------------------------------------

	// test 5 -------------------------------------------
	match.deck.trump = Card{"A", "S"}
	match.playerOneLed = false
	match.storePlayerOneCard(Card{"10", "H"})
	match.storePlayerTwoCard(Card{"9", "C"})
	match.DecideTrickWinner()
	if match.playerOneWonTrick {
		t.Errorf("playerOne didn't win trick, but should have: %v", match.mostRecentlyPlayed)
	}

	if match.playerOneLed {
		t.Error("playerOne isn't lead, but should be")
	}
	// --------------------------------------------------

	// test 6 -------------------------------------------
	match.deck.trump = Card{"10", "H"}
	match.playerOneLed = true
	match.storePlayerOneCard(Card{"A", "H"})
	match.storePlayerTwoCard(Card{"A", "H"})
	match.DecideTrickWinner()
	if !match.playerOneWonTrick {
		t.Errorf("playerOne didn't win trick, but should have: %v", match.mostRecentlyPlayed)
	}

	if !match.playerOneLed {
		t.Error("playerOne isn't lead, but should be")
	}
	// -------------------------------------------------

	// test 7 ------------------------------------------
	match.deck.trump = Card{"10", "H"}
	match.playerOneLed = true
	match.storePlayerOneCard(Card{"A", "C"})
	match.storePlayerTwoCard(Card{"A", "H"})
	match.DecideTrickWinner()
	if match.playerOneWonTrick {
		t.Errorf("playerTwo didn't win trick, but should have: %v", match.mostRecentlyPlayed)
	}

	if match.playerOneLed {
		t.Error("playerTwo isn't lead, but should be")
	}
	// -------------------------------------------------

	// test 8 ------------------------------------------
	match.deck.trump = Card{"10", "H"}
	match.playerOneLed = true
	match.storePlayerOneCard(Card{"Q", "C"})
	match.storePlayerTwoCard(Card{"A", "C"})
	match.DecideTrickWinner()
	if match.playerOneWonTrick {
		t.Errorf("playerTwo didn't win trick, but should have: %v", match.mostRecentlyPlayed)
	}

	if match.playerOneLed {
		t.Error("playerTwo isn't lead, but should be")
	}
	// --------------------------------------------------

	// test 9 ------------------------------------------
	match.deck.trump = Card{"10", "H"}
	match.playerOneLed = false
	match.storePlayerOneCard(Card{"K", "C"})
	match.storePlayerTwoCard(Card{"Q", "C"})
	match.DecideTrickWinner()
	if !match.playerOneWonTrick {
		t.Errorf("playerOne didn't win trick, but should have: %v", match.mostRecentlyPlayed)
	}

	if !match.playerOneLed {
		t.Error("playerOne isn't lead, but should be")
	}
	// --------------------------------------------------
}
