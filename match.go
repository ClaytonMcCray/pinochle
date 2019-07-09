package pinochle

// Match is the pinochle game controller
type Match struct {
	pointValues points
	playerOne   player
	playerTwo   player
}

// Deal will deal cards to playerOne and playerTwo.
func (match *Match) Deal() {}

func (match *Match) MatchOver() bool {}

func (match *Match) PlayerOneHand() Hand {}

func (match *Match) PlayerOneMelds() []meld {}

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

func (match *Match) PlayerOneMeld(attempt meld) bool {}

func (match *Match) PlayerTwoMeld(attempt meld) bool {}

func (match *Match) MeldWasSuccesful() bool {}

func (match *Match) PlayerOneWoneGame() bool {}

func (match *Match) PlayerOneWonMatch() bool {}

func (match *Match) initializePoints() {
	match.pointValues = points{
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
}
