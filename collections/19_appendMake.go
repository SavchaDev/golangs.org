package collections

import "fmt"

// WorkWithSlice test work
func WorkWithSlice() {
	s := []string{}
	lastCap := cap(s)
	for i := 0; i < 1000; i++ {
		s = append(s, "new element")
		if cap(s) != lastCap {
			fmt.Println(cap(s))
			lastCap = cap(s)
		}
	}
}
