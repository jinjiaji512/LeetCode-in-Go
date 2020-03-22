package problem0111

import (
	"fmt"
	"github.com/jinjiaji512/LeetCode-in-Go/kit"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0111(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		s   []int
		ans int
	}{

		{
			[]int{3, 9, 20, kit.NULL, kit.NULL, 15, 7},
			2,
		},

		{
			[]int{1, 2, kit.NULL},
			1,
		},

		{
			[]int{1},
			1,
		},

		{
			[]int{},
			0,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, minDepth(kit.Ints2TreeNode(tc.s)), "输入:%v", tc)
	}
}
