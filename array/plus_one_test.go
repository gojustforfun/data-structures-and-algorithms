package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "one element 0", args: args{digits: []int{0}}, want: []int{1}},
		{name: "one element 9", args: args{digits: []int{9}}, want: []int{1, 0}},
		{name: "two elements without carray", args: args{digits: []int{2,3}}, want: []int{2, 4}},
		{name: "two elements with carray", args: args{digits: []int{3,9}}, want: []int{4, 0}},
		{name: "more elements without carray", args: args{digits: []int{5, 8, 7, 6}}, want: []int{5, 8, 7, 7}},
		{name: "more elements with carray", args: args{digits: []int{9, 9, 9, 9, 9}}, want: []int{1, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, plusOne(tt.args.digits))
		})
	}
}
