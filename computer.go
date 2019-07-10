package pinochle

// Computer is an object satisfying the player interface.
type Computer struct {
	hand []Card
}

func (c *Computer) pushToHand(card Card) {
	c.hand = append(c.hand, card)
}

func (c *Computer) getHand() []Card {
	return c.hand
}
