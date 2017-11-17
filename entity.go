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

// Render action
func (this *entity) do(a action) {
	fmt.Printf("%s: %s\n", this.meta, a.meta)
}
