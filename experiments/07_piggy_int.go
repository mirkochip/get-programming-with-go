package experiments

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomIntCoin() int {
	randomValue := rand.Intn(3)
	randomCoin := 0
	switch randomValue {
	case 0:
		randomCoin = 5
	case 1:
		randomCoin = 10
	case 2:
		randomCoin = 20
	}
	return randomCoin
}

func PiggyInt() {
	var savingsTarget = 20 * 100
	var currentSavings = 0
	for {
		coin := getRandomIntCoin()
		fmt.Printf("Introducing %v USD cents.\n", coin)
		currentSavings += coin
		fmt.Printf("Current balance: %.2f USD.\n", float64(currentSavings)/100)
		time.Sleep(time.Second / 2)
		if currentSavings >= savingsTarget {
			fmt.Printf("Savings target of %v USD just reached.\n", savingsTarget/100)
			break
		}
	}

}
