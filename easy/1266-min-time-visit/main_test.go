package main

import "testing"

func Test_minTimeToVisitAllPoints(t *testing.T) {
	type args struct {
		points [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first case",
			args: args{[][]int{{1, 1}, {3, 4}, {-1, 0}}},
			want: 7,
		},
		{
			name: "second case",
			args: args{[][]int{{3, 2}, {-2, 2}}},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitAllPoints(tt.args.points); got != tt.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", got, tt.want)
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
