package dynamic

import "testing"

func TestClimStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: 20,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimStairs(tt.args.n); got != tt.want {
				t.Errorf("ClimStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
