package experiments

import "fmt"

// Malacandra finds the travel speed to reach Mars in 28 days,
// provided that it's 56.000.000 Km far away from the Earth
func Malacandra() {
	d := 56 * 1e6
	t := 28
	v := d / (float64)(t*24)
	fmt.Printf("Reaching Malacandra in 28 days travelling at just %f Km/h.\n", v)
}
