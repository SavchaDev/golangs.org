package results

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	heidth = 15
)

// Universe in game
type Universe [][]bool

func newUniverse() Universe {
	newU := make(Universe, heidth)
	for i := range newU {
		newU[i] = make([]bool, width)
	}
	return newU
}

func (universe Universe) show() {
	for _, column := range universe {
		for _, val := range column {
			if val == true {
				fmt.Print("*")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
}

func (universe Universe) seed() {
	allLiving := (width * heidth) / 4
	for i := 0; i <= allLiving; i++ {
		universe[rand.Intn(heidth)][rand.Intn(width)] = true
	}
}

func (universe Universe) alive(x, y int) bool {
	x = (x + width) % width
	y = (y + heidth) % heidth
	return universe[y][x]
}

func (universe Universe) neighbors(x, y int) int {

	y = (y + heidth) % heidth
	x = (x + width) % width
	amount := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if v == 0 && h == 0 {
				continue
			}
			if universe.alive(x+h-1, y+v-1) {
				amount++
			}
		}
	}
	return amount
}

func (universe Universe) next(x, y int) bool {
	var live bool
	neighbors := universe.neighbors(x, y)
	if universe.alive(x, y) {
		if neighbors < 2 || neighbors > 3 {
			live = false // dead
		}
	} else {
		if neighbors == 3 {
			live = true // live
		}
	}
	return live
}

func step(universe, paralel Universe) {
	for y, column := range universe {
		for x := range column {
			paralel[y][x] = universe.next(x, y)
		}
	}
}

// GameOfLife - Conway's Game of Life
func GameOfLife() {
	universe := newUniverse()
	universe.seed()

	paralel := newUniverse()
	for i := 0; i < 10; i++ {
		step(universe, paralel)
		universe.show()
		time.Sleep(time.Second)
		universe, paralel = paralel, universe
	}
}
