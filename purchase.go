package purchase // https://exercism.org/tracks/go/exercises/vehicle-purchase/solutions/rchnmy

import "fmt"

type carPrice struct {
    age, originalPrice, discount float64
}

func(c *carPrice) Discount() {
    switch {
        case c.age < 3:
            c.discount = 20
        case (3 <= c.age) && (c.age < 10):
            c.discount = 30
        case 10 <= c.age:
            c.discount = 50
    }
}

func(c carPrice) Estimate() float64 {
    return c.originalPrice - c.originalPrice / (100 / c.discount)
}

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

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    c := &carPrice{
        age:           age,
        originalPrice: originalPrice,
    }
    c.Discount()
    return c.Estimate()
}
