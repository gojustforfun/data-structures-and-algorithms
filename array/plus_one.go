package array

// 66. Plus One
// https://leetcode-cn.com/problems/plus-one/

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	reverse(digits)

	carray, v := 1, 0
	for i := 0; i < len(digits) && carray == 1; i++ {
		v = digits[i] + carray
		digits[i], carray = v%10, v/10
	}
	if carray == 1 {
		digits = append(digits, carray)
	}

	reverse(digits)
	
	return digits
}

func reverse(digits []int) {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
}