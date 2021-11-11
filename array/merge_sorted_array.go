package array

// 88. Merge Sorted Array
// https://leetcode-cn.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	//mergeByInsertionSort(nums1, m, nums2, n)
	mergeByFilter(nums1, m, nums2, n)
}

func mergeByInsertionSort(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	// insertion sort
	for i := m; i < m+n; i++ {
		k := nums1[i]
		j := i
		for ; j-1 >= 0 && nums1[j-1] > k; j-- {
			nums1[j] = nums1[j-1]
		}
		nums1[j] = k
	}
}

func mergeByFilter(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	i, j := m-1, n-1

	for k := m + n - 1; k >= 0; k-- {
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

}
