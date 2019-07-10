package pinochle

// InitializeMatch will build a Match and return it
func InitializeMatch(pOne, pTwo player, shuffle bool) Match {
	match := Match{
		pointValues: initializePoints(),
		playerOne:   pOne,
		playerTwo:   pTwo,
		deck:        buildDeck(shuffle),
	}

	return match
}

// Match is the pinochle game controller
type Match struct {
	pointValues     points
	playerOne       player
	playerTwo       player
	deck            Deck
	dealerPlayerOne bool
}

// Deal will deal cards to playerOne and playerTwo.
func (match *Match) Deal() {
	for i := 0; i < 4; i++ {
		if match.dealerPlayerOne {
			if card, err := match.deck.pop(); err == nil {
				match.playerTwo.pushToHand(card)
			}

			if card, err := match.deck.pop(); err == nil {
				match.playerTwo.pushToHand(card)
			}

			if card, err := match.deck.pop(); err == nil {
				match.playerTwo.pushToHand(card)
			}

			// ************************************************
			if card, err := match.deck.pop(); err == nil {
				match.playerOne.pushToHand(card)
			}

			if card, err := match.deck.pop(); err == nil {
				match.playerOne.pushToHand(card)
			}

			if card, err := match.deck.pop(); err == nil {
				match.playerOne.pushToHand(card)
			}

		} else {
			if card, err := match.deck.pop(); err == nil {
				match.playerOne.pushToHand(card)
			}

			if card, err := match.deck.pop(); err == nil {
				match.playerOne.pushToHand(card)
			}

			if card, err := match.deck.pop(); err == nil {
				match.playerOne.pushToHand(card)
			}

			// ***********************************************
			if card, err := match.deck.pop(); err == nil {
				match.playerTwo.pushToHand(card)
			}

			if card, err := match.deck.pop(); err == nil {
				match.playerTwo.pushToHand(card)
			}

			if card, err := match.deck.pop(); err == nil {
				match.playerTwo.pushToHand(card)
			}
		}
	}

	match.dealerPlayerOne = !match.dealerPlayerOne
	card, err := match.deck.pop()
	if err != nil {
		match.deck.trump = card
	}
}

// PlayerOneHand accesses the interface method player.getHand()
func (match *Match) PlayerOneHand() []Card {
	return match.playerOne.getHand()
}

// PlayerTwoHand accesses the interface method player.getHand()
func (match *Match) PlayerTwoHand() []Card {
	return match.playerTwo.getHand()
}

/*
func (match *Match) MatchOver() bool {}

func (match *Match) PlayerOneMelds() [][]Card {}

func (match *Match) PlayerTwoMelds() [][]Card {}

func (match *Match) TrickPhase() bool {}

func (match *Match) Playoff() bool {}

func (match *Match) PlayerOneTurn() bool {}

func (match *Match) PlayerOnePlayed(card Card) bool {}

func (match *Match) PlayerTwoPlayed(card Card) bool {}

func (match *Match) CardsPlayedInTrick() [2]Card {}

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
