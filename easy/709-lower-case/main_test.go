package main

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first case",
			args: args{str: "Hello"},
			want: "hello",
		},
		{
			name: "second case",
			args: args{str: "here"},
			want: "here",
		},
		{
			name: "third case",
			args: args{str: "LOVELY"},
			want: "lovely",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.str); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
