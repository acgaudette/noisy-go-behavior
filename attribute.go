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

func NewAttribute(meta string) attribute {
	return attribute{
		0,
		0,
		opensimplex.NewWithSeed(rand.Int63()),
		meta,
	}
}
