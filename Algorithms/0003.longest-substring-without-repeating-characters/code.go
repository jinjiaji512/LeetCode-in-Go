package problem0003

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 0
	cache := map[byte]int{} // 存储字符在字符串中的最后位置
	//遍历，如果有重复，起点移动至cache中字符位置的后一字符
	for start, i := 0, 0; i < len(s); i++ {

		//判断是否上一个相同字符位置，是否大于等于start位置，并且i 不等于 start
		//注意 map查找中的ok 与 零值
		if lastCurr, ok := cache[s[i]]; ok && lastCurr >= start && i != start {
			//start位移
			start = lastCurr + 1
		}

		len := i - start + 1
		if maxLen < len {
			maxLen = len
		}

		//记录字符最后出现的位置
		cache[s[i]] = i
	}
	return maxLen
}

/*
用strings.IndexByte代替map缓存
窗口可以在两个边界移动一开始窗口大小为0
随着数组下标的前进窗口的右侧依次增大
每次查询窗口里的字符，若窗口中有查询的字符
窗口的左侧移动到该字符加一的位置
每次记录窗口的最大程度
重复操作直到数组遍历完成
返回最大窗口长度即可
*/
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 0
	//遍历，如果有重复，起点移动至cache中字符位置的后一字符
	for start, i := 0, 0; i < len(s); i++ {
		//判断是否上一个相同字符位置，是否大于等于start位置，并且i 不等于 start
		//注意 map查找中的ok 与 零值
		if lastCurr := strings.IndexByte(s[start:i], s[i]); lastCurr != -1 {
			//start位移
			start = start + lastCurr + 1
		}
		len := i - start + 1
		fmt.Println(i, string(s[i]), start, string(s[start]))
		if maxLen < len {
			maxLen = len
		}
	}
	return maxLen
}

func lengthOfLongestSubstring3(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {
		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}

	return Length
}
