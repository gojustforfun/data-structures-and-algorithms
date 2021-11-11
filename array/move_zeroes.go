package array

// 283. Move Zeroes
// https://leetcode-cn.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	// moveZeroesByRecursion(nums, 0, len(nums))
	// moveZeroesByIteration(nums)
	// moveZeroesByTwoPointer(nums)
	moveZeroesByFilterTemplate(nums)
}

func moveZeroesByRecursion(nums []int, lo, hi int) {
	if hi-lo < 2 {
		return
	}

	moveZeroesByRecursion(nums, lo+1, hi)

	for i := lo; nums[i] == 0 && i+1 < hi; i++ {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}

func moveZeroesByIteration(nums []int) {

	n := len(nums)

	if n < 2 {
		return
	}

	for j := n - 2; j >= 0; j-- {
		for i := j; nums[i] == 0 && i+1 < n && nums[i+1] != 0; i++ {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}

func moveZeroesByTwoPointer(nums []int) {
	/*
		使用双指针
		左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
		右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
		注意到以下性质：
		左指针左边均为非零数；
		右指针左边直到左指针处均为零
	*/
	n, left, right := len(nums), 0, 0
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func moveZeroesByFilterTemplate(nums []int) {
	// 如果nums[i]不为0, n和i始终相同即n++后会立即执行i++
	// 只有当nums[i]为0时, n不移动且指向0,而i++后n和i不再相同
	// 当nums[i]再次不为0时,n指向的是0, i指向的是不为0的数,交换两者即可
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n], nums[i] = nums[i], nums[n]
			n++
		}
	}
}
