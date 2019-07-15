package pinochle

import "fmt"

// InitializeMatch will build a Match and return it
func InitializeMatch(pOne, pTwo player, shuffle bool, playingTo int) Match {
	match := Match{
		pointValues:     initializePoints(),
		playerOne:       pOne,
		playerTwo:       pTwo,
		deck:            buildDeck(shuffle),
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
	playingTo          int
	mostRecentlyPlayed [2]Card
}

// Deal will deal cards to playerOne and playerTwo.
func (match *Match) Deal() error {
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
	pOneWins := match.playerOne.score() >= match.playingTo
	pTwoWins := match.playerTwo.score() >= match.playingTo
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
func (match *Match) TrickPhase() bool {
	return len(match.deck.stack) > 0
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

// Playoff returns true while either player still has cards.
func (match *Match) Playoff() bool {
	return match.playerOne.hasCards() || match.playerTwo.hasCards()
}

// PlayerOneTurn returns the internal value dealerPlayerOne.
func (match *Match) PlayerOneTurn() bool {
	return match.dealerPlayerOne
}

// PlayerOnePlayed pushes playerOne's card into mostRecentlyPlayed, and
// returns an error if playerOne doesn't have the card in question.
func (match *Match) PlayerOnePlayed(card Card) error {
	if match.handContains(match.playerOne.getHand(), card) {
		match.mostRecentlyPlayed[0] = card
		return nil
	}
	return fmt.Errorf("playerOne's hand doesn't contain card %v", card)
}

// PlayerTwoPlayed pushes playerTwo's card into mostRecentlyPlayed.
func (match *Match) PlayerTwoPlayed(card Card) error {
	if match.handContains(match.playerTwo.getHand(), card) {
		match.mostRecentlyPlayed[1] = card
		return nil
	}
	return fmt.Errorf("playerTwo's hand doesn't contain card %v", card)
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

/*
func (match *Match) DecideTrickWinner() {}

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
