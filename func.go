package main

import (
	"fmt"
)

func solve(e entity, pool []action) (result action) {
	min := float64(len(e.state) + 1)

	for _, a := range pool {
		d := a.diff(e.state)

		if d < min {
			min = d
			result = a
		}
	}

	// Debug
	fmt.Printf(
		"[%s] select %s with diff %g\n",
		e.meta,
		result.meta,
		min,
	)

	return
}

func predict(state []attribute, tick int) {
	for i := range state {
		// Rough scaling
		state[i].prediction = 0.5 + 0.5*state[i].noise.Eval2(
			float64(tick),
			state[i].value,
		)
	}
}

func update(entities []entity, pool []action, tick int) {
	for i := range entities {
		result := solve(entities[i], pool)
		entities[i].do(result)
		entities[i].updateValues(result.cost)
		predict(entities[i].state, tick)
	}
}
