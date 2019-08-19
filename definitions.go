package pinochle

import (
	"errors"
	"math/rand"
	"time"
)

type player interface {
	pushToHand(card Card)
	getHand() []Card
	score() int
	getMelds() [][]Card
	hasCards() bool
	play(card Card) (Card, error)
	scoreTrickPoints(points int)
	scoreMeldPoints(points int)
	mergeMeldsAndTricks()
	meldScore() int
	storeMeld(meld []Card)
}

// Card is the fundamental type for each playing card.
type Card struct {
	faceValue string
	suit      string
}

// DummyCard is a blank place holder for compliance with player interface.
var DummyCard = Card{"", ""}

// Deck is a slice of cards representing the stack. It satisfies the Shuffler interface.
type Deck struct {
	stack []Card
	trump Card
}

func (d *Deck) pop() (Card, error) {
	if len(d.stack) < 1 {
		return DummyCard, errors.New("type Deck has no cards")
	}

	top := d.stack[len(d.stack)-1]
	d.stack = d.stack[:len(d.stack)-1]
	return top, nil
}

func (d *Deck) getTrump() (Card, error) {
	if CompareCards(d.trump, DummyCard) {
		return DummyCard, errors.New("type Deck has no valid trump")
	}

	return d.trump, nil
}

var faceValues = []string{"A", "10", "K", "Q", "J", "9"}
var suits = []string{"S", "D", "C", "H"}
var faceValueRanks = map[string]int{
	faceValues[0]: len(faceValues),
	faceValues[1]: len(faceValues) - 1,
	faceValues[2]: len(faceValues) - 2,
	faceValues[3]: len(faceValues) - 3,
	faceValues[4]: len(faceValues) - 4,
	faceValues[5]: len(faceValues) - 5,
}

// buildDeck generates a Deck; it shuffle determines whether or not the Deck is shuffled.
func buildDeck(shuffle bool) Deck {
	var stack []Card
	for _, suit := range suits {
		for _, face := range faceValues {
			stack = append(stack, Card{face, suit}, Card{face, suit})
		}
	}

	if shuffle {
		var shuffledStack = make([]Card, len(stack))
		rand.Seed(time.Now().UTC().UnixNano())
		perm := rand.Perm(len(stack))
		for srcIdx, destIdx := range perm {
			shuffledStack[destIdx] = stack[srcIdx]
		}

		return Deck{shuffledStack, DummyCard}
	}

	return Deck{stack, DummyCard}
}

func removeCard(hand []Card, idx int) []Card {
	return append(hand[:idx], hand[idx+1:]...)
}

// CompareCards determines whether or not two cards are equivalent. This will be useful
// for evaluating the validity of melds.
func CompareCards(first, second Card) bool {
	return first.faceValue == second.faceValue && first.suit == second.suit
}

func compareCardSlices(first, second []Card) bool {
	if len(first) != len(second) {
		return false
	}

	firstMap := make(map[Card]int)
	secondMap := make(map[Card]int)

	for _, card := range first {
		firstMap[card] = 0
		secondMap[card] = 0
	}

	for _, card := range second {
		firstMap[card] = 0
		secondMap[card] = 0
	}

	for _, card := range first {
		firstMap[card]++
	}

	for _, card := range second {
		secondMap[card]++
	}

	for card := range firstMap {
		if firstMap[card] != secondMap[card] {
			return false
		}
	}

	return true
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

type validMeldSlices struct {
	flush           []Card
	royalMarriage   []Card
	spadeMarriage   []Card
	clubMarriage    []Card
	heartMarriage   []Card
	diamondMarriage []Card
	dix             []Card
	hundredAces     []Card
	eightyKings     []Card
	sixtyQueens     []Card
	fortyJacks      []Card
	pinochle        []Card
	doublePinochle  []Card
}
