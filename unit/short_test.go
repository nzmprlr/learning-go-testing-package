package unit

import "testing"

func skipIfShort(t *testing.T) {
	t.Helper()
	if testing.Short() {
		t.Skip()
	}
}

func TestShort(t *testing.T) {
	skipIfShort(t)
}
