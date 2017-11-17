package main

func main() {
	entities := make([]entity, 1)

	entities[0] = entity{
		make([]attribute, 2),
		"eugene",
	}

	entities[0].state[0] = NewAttribute("happiness")
	entities[0].state[1] = NewAttribute("energy")

	pool := make([]action, 2)

	pool[0] = action{
		[]float64{1, 0},
		"sleep",
	}

	pool[1] = action{
		[]float64{0, 1},
		"yell",
	}

	for tick := 0; ; tick++ {
		update(entities, pool, tick)
	}
}