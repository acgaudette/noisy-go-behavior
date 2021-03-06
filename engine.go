package main

import (
	"math/rand"
	"time"
)

type engine struct {
	entities   []entity
	attributes []attribute
}

func NewEngine(attributes ...attribute) engine {
	// Seed the RNG
	rand.Seed(time.Now().Unix())

	// Use shared attribute specification
	return engine{
		make([]entity, 0),
		attributes,
	}
}

// Create new entity and add to engine
func (this *engine) addEntity(label string) {
	this.entities = append(
		this.entities,
		entity{this.attributes, label},
	)
}

// Update all entities in engine
func (this *engine) update(pool []action, tick int) {
	for i := range this.entities {
		this.entities[i].update(pool, tick)
	}
}
