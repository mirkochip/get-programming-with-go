package experiments

import (
	"fmt"
	"math/rand"
	"time"
)

// Guess a random number 1 <= n <= 100 and will stop running once
// guessed the right number defined as toGuess variable.
func Guess() {
	toGuess := 10
	for {
		guessed := rand.Intn(99) + 1
		fmt.Printf("Given %d, guessed: %d\n", toGuess, guessed)
		time.Sleep(time.Second)
		if guessed == toGuess {
			fmt.Println("Guessed the right number, exiting...")
			break
		}
		if guessed > toGuess {
			fmt.Printf("Wrong guess: %d is greater than the right value.\n", guessed)
		} else {
			fmt.Printf("Wrong guess: %d is smaller than the right value.\n", guessed)
		}
	}
}
