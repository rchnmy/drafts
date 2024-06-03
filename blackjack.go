package blackjack // https://exercism.org/tracks/go/exercises/blackjack/solutions/rchnmy

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int
    lowest := []string{"two", "three", "four", "five", "six", "seven", "eight", "nine"}

    switch card {
        case "ace":
            value = 11
        case "king", "queen", "jack", "ten":
            value = 10
        case "nine", "eight", "seven", "six", "five", "four", "three", "two":
            for i, _ := range lowest {
                switch lowest[i] {
                    case card:
                        value = i + 2
                }
            }
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    playerHand := ParseCard(card1) + ParseCard(card2)
    dealerHand := ParseCard(dealerCard)
    var decision string

    switch {
        case playerHand == 22:
            decision = "P"
        case playerHand == 21:
            switch {
                case dealerHand < 10:
                    decision = "W"
                default:
                    decision = "S"
            }
        case 17 <= playerHand && playerHand <= 20:
            decision = "S"
        case 12 <= playerHand && playerHand <= 16:
            switch {
                case 7 <= dealerHand:
                    decision = "H"
                default:
                    decision = "S"
            }
        case playerHand <= 11:
            decision = "H"
    }
    return decision
}
