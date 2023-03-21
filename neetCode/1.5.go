package neetCode

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
// 示例 1:
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
//
// link: https://leetcode-cn.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)
	for _, num := range nums {
		hashmap[num]++
	}
	ret := make([][2]int, k)
	for key, value := range hashmap {
		// 判断 ret 最后一个元素是否小于当前元素
		if ret[k-1][1] < value {
			ret[k-1][0] = key
			ret[k-1][1] = value
			// 亮点：因为 ret 除了最后一个元素都是有序的，所以只需要遍历一遍
			for i := k - 1; i > 0; i-- {
				if ret[i][1] > ret[i-1][1] {
					ret[i], ret[i-1] = ret[i-1], ret[i]
				}
			}
		}
	}
	retNums := make([]int, k)
	for i := 0; i < k; i++ {
		retNums[i] = ret[i][0]
	}
	return retNums
}
