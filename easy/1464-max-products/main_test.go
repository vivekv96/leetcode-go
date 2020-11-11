package main

import "testing"

func Test_maxProduct(t *testing.T) {
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
			args: args{nums: []int{3, 4, 5, 2}},
			want: 12,
		},
		{
			name: "second case",
			args: args{nums: []int{1, 5, 4, 5}},
			want: 16,
		},
		{
			name: "third case",
			args: args{nums: []int{3, 7}},
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
