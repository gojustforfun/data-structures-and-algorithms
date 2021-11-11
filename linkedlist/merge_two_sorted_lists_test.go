package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{name: "l1 empty, l2 not", args: args{l1: (*ListNode)(nil), l2: &ListNode{Val: 1}}, want: &ListNode{Val: 1}},
		{name: "l2 empty, l1 not", args: args{l1:&ListNode{Val: 1}, l2:  (*ListNode)(nil)}, want: &ListNode{Val: 1}},
		{name: "l2 == l1 ", args: args{l1:&ListNode{Val: 1}, l2:&ListNode{Val: 1}}, want: &ListNode{Val: 1, Next: &ListNode{Val: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeTwoLists(tt.args.l1, tt.args.l2))
		})
	}
}
