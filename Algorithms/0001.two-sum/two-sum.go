package problem0001

func twoSum(nums []int, target int) []int {
	//暴力法 n^2
	//map缓存 n
	if len(nums) < 2 {
		return []int{}
	}
	cache := map[int]int{}
	for i, v := range nums {
		cache[v] = i
	}
	for i, v := range nums {
		if j, ok := cache[target-v]; ok && j > i {
			return []int{i, j}
		}
	}
	return []int{}
}
