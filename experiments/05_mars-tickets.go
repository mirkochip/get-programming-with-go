package experiments

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mirkochip/get-programming-with-go/common"
)

var tickets = 10
var distance = 62100000

func getProvider() string {
	n := rand.Intn(3) + 1
	provider := ""
	switch n {
	case 1:
		provider = "Virgin Galactic"
	case 2:
		provider = "SpaceX"
	case 3:
		provider = "Space Adventures"
	}
	return provider
}

// getRate returns:
// 0: one-way
// 1: round-trip
func getType() int {
	return rand.Intn(2)
}

// getRate returns a rate 36.000 USD <= r <= 50.000 USD
func getRate(speed int) int {
	rate := 0
	switch {
	case speed >= 16 && speed < 20:
		rate = rand.Intn(4000) + 36000
	case speed >= 20 && speed < 25:
		rate = rand.Intn(10000) + 40000
	case speed >= 25 && speed <= 30:
		rate = rand.Intn(6001) + 50000
	}
	return rate / 1000
}

// getSpeed return speed 16 Km/s <= s <= 30 Km/s
// please note -> rand.Intn(n) where n => [0, n) (half-open interval)
func getSpeed() int {
	return (rand.Intn(15) + 16) * 3600
}

// MarsTickets output example:
/*
   Provider           Dd O/W or R/T     Rate
   ==============================================
   Space Adventures   37 Round/Trip     $  72.000
   SpaceX             32 One/Way        $  47.000
   SpaceX             42 One/Way        $  37.000
   Virgin Galactic    31 Round/Trip     $  90.000
   SpaceX             29 Round/Trip     $  94.000
   Space Adventures   27 One/Way        $  54.000
   Space Adventures   34 Round/Trip     $  94.000
   Virgin Galactic    26 Round/Trip     $ 104.000
   SpaceX             32 One/Way        $  41.000
   Space Adventures   31 One/Way        $  43.000
*/
func MarsTickets() {
	// Seed function init the default Source since different behavior is required for each run.
	// https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%-18v %v %-14v %v\n", "Provider", "Dd", "O/W or R/T", "Rate")
	common.PrintSymbols("=", 46)
	for i := 0; i < tickets; i++ {
		provider := getProvider()
		speed := getSpeed()
		rate := 0
		tripType := getType()
		tripLabel := ""
		if tripType == 0 {
			rate = getRate(speed / 3600)
			tripLabel = "One/Way"
		} else {
			rate = getRate(speed/3600) * 2
			tripLabel = "Round/Trip"
		}
		duration := (distance / speed) / 24
		fmt.Printf("%-18v %v %-14v $%+4v.000\n", provider, duration, tripLabel, rate)
	}
}
