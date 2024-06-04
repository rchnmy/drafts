package cards // https://exercism.org/tracks/go/exercises/card-tricks/solutions/rchnmy

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    var card int
    switch {
        case index < 0 || index > len(slice) - 1:
            card = -1
        default:
            for i, _ := range slice {
                switch i {
                    case index:
                        card = slice[i]
                }
            }
    }
    return card
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    switch {
        case index < 0 || index > len(slice) - 1:
            slice = append(slice, value)
        default:
            slice[index] = value
    }
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    switch {
        case len(values) < 1:
            return slice
        default:
            for _, s := range slice {
                values = append(values, s)	
            }
            return values
    }
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    switch {
        case index < 0 || index > len(slice) - 1:
            return slice
        default:
            return append(slice[:index], slice[index + 1:]...)
    }
}
