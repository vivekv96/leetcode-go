package main

import "testing"

func Test_subtractProductAndSum(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first case",
			args: args{n: 234},
			want: 15,
		},
		{
			name: "second case",
			args: args{n: 4421},
			want: 21,
		},
		{
			name: "third case",
			args: args{n: 1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractProductAndSum(tt.args.n); got != tt.want {
				t.Errorf("subtractProductAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
