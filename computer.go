package pinochle

import "errors"

// Computer is an object satisfying the player interface.
type Computer struct {
	hand         []Card
	currentScore int
	melds        [][]Card
}

func (c *Computer) getMelds() [][]Card {
	return c.melds
}

func (c *Computer) hasCards() bool {
	return len(c.hand) > 0
}

func (c *Computer) score() int {
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
