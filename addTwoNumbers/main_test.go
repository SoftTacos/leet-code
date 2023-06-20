package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "simple",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				l2: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
		{
			name: "simple",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				l2: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
