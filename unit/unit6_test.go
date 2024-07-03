package unit

import (
	mock_unit "gotesting/unit/mocks/unit"
	"sync"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestFooAsyncWrong(t *testing.T) {
	mockFooer := mock_unit.NewMockFooer(gomock.NewController(t))
	mockFooer.
		EXPECT().
		Foo()

	FooAsync(mockFooer)
}

func TestFooAsync(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	mockFooer := mock_unit.NewMockFooer(gomock.NewController(t))
	mockFooer.
		EXPECT().
		Foo().
		Do(func() {
			wg.Done()
		})

	FooAsync(mockFooer)
}
