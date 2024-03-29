package pinochle

import "fmt"

// Human is an object satisfying the player interface.
type Human struct {
	hand              []Card
	currentTrickScore int
	currentMeldScore  int
	currentScore      int
	melds             [][]Card
}

func (h *Human) storeMeld(meld []Card) {
	h.melds = append(h.melds, meld)
}

func (h *Human) scoreTrickPoints(points int) {
	h.currentTrickScore += points
}

func (h *Human) scoreMeldPoints(points int) {
	h.currentMeldScore += points
}

func (h *Human) mergeMeldsAndTricks() {
	h.currentScore = h.currentMeldScore + h.currentTrickScore
}

func (h *Human) meldScore() int {
	return h.currentMeldScore
}

func (h *Human) hasCards() bool {
	return len(h.hand) > 0
}

func (h *Human) getMelds() [][]Card {
	return h.melds
}

func (h *Human) score() int {
	h.mergeMeldsAndTricks()
	return h.currentScore
}

func (h *Human) pushToHand(card Card) {
	h.hand = append(h.hand, card)
}

func (h *Human) getHand() []Card {
	return h.hand
}

func (h *Human) handContains(card Card) (bool, int) {
	for idx, cardInHand := range h.hand {
		if CompareCards(card, cardInHand) {
			return true, idx
		}
	}
	return false, -1
}

// play for Human takes in the card the human wants to play,
// validates it, and returns the card if it's a valid card. If it's
// not a valid card, a non-nil error will be returned.
func (h *Human) play(card Card) (Card, error) {
	success, idx := h.handContains(card)
	if success {
		h.hand = removeCard(h.hand, idx)
		return card, nil
	}
	return DummyCard, fmt.Errorf("hand of Human does not contain card %v", card)
}
