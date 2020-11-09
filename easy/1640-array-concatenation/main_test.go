package main

import "testing"

func Test_canFormArray(t *testing.T) {
	type args struct {
		arr    []int
		pieces [][]int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first case",
			args: args{
				arr:    []int{85},
				pieces: [][]int{{85}},
			},
			want: true,
		},
		{
			name: "second case",
			args: args{
				arr:    []int{15, 88},
				pieces: [][]int{{88}, {15}},
			},
			want: true,
		},
		{
			name: "third case",
			args: args{
				arr:    []int{49, 18, 16},
				pieces: [][]int{{16, 18, 49}},
			},
			want: false,
		},
		{
			name: "fourth case",
			args: args{
				arr:    []int{91, 4, 64, 78},
				pieces: [][]int{{78}, {4, 64}, {91}},
			},
			want: true,
		},
		{
			name: "fifth case",
			args: args{
				arr:    []int{1, 3, 5, 7},
				pieces: [][]int{{2, 4, 6, 8}},
			},
			want: false,
		},
		{
			name: "sixth case",
			args: args{
				arr:    []int{91, 4, 64, 1, 78},
				pieces: [][]int{{78}, {4, 64, 3}, {91}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFormArray(tt.args.arr, tt.args.pieces); got != tt.want {
				t.Errorf("canFormArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
