package pinochle

import (
	"errors"
	"fmt"
)

// InitializeMatch will build a Match and return it
func InitializeMatch(pOne, pTwo player, playingTo int) Match {
	match := Match{
		pointValues:     initializePoints(),
		playerOne:       pOne,
		playerTwo:       pTwo,
		playingTo:       playingTo,
		dealerPlayerOne: true,
	}

	return match
}

// Match is the pinochle game controller
type Match struct {
	pointValues        points
	playerOne          player
	playerTwo          player
	deck               Deck
	dealerPlayerOne    bool
	playerOneWonTrick  bool
	playerOneLed       bool
	playingTo          int
	mostRecentlyPlayed [2]Card
	meldSlices         validMeldSlices
}

// NewGame initializes a new game, consisting of a trick phase and a playoff.
func (match *Match) NewGame(shuffle bool) error {
	match.deck = buildDeck(shuffle)
	err := match.deal()
	match.playerOneLed = match.dealerPlayerOne
	match.buildMeldSlices()
	return err
}

func (match *Match) buildMeldSlices() {
	suit := match.deck.trump.suit
	match.meldSlices.flush = []Card{Card{"A", suit}, Card{"10", suit}, Card{"K", suit}, Card{"Q", suit}, Card{"J", suit}}
	match.meldSlices.royalMarriage = []Card{Card{"K", suit}, Card{"Q", suit}}
	match.meldSlices.clubMarriage = []Card{Card{"K", "C"}, Card{"Q", "C"}}
	match.meldSlices.spadeMarriage = []Card{Card{"K", "S"}, Card{"Q", "S"}}
	match.meldSlices.diamondMarriage = []Card{Card{"K", "D"}, Card{"Q", "D"}}
	match.meldSlices.heartMarriage = []Card{Card{"K", "H"}, Card{"Q", "H"}}
	match.meldSlices.dix = []Card{Card{"9", suit}}
	match.meldSlices.hundredAces = []Card{Card{"A", "S"}, Card{"A", "H"}, Card{"A", "C"}, Card{"A", "D"}}
	match.meldSlices.eightyKings = []Card{Card{"K", "S"}, Card{"K", "H"}, Card{"K", "C"}, Card{"K", "D"}}
	match.meldSlices.sixtyQueens = []Card{Card{"Q", "S"}, Card{"Q", "H"}, Card{"Q", "C"}, Card{"Q", "D"}}
	match.meldSlices.fortyJacks = []Card{Card{"J", "S"}, Card{"J", "H"}, Card{"J", "C"}, Card{"Q", "D"}}
	match.meldSlices.pinochle = []Card{Card{"Q", "S"}, Card{"J", "D"}}
	match.meldSlices.doublePinochle = []Card{Card{"Q", "S"}, Card{"J", "D"}, Card{"Q", "S"}, Card{"J", "D"}}
}

// deal will deal cards to playerOne and playerTwo.
func (match *Match) deal() error {
	for i := 0; i < 4; i++ {
		if match.dealerPlayerOne {

			for j := 0; j < 3; j++ {
				card, err := match.deck.pop()
				if err != nil {
					return err
				}

				match.playerTwo.pushToHand(card)
			}

			for j := 0; j < 3; j++ {
				card, err := match.deck.pop()
				if err != nil {
					return err
				}

				match.playerOne.pushToHand(card)
			}

		} else {

			for j := 0; j < 3; j++ {
				card, err := match.deck.pop()
				if err != nil {
					return err
				}

				match.playerOne.pushToHand(card)
			}

			for j := 0; j < 3; j++ {
				card, err := match.deck.pop()
				if err != nil {
					return err
				}

				match.playerTwo.pushToHand(card)
			}
		}
	}

	match.dealerPlayerOne = !match.dealerPlayerOne
	card, err := match.deck.pop()
	if err != nil {
		return err
	}
	match.deck.trump = card
	return nil
}

// PlayerOneHand accesses the interface method player.getHand()
func (match *Match) PlayerOneHand() []Card {
	return match.playerOne.getHand()
}

// PlayerTwoHand accesses the interface method player.getHand()
func (match *Match) PlayerTwoHand() []Card {
	return match.playerTwo.getHand()
}

