package unit

import "time"

type Unit3 struct {
	multiplier Unit3MultiplierService
}

func (u Unit3) Multiplier(i int) int {
	return i * u.multiplier.Get()
}

type Unit3MultiplierService struct {
	multiplier int
}

func (ms Unit3MultiplierService) Get() int {
	time.Sleep(time.Second * 2)
	return ms.multiplier
}
