//go:build !race

package unit

import "testing"

func TestRace(t *testing.T) {
	Race()
}
