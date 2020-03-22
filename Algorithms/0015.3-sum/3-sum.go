package problem0015

import "sort"

//先排序，nlogn
//对撞指针。
func threeSum(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		iv := nums[i]
		//对i进行去重
		if i > 0 && iv == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			jv := nums[j]
			kv := nums[k]
			if iv+jv+kv < 0 {
				j++
			} else if iv+jv+kv > 0 {
				k--
			} else if j > i+1 && jv == nums[j-1] {
				//对j进行去重
				j++
			} else if k < len(nums)-1 && kv == nums[k+1] {
				//对k进行去重
				j++
			} else {
				results = append(results, []int{iv, jv, kv})
				j++
			}
		}
	}
	return results
}
func threeSum3(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	k := 0
	for k < len(nums)-2 {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			k++
			continue
		}
		firstIndex := 1 + k
		lastIndex := len(nums) - 1
		for firstIndex < lastIndex {
			if nums[k]+nums[firstIndex]+nums[lastIndex] == 0 {
				// if !ifRepeat(res, []int{nums[k], nums[firstIndex], nums[lastIndex]}) {
				// 	res = append(res, []int{nums[k], nums[firstIndex], nums[lastIndex]})
				// }
				res = append(res, []int{nums[k], nums[firstIndex], nums[lastIndex]})
				nextRight := 1
				for {
					if firstIndex < lastIndex-nextRight && nums[lastIndex] == nums[lastIndex-nextRight] {
						nextRight++
					} else {
						break
					}
				}
				lastIndex = lastIndex - nextRight
				nextLeft := 1
				for {
					if firstIndex+nextLeft < lastIndex && nums[firstIndex] == nums[firstIndex+nextLeft] {
						nextLeft++
					} else {
						break
					}
				}
				firstIndex = firstIndex + nextLeft
			} else if nums[k]+nums[firstIndex]+nums[lastIndex] > 0 {
				lastIndex--
			} else {
				firstIndex++
			}
		}
		k++
	}
	return res
}

// n^2 复杂度，超出时间限制
func threeSum2(nums []int) [][]int {
	//map方式缓存两数之和，再遍历查找
	cache := map[int][][]int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if _, ok := cache[sum]; !ok {
				cache[sum] = [][]int{}
			}
			cache[sum] = append(cache[sum], []int{i, j})
		}
	}
	results := [][]int{}
	resultCache := map[[3]int]bool{}
	//遍历查找，并去重
	for i, v := range nums {
		if ss, ok := cache[0-v]; ok {
			for _, s := range ss {
				if i == s[0] || i == s[1] {
					continue
				}
				s3Slice := []int{v, nums[s[0]], nums[s[1]]}
				sort.Ints(s3Slice)
				s3 := [3]int{s3Slice[0], s3Slice[1], s3Slice[2]}
				if _, ok := resultCache[s3]; !ok {
					results = append(results, s3[:])
					resultCache[s3] = true
				}
			}
		}
	}
	return results
}
