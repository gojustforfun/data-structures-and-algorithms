package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{name:"nil", args: args{head: (*ListNode)(nil)}, want: (*ListNode)(nil)},
		{name:"empty", args: args{head: &ListNode{}}, want: &ListNode{}},
		{name:"one element", args: args{head: &ListNode{Val: 1}}, want: &ListNode{Val:1}},
		{name:"two elements", args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, want: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseList(tt.args.head))
		})
	}
}
