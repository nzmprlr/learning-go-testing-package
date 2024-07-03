package unit

type Unit2 struct {
	multiplier int
}

func (u Unit2) Multiplier(i int) int {
	return i * u.multiplier
}
