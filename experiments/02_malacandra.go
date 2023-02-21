package experiments

import "fmt"

// speedCalculator calculates the speed (Km/h) given a distance (Km) and a duration (hh)
func speedCalculator(distance float64, duration float64) float64 {
	return distance / duration
}

// Malacandra finds the travel speed to reach Mars in 28 days,
// provided that it's 56.000.000 Km far away from the Earth
func Malacandra() {
	d := 56 * 1e6
	t := 28.0
	fmt.Printf("Malacandra is reacheable in 28 days travelling at %f Km/h.\n", speedCalculator(d, t*24))
}
