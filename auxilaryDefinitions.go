package pinochle

type player interface {
}

type Card struct {
	faceValue string
	suit      string
}

type Hand []Card

type Deck []Card

type meld []Card

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
