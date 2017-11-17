package main

import (
	"math"
)

type action struct {
	cost []float64
	meta string
}

// Returns difference between action cost and entity state prediction
func (this *action) diff(state []attribute) (result float64) {
	for i := 0; i < len(this.cost); i++ {
		result += math.Abs(state[i].prediction - this.cost[i])
	}

	return
}
