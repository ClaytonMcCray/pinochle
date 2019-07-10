package pinochle

// Human is an object satisfying the player interface.
type Human struct {
	hand []Card
}

func (h *Human) pushToHand(card Card) {
	h.hand = append(h.hand, card)
}

func (h *Human) getHand() []Card {
	return h.hand
}
