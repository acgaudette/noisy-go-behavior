package main

import (
	"math"
)

type action struct {
	cost []float64
	meta string
}

// Returns summed difference between action cost and target slice
func (this *action) error(target []float64) (result float64) {
	for i := range this.cost {
		result += math.Abs(target[i] - this.cost[i])
	}

	return
}
