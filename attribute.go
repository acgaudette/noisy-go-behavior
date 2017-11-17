package main

import (
	"github.com/ojrac/opensimplex-go"
	"math/rand"
)

type attribute struct {
	value      float64
	prediction float64
	noise      *opensimplex.Noise
	meta       string
}

// Predict value for next tick
func (this attribute) predict(tick int) {
	// Rough scaling
	this.prediction = 0.5 + 0.5*this.noise.Eval2(
		float64(tick)*DELTA_SCALE,
		this.value,
	)
}

func NewAttribute(meta string) attribute {
	return attribute{
		0,
		0,
		opensimplex.NewWithSeed(rand.Int63()),
		meta,
	}
}
