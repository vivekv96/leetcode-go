package main

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	var (
		t1   TreeNode
		t2   TreeNode
		want TreeNode
	)

	t1.Val = 1
	t1.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}
	t1.Right = &TreeNode{Val: 2}

	t2.Val = 2
	t2.Left = &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}
	t2.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}

	want.Val = 3
	want.Left = &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 4}}
	want.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}

	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "first case",
			args: args{
				t1: &t1,
				t2: &t2,
			},
			want: &want,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
