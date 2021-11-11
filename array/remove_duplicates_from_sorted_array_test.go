package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Empty", args: args{}, want: 0},
		{name: "One Element", args: args{nums:[]int{1}}, want: 1},
		{name: "Another One Element", args: args{nums:[]int{2}}, want: 1},
		{name: "Two Elements", args: args{nums:[]int{1, 2}}, want: 2},
		{name: "With Duplicates All", args: args{nums:[]int{1, 1}}, want: 1},
		{name: "With Duplicates Head", args: args{nums:[]int{1, 1, 2}}, want: 2},
		{name: "With Duplicates Middle", args: args{nums:[]int{1, 2, 2, 3}}, want: 3},
		{name: "With Duplicates Tail", args: args{nums:[]int{1, 2, 2}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeDuplicates(tt.args.nums))
		})
	}
}
