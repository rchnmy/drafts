package birdwatcher // https://exercism.org/tracks/go/exercises/bird-watcher/solutions/rchnmy

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var b int
    for i := 0; i < len(birdsPerDay); i++ {
        b += birdsPerDay[i]
    }
    return b
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int { 
    var b int
    for i := 7 * week - 7; i < 7 * week; i++ {
        b += birdsPerDay[i]
    }
    return b
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] = birdsPerDay[i] + 1
    }
    return birdsPerDay
}
