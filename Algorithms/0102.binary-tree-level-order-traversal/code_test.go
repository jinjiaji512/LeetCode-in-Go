package problem0102

import (
	"fmt"
	"github.com/jinjiaji512/LeetCode-in-Go/kit"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0102(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		s   []int
		ans [][]int
	}{

		{
			[]int{3, 9, 20, kit.NULL, kit.NULL, 15, 7},
			[][]int{[]int{3},
				[]int{9, 20},
				[]int{15, 7}},
		},
		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, levelOrder(kit.Ints2TreeNode(tc.s)), "输入:%v", tc)
	}
}
