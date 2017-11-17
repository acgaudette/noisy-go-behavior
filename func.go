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

func predict(e entity, tick int) {
	for i := range e.state {
		// Rough scaling
		e.state[i].prediction = 0.5 + 0.5*e.state[i].noise.Eval2(
			float64(tick)*DELTA_SCALE,
			e.state[i].value,
		)

		// Debug
		fmt.Printf(
			"\t%s.%s -> %g\n",
			e.meta,
			e.state[i].meta,
			e.state[i].prediction,
		)
	}
}

func update(entities []entity, pool []action, tick int) {
	for i := range entities {
		result := solve(entities[i], pool)
		entities[i].do(result)
		entities[i].updateValues(result.cost)
		predict(entities[i], tick)
	}
}
