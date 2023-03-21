package neetCode

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
// 示例：
// 输入：strs = ["eat","tea","tan","ate","nat","bat"]
// 输出：[["bat"],["nat","tan"],["ate","eat","tea"]]
//
// link: https://leetcode-cn.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	hashmap := make(map[int][]string)
	for i := 0; i < len(strs); i++ {
		count := 0
		for j := 0; j < len(strs[i]); j++ {
			count += sabotage(strs[i][j])
		}
		hashmap[count] = append(hashmap[count], strs[i])
	}
	ret := make([][]string, len(hashmap))
	for _, v := range hashmap {
		ret = append(ret, v)
	}
	return ret
}

func sabotage(v uint8) int {
	i := int(v)
	return i * i * i * i
}
