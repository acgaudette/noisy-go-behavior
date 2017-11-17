package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Entity population
	entities := []entity{
		entity{
			[]attribute{
				NewAttribute("happiness"),
				NewAttribute("energy"),
			},
			"eugene",
		},
	}

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
		for i := range entities {
			entities[i].update(pool, tick)
		}

		time.Sleep(time.Second)
		fmt.Println()
	}
}
