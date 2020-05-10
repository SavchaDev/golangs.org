package results

import (
	"fmt"
	"math/rand"
)

const dist = 62100000

// FlightToMars sss
func FlightToMars() {
	fmt.Println("Spiceline         Days   Trip type    Price")
	fmt.Println("===========================================")
	for count := 10; count > 0; count-- {
		//Spaceline
		var Spaceline = ""
		if num := rand.Intn(3) + 1; num == 1 {
			Spaceline = "Space Adventures"
		} else if num == 2 {
			Spaceline = "SpaceX"
		} else if num == 3 {
			Spaceline = "Virgin Galactic"
		}
		//days
		speed := (rand.Intn(15) + 16)
		days := dist / speed / 86400
		//price
		price := 20 + speed
		//tripType
		var tripType = ""
		if rand.Intn(2) == 1 {
			tripType = "One-way"
		} else {
			tripType = "Round-trip"
			days *= 2
			price *= 2
		}

		fmt.Printf("%-17v %-6v %-10v %4v $\n", Spaceline, days, tripType, price)
	}
}