// MatchOver returns whether or not one of the players but not both has reached the cap.
func (match *Match) MatchOver() bool {
	pOneWins := match.playerOne.meldScore() >= match.playingTo
	pTwoWins := match.playerTwo.meldScore() >= match.playingTo
	return (pOneWins || pTwoWins) && !(pOneWins && pTwoWins)
}

// PlayerOneMelds returns a slice of card slices; the internal slices are the individual melds
func (match *Match) PlayerOneMelds() [][]Card {
	return match.playerOne.getMelds()
}

// PlayerTwoMelds returns a slice of card slices; the internal slices are the individual melds
func (match *Match) PlayerTwoMelds() [][]Card {
	return match.playerTwo.getMelds()
}

// TrickPhase returns true while there are cards in the stack.
func (match *Match) TrickPhase() (trickPhase, lastCard bool) {
	trickPhase = len(match.deck.stack) > 0
	lastCard = len(match.deck.stack) < 2
	return trickPhase, lastCard
}

// DealToPlayerOne deals a card to playerOne
func (match *Match) DealToPlayerOne() error {
	card, err := match.deck.pop()
	if err != nil {
		return err
	}

	match.playerOne.pushToHand(card)
	return nil
}

// DealToPlayerTwo deals a card to playerTwo
func (match *Match) DealToPlayerTwo() error {
	card, err := match.deck.pop()
	if err != nil {
		return err
	}

	match.playerTwo.pushToHand(card)
	return nil
}

// DealTrumpToPlayerOne pushes the trump card to playerOne's hand.
func (match *Match) DealTrumpToPlayerOne() error {
	card, err := match.deck.getTrump()
	if err != nil {
		return err
	}

	match.playerOne.pushToHand(card)
	return nil
}

// DealTrumpToPlayerTwo pushes the trump card to playerOne's hand.
func (match *Match) DealTrumpToPlayerTwo() error {
	card, err := match.deck.getTrump()
	if err != nil {
		return err
	}

	match.playerTwo.pushToHand(card)
	return nil
}

// Playoff returns true while either player still has cards.
func (match *Match) Playoff() bool {
	trickPhase, _ := match.TrickPhase()
	return (match.playerOne.hasCards() || match.playerTwo.hasCards()) && !trickPhase
}

// PlayerOneTurn returns the internal value dealerPlayerOne.
/*
func (match *Match) PlayerOneTurn() bool {
}
*/

func (match *Match) storePlayerOneCard(card Card) {
	match.mostRecentlyPlayed[0] = card
}

func (match *Match) storePlayerTwoCard(card Card) {
	match.mostRecentlyPlayed[1] = card
}

// PlayerOnePlayed validates the card that playerOne wants to play,
// then stores it. An error is returned if playerOne's hand doesn't contain the card.
func (match *Match) PlayerOnePlayed(card Card) error {
	validatedCard, err := match.playerOne.play(card)
	if err != nil {
		return err
	}

	match.storePlayerOneCard(validatedCard)
	return nil
}

// PlayerTwoPlayed validates the card that playerTwo wants to play,
// then stores it. An error is returned if playerTwo's hand doesn't contain the card.
func (match *Match) PlayerTwoPlayed(card Card) error {
	validatedCard, err := match.playerTwo.play(card)
	if err != nil {
		return err
	}

	match.storePlayerTwoCard(validatedCard)
	return nil
}

func (match *Match) handContains(hand []Card, card Card) bool {
	for _, c := range hand {
		if CompareCards(c, card) {
			return true
		}
	}

	return false
}

// CardsPlayedInTrick returns an array of the cards played in the most recent trick.
// It returns an error if match.mostRecentlyPlayed contains one or more DummyCard's.
func (match *Match) CardsPlayedInTrick() ([2]Card, error) {
	if CompareCards(match.mostRecentlyPlayed[0], DummyCard) || CompareCards(match.mostRecentlyPlayed[1], DummyCard) {
		return match.mostRecentlyPlayed, fmt.Errorf("match.mostRecentlyPlayed contains %v", match.mostRecentlyPlayed)
	}

	return match.mostRecentlyPlayed, nil
}

func (match *Match) setNextTrickLeader() {
	match.playerOneLed = match.playerOneWonTrick
}

