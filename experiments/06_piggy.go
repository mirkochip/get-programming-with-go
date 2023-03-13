package experiments

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomCoin() float64 {
	randomValue := rand.Intn(3)
	randomCoin := 0.0
	switch randomValue {
	case 0:
		randomCoin = 0.05
	case 1:
		randomCoin = 0.10
	case 2:
		randomCoin = 0.20
	}
	return randomCoin
}

func Piggy() {
	var savingsTarget = 20.0
	var currentSavings = 0.0
	for {
		coin := getRandomCoin()
		fmt.Printf("Introducing %.2f USD.\n", coin)
		currentSavings += coin
		fmt.Printf("Current balance: %.2f USD.\n", currentSavings)
		time.Sleep(time.Second / 2)
		if currentSavings >= savingsTarget {
			fmt.Printf("Savings target of %v USD just reached.\n", savingsTarget)
			break
		}
	}
}
