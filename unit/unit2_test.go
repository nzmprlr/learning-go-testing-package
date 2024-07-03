package unit

import "testing"

func TestUnit2_Multiplier(t *testing.T) {
	type fields struct {
		multiplier int
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
		{"1", fields{1}, args{2}, 2},
		{"2", fields{2}, args{3}, 6},
		{"3", fields{3}, args{4}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Unit2{
				multiplier: tt.fields.multiplier,
			}
			if got := u.Multiplier(tt.args.i); got != tt.want {
				t.Errorf("Unit2.Multiplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
