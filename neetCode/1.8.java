import java.util.Arrays;
import java.util.List;

// 设计一个将字符串列表编码为字符串的算法。 已经编码的字符串之后会通过网络发送同时也会被解码回到原始的字符串列表。
// 请实现 encode 和 decode

// 样例
// 给定 strs = ["lint","code","love","you"]
// 返回 "lint:;code:;love:;you"

// link: https://www.lintcode.com/problem/659/


public class Solution {

    public String encode(List<String> strs) {
        StringBuilder sb = new StringBuilder();
        for (String str : strs) {
            for (char c : str.toCharArray()) {
                if (c == ':') {
                    sb.append("::");
                } else {
                    sb.append(c);
                }
            }
            sb.append(":;");
        }
        sb.delete(sb.length() - 2, sb.length());
        return sb.toString();
    }

    public List<String> decode(String str) {
        List<String> result = Arrays.asList(str.split(":;"));
        result.replaceAll(s -> s.replace("::", ":"));
        return result;
    }
}
