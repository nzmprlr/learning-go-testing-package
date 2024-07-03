package unit

import (
	mock_unit "gotesting/unit/mocks/unit"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestUnit4_Multiplier(t *testing.T) {
	controller := gomock.NewController(t)
	mockMultiplierService := mock_unit.NewMockMultiplierService(controller)

	type fields struct {
		multiplier MultiplierService
		r          int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"1", fields{mockMultiplierService, 1}, args{2}, 2},
		{"2", fields{mockMultiplierService, 2}, args{3}, 6},
		{"3", fields{mockMultiplierService, 3}, args{4}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMultiplierService.EXPECT().Get().Return(tt.fields.r).Times(1)

			u := Unit4{
				multiplier: tt.fields.multiplier,
			}

			if got := u.Multiplier(tt.args.i); got != tt.want {
				t.Errorf("Unit4.Multiplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
