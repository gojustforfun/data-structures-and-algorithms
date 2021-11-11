package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "one element, one zero", args: args{nums: []int{0}}, want: []int{0}},
		{name: "one element, no zero", args: args{nums: []int{1}}, want: []int{1}},
		{name: "two elements, no zero", args: args{nums: []int{3, 2}}, want: []int{3, 2}},
		{name: "two elements, more zeroes", args: args{nums: []int{0, 3, 0, 2, 0}}, want: []int{3, 2, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
