package pinochle

type player interface {
}

// Card is the fundamental type for each playing card.
type Card struct {
	faceValue string
	suit      string
}

// Hand is a slice of cards belonging to each player.
type Hand []Card

// Deck is a slice of cards representing the stack.
type Deck []Card

type meld []Card

var faceValues = []string{"A", "10", "K", "Q", "J", "9"}
var suits = []string{"S", "D", "Q", "H"}

func buildDeck() Deck {
	var deck Deck
	for _, suit := range suits {
		for _, face := range faceValues {
			deck = append(deck, Card{face, suit})
		}
	}

	return deck
}

// Top level points container.
type points struct {
	ace   int
	ten   int
	king  int
	queen int
	jack  int

	classA
	classB
	classC
}

// Construct holding Class A meld points.
type classA struct {
	flush         int
	royalMarriage int
	marriage      int
	dix           int
}

// Construct holding Class B meld points.
type classB struct {
	hundredAces int
	eightyKings int
	sixtyQueens int
	fortyJacks  int
}

// Construct holding Class C meld points.
type classC struct {
	pinochle       int
	doublePinochle int
}
