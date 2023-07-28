package experiments

import (
	"fmt"
)

const distanceFromEarth = 236e15 // km
const lightSpeed = 299792        // km/s
const lightYear = lightSpeed * 60 * 60 * 24 * 365

func Canis() {
	fmt.Printf("Light Year is: %d km/year \n", lightYear)
	lightYears := distanceFromEarth / lightYear
	fmt.Printf("We need %f years to reach the Canis Galazy travelling at the light speed!\n", lightYears)
}
