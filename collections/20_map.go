package collections

import "fmt"

// TestMap testing
func TestMap() {
	// temperature := map[string]int{
	// 	"Earth": 15,
	// 	"Mars":  -65,
	// 	//	"Moon":  -10,
	// }
	temperature := make(map[string]int, 4)
	temperature["Earth"] = 15

	fmt.Printf("average temperature on earth: %v C\n", temperature["Earth"])
	temperature["Mars"] = -100
	fmt.Println(temperature)
	tempMoon, ok := temperature["Moon"]
	if ok {
		fmt.Println(tempMoon)
	} else {
		fmt.Println("moon not found")
	}
	newTemperature := temperature
	delete(newTemperature, "Moon")
	fmt.Println(temperature)
	fmt.Println(newTemperature)
}

// WordFrequency couts the frequency of words in the text
func WordFrequency() {

}
