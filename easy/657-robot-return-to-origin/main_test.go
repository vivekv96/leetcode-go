package main

import "testing"

func Test_judgeCircle(t *testing.T) {
	type args struct {
		moves string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first case",
			args: args{moves: "UD"},
			want: true,
		},
		{
			name: "second case",
			args: args{moves: "LL"},
			want: false,
		},
		{
			name: "third case",
			args: args{moves: "RRDD"},
			want: false,
		},
		{
			name: "fourth case",
			args: args{moves: "LDRRLRUULR"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.args.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
