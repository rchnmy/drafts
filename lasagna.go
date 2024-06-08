package lasagna // https://exercism.org/tracks/go/exercises/lasagna-master/solutions/rchnmy

import (
    "sync"
)

func PreparationTime(layers []string, minutes int) int {
    switch minutes {
    case 0:
        return len(layers) * 2
    default:
        return len(layers) * minutes
    }
}

func Quantities(layers []string) (int, float64) {
    switch {
    case len(layers) < 1:
        return 0, 0.0
    default:

    var (
        noodles int
        sauce   float64
        wg      sync.WaitGroup
    )

    wg.Add(2)
    go func() {
        defer wg.Done()
        for _, l := range layers {
            switch l {
            case "noodles":
                noodles++
            }
        }
       noodles *= 50
    }()

    go func() {
        defer wg.Done()
        for _, l := range layers {
            switch l {
            case "sauce":
                sauce++
            }
        }
       sauce *= 0.2
    }()
    wg.Wait()

    return noodles, sauce
    }
}

func AddSecretIngredient(friendsList, myList []string) {
    for i, v := range myList {
        switch v {
        case "?":
            myList[i] = friendsList[len(friendsList) - 1]
        }
    }
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    var scaledQuantities []float64
    for _, q := range quantities {
        scaledQuantities = append(scaledQuantities, q * float64(portions) / 2.0)
    }
    return scaledQuantities
}
