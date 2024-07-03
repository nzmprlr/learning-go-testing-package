package unit_test

import (
	"gotesting/unit"
	"testing"
)

func TestMultiplier(t *testing.T) {
	cases := []struct {
		x, m, want int
	}{
		{1, 2, 2},
		{2, 3, 6},
		{3, 4, 13},
	}

	for _, c := range cases {
		got := unit.Multiplier(c.x, c.m)
		if got != c.want {
			t.Errorf("want: %d, got: %d", c.want, got)
		}
	}
}
