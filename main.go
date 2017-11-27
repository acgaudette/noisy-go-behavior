package main

import (
	"fmt"
	"time"
)

func main() {
	// Initialize engine with attributes
	behavior := NewEngine(
		NewAttribute("happiness"),
		NewAttribute("energy"),
	)

	// Add entities
	behavior.addEntity("eugene")
	behavior.addEntity("jurgen")

	// Action pool
	pool := []action{
		action{
			[]float64{1, 0},
			"sleep",
		},
		action{
			[]float64{0, 1},
			"yell",
		},
	}

	// Update loop
	for tick := 0; ; tick++ {
		behavior.update(pool, tick)
		time.Sleep(time.Second)
		fmt.Println()
	}
}
