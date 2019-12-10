package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	hashMap := map[int]int{}
	for idx, val := range nums {
		hashMap[val] = idx
	}

	for idx, src := range nums {
		dst := target - src
		_, exist := hashMap[dst]
		if exist && hashMap[dst] != idx {
			res = append(res, idx)
			res = append(res, hashMap[dst])
			break
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	res := []int{}
	hashMap := map[int]int{}
	for idx := len(nums) - 1; idx >= 0; idx-- {
		dst := target - nums[idx]
		_, exist := hashMap[dst]
		if exist {
			res = append(res, idx)
			res = append(res, hashMap[dst])
			break
		} else {
			hashMap[nums[idx]] = idx
		}
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum2(nums, target)

	for _, val := range res {
		fmt.Printf("%v\n", val)
	}
}
