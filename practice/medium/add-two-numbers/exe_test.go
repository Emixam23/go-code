package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			name: "test1",
			args: args{
				l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
				l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			name: "test2",
			args: args{
				l1: &ListNode{Val: 0},
				l2: &ListNode{Val: 0},
			},
			want: &ListNode{Val: 0},
		},
		{
			name: "test3",
			args: args{
				l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
				l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
			},
			want: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			assert.NotNil(t, got)

			var gotIT, wantIT = got, tt.want
			for gotIT != nil && wantIT != nil {
				assert.Equalf(t, wantIT.Val, gotIT.Val, "addTwoNumbers(%v, %v) => %v but got %v", tt.args.l1.GetNumber(), tt.args.l2.GetNumber(), tt.want.GetNumber(), got.GetNumber())
				gotIT = gotIT.Next
				wantIT = wantIT.Next
			}
			assert.Nilf(t, gotIT, "addTwoNumbers(%v, %v) => %v but got %v", tt.args.l1.GetNumber(), tt.args.l2.GetNumber(), tt.want.GetNumber(), got.GetNumber())
			assert.Nilf(t, wantIT, "addTwoNumbers(%v, %v) => %v but got %v", tt.args.l1.GetNumber(), tt.args.l2.GetNumber(), tt.want.GetNumber(), got.GetNumber())
		})
	}
}
