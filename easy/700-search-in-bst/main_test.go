package main

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	tree := new(TreeNode)

	tree.Val = 4
	tree.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	tree.Right = &TreeNode{Val: 7}

	type args struct {
		root *TreeNode
		val  int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "first case",
			args: args{
				root: tree,
				val:  12,
			},
			want: nil,
		},
		{
			name: "second case",
			args: args{
				root: tree,
				val:  2,
			},
			want: tree.Left,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
