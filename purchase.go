package purchase // https://exercism.org/tracks/go/exercises/vehicle-purchase/solutions/rchnmy

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    switch kind {
        case "car", "truck":
            return true
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var choice string
    if option1 < option2 { choice = option1 } else { choice = option2 }
        return fmt.Sprintf("%s is clearly the better choice.", choice)
}

func CalculateResellPrice(originalPrice, age float64) float64 {
    // Overengeenering for practice sake
    var discount, actualPrice float64
    switch {
        case 3 > age:
            discount += 20
        case (3 <= age) && (age < 10):
            discount += 30
        case 10 <= age:
            discount += 50
    }
    actualPrice = originalPrice - originalPrice / (100 / discount)
    return actualPrice
}
