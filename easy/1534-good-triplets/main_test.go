package main

import "testing"

func Test_countGoodTriplets(t *testing.T) {
	type args struct {
		arr []int
		a   int
		b   int
		c   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first case",
			args: args{
				arr: []int{3, 0, 1, 1, 9, 7},
				a:   7,
				b:   2,
				c:   3,
			},
			want: 4,
		},
		{
			name: "second case",
			args: args{
				arr: []int{1, 1, 2, 2, 3},
				a:   0,
				b:   0,
				c:   1,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodTriplets(tt.args.arr, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("countGoodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		i int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first case",
			args: args{i: -10},
			want: 10,
		},
		{
			name: "second case",
			args: args{i: 0},
			want: 0,
		},
		{
			name: "third case",
			args: args{i: -10},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.i); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
