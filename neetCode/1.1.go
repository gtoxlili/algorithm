package neetCode

import "sort"

// 给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
//
// 示例 1：
// 输入：nums = [1,2,3,1]
// 输出：true
//
// link: https://leetcode-cn.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	l := len(nums)
	if l <= 1 {
		return false
	}
	if l == 2 {
		return nums[0] == nums[1]
	}
	if !sort.IntsAreSorted(nums) {
		sort.Ints(nums)
	}
	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
