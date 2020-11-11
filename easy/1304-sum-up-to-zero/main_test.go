package main

import (
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "first case",
			args: args{n: 1},
		},
		{
			name: "second case",
			args: args{n: 11},
		},
		{
			name: "third case",
			args: args{n: 111},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumZero(tt.args.n); sum(got) != 0 {
				t.Errorf("sumZero() = %v, want %v", got, 0)
			}
		})
	}
}

func sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}

	return
}

func Test_sum(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first case",
			args: args{nums: []int{1, 2, 3}},
			want: 6,
		},
		{
			name: "second case",
			args: args{nums: []int{1, 2, 3, -2}},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.nums); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
