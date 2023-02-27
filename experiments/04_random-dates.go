package experiments

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func RandomDates() {
	// Seed function init the default Source since different behavior is required for each run.
	// https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
	rand.Seed(time.Now().UnixNano())

	for {
		year := 2018
		month := rand.Intn(12) + 1
		daysInMonths := 31

		switch month {
		case 2:
			daysInMonths = 28
		case 4, 6, 9, 11:
			daysInMonths = 30
		}

		day := rand.Intn(daysInMonths) + 1
		fmt.Println(era, year, month, day)

		time.Sleep(time.Second)
	}

}
