package collatzconjecture // https://exercism.org/tracks/go/exercises/collatz-conjecture/solutions/rchnmy

import (
    "errors"
)

func CollatzConjecture(n int) (int, error) {
    switch {
        case n <= 0:
            return 0, errors.New("The number has to be positive and greater than null.")
        default:
            var v int
            for n != 1 {
                v++
                switch n % 2 == 0 {
                    case true:
                        n /= 2
                    default:
                        n = n * 3 + 1
                }
            }
        return v, nil
    }
}
