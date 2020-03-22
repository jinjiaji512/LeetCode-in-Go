package problem0111

import (
	"math"
	"strings"
	"unicode"
)

//link: https://leetcode-cn.com/problems/string-to-integer-atoi/solution/gojie-fa-liang-bu-zou-jian-dan-di-luo-ji-chu-li-zh/

//先处理掉一些非法数值，留下纯数字
//接着对纯数字字符串进行处理，判断是否溢出
func myAtoi(str string) int {
	return convert(clean(str))
}

func clean(s string) (sign int, abs string) {
	//去除首尾空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	//判断是否是有效字符
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
	case '-':
		sign = -1
	case '+':
		sign = 1
	default:
		//非法字符开头
		return
	}
	//遍历字符串字符，出现非数字字符，截断
	rs := []rune(s)
	for i := 1; i < len(rs); i++ {
		if !unicode.IsDigit(rs[i]) {
			rs = rs[:i]
			break
		}
	}
	if sign == 1 || sign == -1 {
		rs = rs[1:]
	} else {
		sign = 1
	}
	abs = string(rs)
	return
}

// 接收的输入是已经处理过的纯数字
func convert(sign int, absStr string) int {
	//数据不合法
	if len(absStr) == 0 {
		return 0
	}
	absNum := 0
	//计算数字，并判断数据是否超出int32边界
	for _, v := range absStr {
		absNum = absNum*10 + int(v-'0')
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && absNum*-1 < math.MinInt32:
			return math.MinInt32
		}
	}
	return absNum * sign
}
