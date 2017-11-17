package main

import (
	"math"
)

type action struct {
	cost []float64
	meta string
}

func (this *action) diff(state []attribute) (result float64) {
	for i := 0; i < len(this.cost); i++ {
		result += math.Abs(this.cost[i] - state[i].value)
	}

	return
}
