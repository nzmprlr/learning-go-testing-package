package unit

import "testing"

func TestUnit3_Multiplier(t *testing.T) {
	t.Parallel()
	type fields struct {
		ms Unit3MultiplierService
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
		{"1", fields{Unit3MultiplierService{1}}, args{2}, 2},
		{"2", fields{Unit3MultiplierService{2}}, args{3}, 6},
		{"3", fields{Unit3MultiplierService{3}}, args{4}, 12},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := Unit3{
				multiplier: tt.fields.ms,
			}
			if got := u.Multiplier(tt.args.i); got != tt.want {
				t.Errorf("Unit3.Multiplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
