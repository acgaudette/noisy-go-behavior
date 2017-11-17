package main

import (
	"fmt"
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

func (this *entity) solve(pool []action) (result action) {
	min := float64(len(this.state) + 1)

	for _, a := range pool {
		d := a.diff(this.state)

		if d < min {
			min = d
			result = a
		}
	}

	// Debug
	fmt.Printf(
		"[%s] select %s with diff %g\n",
		this.meta,
		result.meta,
		min,
	)

	return
}

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
