package benchmark

import "testing"

func BenchmarkStd(b *testing.B) {
	for range b.N {
		Std()
	}
}

func BenchmarkGoJson(b *testing.B) {
	for range b.N {
		GoJson()
	}
}

func BenchmarkSonic(b *testing.B) {
	for range b.N {
		Sonic()
	}
}
