package main

import "testing"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first case",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "second case",
			args: args{
				head: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: 0,
		},
		{
			name: "third case",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: 1,
		},
		{
			name: "fourth case",
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  0,
						Next: nil,
					},
				},
			},
			want: 0,
		},
		{
			name: "fifth case",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 1,
											Next: &ListNode{
												Val: 1,
												Next: &ListNode{
													Val: 1,
													Next: &ListNode{
														Val: 0,
														Next: &ListNode{
															Val: 0,
															Next: &ListNode{
																Val: 0,
																Next: &ListNode{
																	Val: 0,
																	Next: &ListNode{
																		Val: 0,
																		Next: &ListNode{
																			Val:  0,
																			Next: nil,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: 18880,
		},
		{
			name: "sixth case",
			args: args{
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  0,
							Next: nil,
						},
					},
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
