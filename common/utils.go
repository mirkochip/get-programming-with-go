package common

import "fmt"

func PrintSymbols(symbol string, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(symbol)
		if i == n-1 {
			fmt.Print("\n")
		}
	}
}
