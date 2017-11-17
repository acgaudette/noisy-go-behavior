package main

import (
	"fmt"
)

type entity struct {
	state []attribute
	meta  string
}

func (this *entity) updateValues(values []float64) {
	for i := 0; i < len(this.state); i++ {
		this.state[i].value = values[i]
	}
}

func (this *entity) do(a action) {
	fmt.Printf("%s: %s\n", this.meta, a.meta)
}
