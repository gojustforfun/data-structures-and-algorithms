package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "nil", args: args{head: (*ListNode)(nil)}, want: (*ListNode)(nil)},
		{name: "1 node", args: args{head: &ListNode{}}, want: &ListNode{}},
		{name: "2 nodes, k=1", args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, k: 1}, want: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{name: "2 nodes, k=2", args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, k: 2}, want: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
		{name: "3 nodes, k=2", args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, k: 2}, want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val:3}}}},
		{name: "4 nodes, k=2", args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, k: 2}, want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
		{name: "5 nodes, k=2", args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, k: 2}, want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseKGroup(tt.args.head, tt.args.k))
		})
	}
}
