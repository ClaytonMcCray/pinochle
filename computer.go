package pinochle

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
