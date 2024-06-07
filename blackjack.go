package blackjack // https://exercism.org/tracks/go/exercises/blackjack/solutions/rchnmy

var cardValue = map[string]int {
    "ace":   11,
    "king":  10,
    "queen": 10,
    "jack":  10,
    "ten":   10,
    "nine":  9,
    "eight": 8,
    "seven": 7,
    "six":   6,
    "five":  5,
    "four":  4,
    "three": 3,
    "two":   2,
}

var result = map[string]string {
    "Tie":   "T",
    "Split": "P",
    "Win":   "W",
    "Stand": "S",
    "Hit":   "H",
}

type player struct {
    Dealer   int
    Self     int
    Hand     string
    Decision string
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    if _, ok := cardValue[card]; !ok {
        return 0
    }
    return cardValue[card]
}

// evalDealer calculates the dealer's cards value.
func(p *player) evalDealer(cards ...string) {
    for _, c := range cards {
        p.Dealer += ParseCard(c)
    }
}

// evalSelf calculates the player's cards value.
func(p *player) evalSelf(cards ...string) {
    for _, c := range cards {
        p.Self += ParseCard(c)
    }
}

// isStandoff checks if there is a tie.
func(p *player) isStandoff() bool {
    switch p.Self {
    case p.Dealer:
        return true
    default:
        p.isNatural()
    }
    return false
}

// isNatural checks if the player has blackjack
// and sets the corresponding hand type.
func(p *player) isNatural() {
    switch p.Self {
    case 22:
        p.Hand = "Unnatural"
    case 21:
        p.Hand = "Natural"
    default:
        p.isSoft()
    }
}

// isSoft is supposed to check if the player has an ace
// and set the corresponding hand type. But Alex's strategy
// is a bit different.
func(p *player) isSoft() {
    switch {
    case 12 <= p.Self:
        p.Hand = "Soft"
    default:
        p.Hand = "Hard"
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    p := &player{}
    p.evalDealer(dealerCard)
    p.evalSelf(card1, card2)

    switch {
    case p.isStandoff():
        p.Decision = result["Tie"]
    default:
        switch p.Hand {
        case "Unnatural":
            p.Decision = result["Split"]
        case "Natural":
            switch {
            case 10 <= p.Dealer:
                p.Decision = result["Stand"]
            default:
                p.Decision = result["Win"]
            }
        case "Soft":
            switch {
            case 16 < p.Self:
                p.Decision = result["Stand"]
            default:
                switch {
                case 7 <= p.Dealer:
                    p.Decision = result["Hit"]
                default:
                    p.Decision = result["Stand"]
                }
            }
        default:
            p.Decision = result["Hit"]
        }
    }
    return p.Decision
}
