package pinochle

// Human is an object satisfying the player interface.
type Human struct {
	hand         []Card
	currentScore int
	melds        [][]Card
}

func (h *Human) getMelds() [][]Card {
	return h.melds
}

func (h *Human) score() int {
	return h.currentScore
}

func (h *Human) pushToHand(card Card) {
	h.hand = append(h.hand, card)
}

func (h *Human) getHand() []Card {
	return h.hand
}
