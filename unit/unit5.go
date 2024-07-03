package unit

//go:generate mockgen -source=$GOFILE -destination=./mocks/$GOPACKAGE/${GOFILE}mock.go
type Printer interface {
	Print([]int)
}

type Unit5 struct {
	printer Printer
}

func (u Unit5) MapToSlicePrint(x int) {
	m := map[int]int{}
	for i := range x {
		m[i] = i
	}

	slice := []int{}
	for v := range m {
		slice = append(slice, v)
	}

	u.printer.Print(slice)
}
