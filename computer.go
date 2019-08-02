package pinochle

import "errors"

// Computer is an object satisfying the player interface.
type Computer struct {
	hand              []Card
	currentMeldScore  int
	currentTrickScore int
	currentScore      int
	melds             [][]Card
}

func (c *Computer) scoreTrickPoints(points int) {
	c.currentTrickScore += points
}

func (c *Computer) scoreMeldPoints(points int) {
	c.currentMeldScore += points
}

func (c *Computer) getMelds() [][]Card {
	return c.melds
}

func (c *Computer) hasCards() bool {
	return len(c.hand) > 0
}

func (c *Computer) mergeMeldsAndTricks() {
	c.currentScore = c.currentMeldScore + c.currentTrickScore
}

func (c *Computer) meldScore() int {
	return c.currentMeldScore
}

func (c *Computer) score() int {
	c.mergeMeldsAndTricks()
	return c.currentScore
}

func (c *Computer) pushToHand(card Card) {
	c.hand = append(c.hand, card)
}

func (c *Computer) getHand() []Card {
	return c.hand
}

// play for Computer takes in a DummyCard for the sake of satisfying
// the player interface. It returns the card the Computer chooses to play.
func (c *Computer) play(card Card) (Card, error) {
	if c.hasCards() {
		temp := c.hand[len(c.hand)-1]
		c.hand = removeCard(c.hand, len(c.hand)-1)
		return temp, nil
	}

	return card, errors.New("hand of Computer is empty")
}
