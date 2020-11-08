package main

import "testing"

func Test_balancedStringSplit(t *testing.T) {
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
			args: args{s: "RLRRLLRLRL"},
			want: 4,
		},
		{
			name: "second case",
			args: args{s: "RLLLLRRRLR"},
			want: 3,
		},
		{
			name: "third case",
			args: args{s: "LLLLRRRR"},
			want: 1,
		},
		{
			name: "fourth case",
			args: args{s: "RLRRRLLRLL"},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
