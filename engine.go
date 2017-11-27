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
	rand.Seed(time.Now().Unix())

	return engine{
		make([]entity, 0),
		attributes,
	}
}

func (this engine) addEntity(label string) {
	this.entities = append(
		this.entities,
		entity{this.attributes, label},
	)
}

func (this engine) update(pool []action, tick int) {
	for i := range this.entities {
		entities[i].update(pool, tick)
	}
}
