package array

// 26. Remove Duplicates from Sorted Array
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	// 过滤器模版,用于解决“数组保序过滤问题”
	// 遍历所有元素 
	n := 0
	for i := 0; i < len(nums); i++ {
		// 要还是不要,判断条件
		// 第一个肯定要(i == 0)
		// 从第二个开始,不重复的要 nums[i] != nums[i-1]
		if i == 0 || nums[i] != nums[i-1]  {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}