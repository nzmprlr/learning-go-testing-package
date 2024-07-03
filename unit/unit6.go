package unit

import "time"

//go:generate mockgen -source=$GOFILE -destination=./mocks/$GOPACKAGE/${GOFILE}mock.go
type Fooer interface {
	Foo()
}

func FooAsync(fooer Fooer) {
	go func() {
		time.Sleep(100 * time.Millisecond)
		fooer.Foo()
	}()
}
