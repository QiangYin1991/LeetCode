package main

import "fmt"

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Sort(sort.IntSlice(nums))
	for idx, _ := range nums {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}
		if nums[idx] > 0 || (idx < len(nums)-1 && nums[idx]+nums[idx+1] > 0) {
			continue
		}
		left, right := idx+1, len(nums)-1
		for left < right {
			sum := nums[idx] + nums[left] + nums[right]
			if sum == 0 {
				var arr []int
				arr = append(arr, nums[idx], nums[left], nums[right])
				res = append(res, arr)
				for left+1 < right && nums[left] == nums[left+1] {
					left++
				}
				for right-1 > left && nums[right] == nums[right-1] {
					right--
				}
				left, right = left+1, right-1
			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{}
	//nums := []int{0, 0, 0, 0}
	//nums := []int{3, 0, -2, -1, 1, 2}
	//nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	res := threeSum(nums)
	fmt.Println("[")
	for _, arr := range res {
		fmt.Print("[")
		for _, val := range arr {
			fmt.Print(val, ",")
		}
		fmt.Println("],")
	}
	fmt.Println("]")
}