// DecideTrickWinner sets match.playerOneWonTrick based off of match.mostRecentlyPlayed.
func (match *Match) DecideTrickWinner() error {

	defer match.setNextTrickLeader()

	if CompareCards(DummyCard, match.mostRecentlyPlayed[0]) {
		return errors.New("playerOne's card not properly stored")
	}

	if CompareCards(DummyCard, match.mostRecentlyPlayed[1]) {
		return errors.New("playerTwo's card not properly stored")
	}

	pOneCard := match.mostRecentlyPlayed[0]
	pTwoCard := match.mostRecentlyPlayed[1]
	trump, err := match.deck.getTrump()

	if err != nil {
		return err
	}

	if pOneCard.suit == trump.suit && pTwoCard.suit != trump.suit {
		match.playerOneWonTrick = true
		return nil
	} else if pOneCard.suit != trump.suit && pTwoCard.suit == trump.suit {
		match.playerOneWonTrick = false
		return nil
	} else {
		if match.playerOneLed {
			if pTwoCard.suit != pOneCard.suit {
				match.playerOneWonTrick = true
				return nil
			} else if faceValueRanks[pOneCard.faceValue] >= faceValueRanks[pTwoCard.faceValue] {
				match.playerOneWonTrick = true
				return nil
			} else {
				match.playerOneWonTrick = false
				return nil
			}
		} else {
			if pOneCard.suit != pTwoCard.suit {
				match.playerOneWonTrick = false
				return nil
			} else if faceValueRanks[pTwoCard.faceValue] >= faceValueRanks[pOneCard.faceValue] {
				match.playerOneWonTrick = false
				return nil
			} else {
				match.playerOneWonTrick = true
				return nil
			}
		}
	}
}

// validateMeld determines whether or not a real meld has been played, and then
// returns the correct point value
func (match *Match) validateMeld(attempt []Card) (int, error) {
	var scoredPoints int
	var err error
	if compareCardSlices(attempt, match.meldSlices.flush) {
		scoredPoints = match.pointValues.classA.flush
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.royalMarriage) {
		scoredPoints = match.pointValues.classA.royalMarriage
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.clubMarriage) {
		scoredPoints = match.pointValues.classA.marriage
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.diamondMarriage) {
		scoredPoints = match.pointValues.classA.marriage
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.spadeMarriage) {
		scoredPoints = match.pointValues.classA.marriage
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.heartMarriage) {
		scoredPoints = match.pointValues.classA.marriage
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.dix) {
		scoredPoints = match.pointValues.classA.dix
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.hundredAces) {
		scoredPoints = match.pointValues.classB.hundredAces
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.eightyKings) {
		scoredPoints = match.pointValues.classB.eightyKings
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.sixtyQueens) {
		scoredPoints = match.pointValues.classB.sixtyQueens
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.fortyJacks) {
		scoredPoints = match.pointValues.classB.fortyJacks
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.pinochle) {
		scoredPoints = match.pointValues.classC.pinochle
		err = nil
	} else if compareCardSlices(attempt, match.meldSlices.doublePinochle) {
		scoredPoints = match.pointValues.classC.doublePinochle
		err = nil
	} else {
		scoredPoints = 0
		err = errors.New("attempted meld is invalid")
	}

	return scoredPoints, err
}

/*
func (match *Match) AssignTrickPoints() {}

func (match *Match) PlayerOneWonTrick() bool {}

func (match *Match) DoneMelding() bool {}

func (match *Match) PlayerOneMeldableCards() []Card {}

func (match *Match) PlayerTwoMeldableCards() []Card {}

func (match *Match) PlayerOneMeld(attempt []Card) bool {}

func (match *Match) PlayerTwoMeld(attempt []Card) bool {}

func (match *Match) MeldWasSuccesful() bool {}

func (match *Match) PlayerOneWoneGame() bool {}

func (match *Match) PlayerOneWonMatch() bool {}

*/

func initializePoints() points {
	pointValues := points{
		ace:   11,
		ten:   10,
		king:  4,
		queen: 3,
		jack:  2,
		classA: classA{
			flush:         150,
			royalMarriage: 40,
			marriage:      20,
			dix:           10,
		},
		classB: classB{
			hundredAces: 100,
			eightyKings: 80,
			sixtyQueens: 60,
			fortyJacks:  40,
		},
		classC: classC{
			pinochle:       40,
			doublePinochle: 300,
		},
	}

	return pointValues
}
