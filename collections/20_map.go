package collections

import (
	"fmt"
	"strings"
)

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
	text := `
	As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, 
	and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight 
	in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same 
	resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled 
	across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering 
	unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man.`
	frequency := prepareText(text)
	for i, val := range frequency {
		if val > 1 {
			fmt.Printf("%v - %v\n", i, val)
		}
	}
}

func prepareText(text string) map[string]int {
	freq := make(map[string]int)

	text = strings.ToLower(text)
	sliceText := strings.Fields(text)
	for _, val := range sliceText {
		val = strings.Trim(val, `.,;:-_!?"`)
		freq[val]++
	}

	return freq
}
