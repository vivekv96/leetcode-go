package main

import "testing"

func Test_generateTheString(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "first case",
			args: args{n: 5},
		},
		{
			name: "second case",
			args: args{n: 1},
		},
		{
			name: "third case",
			args: args{n: 112},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTheString(tt.args.n); hasOddCount(got) != true || len(got) != tt.args.n {
				t.Errorf("generateTheString()'s count = %v, want %v", got, true)
			}
		})
	}
}

func Test_hasOddCount(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "first case",
			s:    "aabbb",
			want: false,
		},
		{
			name: "second case",
			s:    "aaabbb",
			want: true,
		},
		{
			name: "third case",
			s:    "abbbcdd",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := hasOddCount(tc.s); got != tc.want {
				t.Errorf("hasOddCount() = %v, want %v", got, tc.want)
			}
		})
	}
}

func hasOddCount(s string) bool {
	m := make(map[string]int)

	for i := range s {
		m[string(s[i])] += 1
	}

	for _, value := range m {
		if value%2 != 1 {
			return false
		}
	}

	return true
}
