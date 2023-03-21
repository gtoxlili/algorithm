package neetCode

import (
	"unicode"
)

// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
// 字母和数字都属于字母数字字符。
//
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。

// link: https://leetcode-cn.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	r, l := 0, len(s)-1
	for l > r {
		var rc, lc uint8
		for {
			rc = s[r]
			r++
			if r > l {
				goto end
			}
			if unicode.IsLetter(rune(rc)) || unicode.IsNumber(rune(rc)) {
				break
			}
		}
		for {
			lc = s[l]
			l--
			if unicode.IsLetter(rune(lc)) || unicode.IsNumber(rune(lc)) {
				break
			}
		}
		diff := rc - lc
		if diff == 224 && rc == '0' {
			return false
		}
		if rc != lc && diff != 224 && diff != 32 {
			return false
		}
	}
end:
	return true
}
