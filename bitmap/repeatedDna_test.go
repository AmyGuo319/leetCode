package bitmap

import "fmt"

//187. 重复的DNA序列
//所有 DNA 都由一系列缩写为 'A'，'C'，'G' 和 'T' 的核苷酸组成，例如："ACGAATTCCG"。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
//
//编写一个函数来找出所有目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。
//
//
//
//示例 1：
//
//输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//输出：["AAAAACCCCC","CCCCCAAAAA"]
//示例 2：
//
//输入：s = "AAAAAAAAAAAAA"
//输出：["AAAAAAAAAA"]

func findRepeatedDnaSequences(s string) []string {
	mp := make(map[string]int)
	for i := 0; i+10 <= len(s); i++ {
		mp[s[i:i+10]]++
	}
	res := make([]string, 0)
	for k, v := range mp {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

func Example_findRepeatedDnaSequences() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
	//Output:
	//[AAAAACCCCC CCCCCAAAAA]
	//[AAAAAAAAAA]
	//[AAAAAAAAAA]
}
