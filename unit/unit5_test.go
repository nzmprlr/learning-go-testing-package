package unit

import (
	"fmt"
	mock_unit "gotesting/unit/mocks/unit"
	"slices"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestUnit5_MapToSlicePrint(t *testing.T) {
	mock := mock_unit.NewMockPrinter(gomock.NewController(t))
	mock.EXPECT().Print(gomock.Any())
	mock.EXPECT().Print(&sliceMatcher{[]int{0, 2, 1, 3, 4}})
	mock.EXPECT().Print(gomock.InAnyOrder([]int{0, 1, 2, 3, 4}))

	u := Unit5{mock}
	u.MapToSlicePrint(5)
	u.MapToSlicePrint(5)
	u.MapToSlicePrint(5)
}

type sliceMatcher struct {
	this []int
}

func (m *sliceMatcher) Matches(that any) bool {
	t, ok := that.([]int)
	if !ok {
		return false
	}

	count := 0
	for _, this := range m.this {
		if slices.Contains(t, this) {
			count++
		}
	}

	return len(t) == count
}

func (m *sliceMatcher) String() string {
	return fmt.Sprintf("has the same elements as %v", m.this)
}
