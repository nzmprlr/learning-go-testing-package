package unit

import "time"

type Unit4 struct {
	multiplier MultiplierService
}

func (u Unit4) Multiplier(i int) int {
	return i * u.multiplier.Get()
}

//go:generate mockgen -source=$GOFILE -destination=./mocks/$GOPACKAGE/${GOFILE}mock.go
type MultiplierService interface {
	Get() int
}

type Unit4MultiplierService struct {
	multiplier int
}

func (ms Unit4MultiplierService) Get() int {
	time.Sleep(time.Second * 2)
	return ms.multiplier
}
