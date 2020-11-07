package main

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first case",
			args: args{s: "(1+(2*3)+((8)/4))+1"},
			want: 3,
		},
		{
			name: "second case",
			args: args{s: "(1)+((2))+(((3)))"},
			want: 3,
		},
		{
			name: "third case",
			args: args{s: "1+(2*3)/(2-1)"},
			want: 1,
		},
		{
			name: "fourth case",
			args: args{s: "1"},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.s); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
