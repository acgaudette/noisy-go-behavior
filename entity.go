package main

import (
	"fmt"
	"math"
)

type entity struct {
	state []attribute
	meta  string
}

func (this *entity) updateValues(values []float64) {
	for i := range this.state {
		this.state[i].value = values[i]

		// Debug
		fmt.Printf(
			"\t%s.%s = %g\n",
			this.meta,
			this.state[i].meta,
			this.state[i].value,
		)
	}
}

// Predict state for next tick
func (this *entity) predict(tick int) {
	for i := range this.state {
		this.state[i].predict(tick)

		// Debug
		fmt.Printf(
			"\t%s.%s -> %g\n",
			this.meta,
			this.state[i].meta,
			this.state[i].prediction,
		)
	}
}

// Get difference between state prediction and actual state
func (this *entity) diff() []float64 {
	result := make([]float64, len(this.state))

	for i := range this.state {
		result[i] = math.Abs(this.state[i].prediction - this.state[i].value)
	}

	return result
}

// Select the best action from a pool with the state prediction
func (this *entity) solve(pool []action) (result action) {
	// Calculate diff target to match against
	target := this.diff()
	min := float64(len(target) + 1)

	// Look for action with best fit
	for _, a := range pool {
		d := a.error(target)

		if d < min {
			min = d
			result = a
		}
	}

	// Debug
	fmt.Printf(
		"[%s] select %s with error %g\n",
		this.meta,
		result.meta,
		min,
	)

	return
}

// Solve for and perform action; predict next tick
func (this *entity) update(pool []action, tick int) {
	result := this.solve(pool)
	this.do(result)
	this.updateValues(result.cost)

	this.predict(tick)
}

// Render action
func (this *entity) do(a action) {
	fmt.Printf("%s: %s\n", this.meta, a.meta)
}
