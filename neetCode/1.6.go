package neetCode

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
//
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

// 示例 1：
// 输入：nums = [1,2,3,4]
// 输出：[24,12,8,6]

// link: https://leetcode-cn.com/problems/product-of-array-except-self/

func productExceptSelf(nums []int) []int {
	l := len(nums)
	// L[i] 表示 nums[0:i] 的乘积
	// R[i] 表示 nums[i:l] 的乘积
	L, R := make([]int, l), make([]int, l)
	L[0] = 1
	R[l-1] = 1
	for i := 1; i < l; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	for i := l - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	for i := 0; i < l; i++ {
		nums[i] = L[i] * R[i]
	}
	return nums
}
