package neetCode

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
// 示例:
// 输入: s = "anagram", t = "nagaram"
// 输出: true
//
// link: https://leetcode-cn.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	count := 0
	for i := 0; i < ls; i++ {
		count += sabot(s[i]) - sabot(t[i])
	}
	return count == 0
}

// 解题要点：破坏 UTF-8 Char 间的线性关系
func sabot(v uint8) int {
	i := int(v)
	return i * i * i * i
}
