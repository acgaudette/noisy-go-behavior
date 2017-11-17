package main

func solve(state []attribute, pool []action) (result action) {
	min := float64(len(state) + 1)

	for _, a := range pool {
		d := a.diff(state)

		if d < min {
			min = d
			result = a
		}
	}

	return
}

func predict(state []attribute, tick int) {
	for _, a := range state {
		a.prediction = a.noise.Eval2(float64(tick), a.value)
	}
}

func update(entities []entity, pool []action, tick int) {
	for _, e := range entities {
		result := solve(e.state, pool)
		e.do(result)

		e.updateValues(result.cost)
		predict(e.state, tick)
	}
}
