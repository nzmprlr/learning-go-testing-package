package fuzz

import "testing"

func FuzzEcho(f *testing.F) {
	f.Add(1)
	f.Add(2)
	f.Fuzz(func(t *testing.T, x int) {
		echo := Echo(x)

		if echo != x {
			t.Errorf("echo want: %d, got: %d", x, echo)
		}
	})
}
