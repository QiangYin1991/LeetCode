package main

import "fmt"

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Sort(sort.IntSlice(nums))
	for idx, mid := range nums {
		left, right := 0, len(nums)-1
		for left < idx && right > idx {
			addRes := nums[left] + nums[right]
			if addRes == -mid {
				var arr []int
				arr = append(arr, nums[left])
				arr = append(arr, mid)
				arr = append(arr, nums[right])
				res = append(res, arr)
				left, right = left+1, right-1
			} else if addRes < -mid {
				left++
			} else if addRes > -mid {
				right--
			}
		}
	}
	preArr := []int{}
	var res2 [][]int
	for _, arr := range res {
		if len(preArr) == len(arr) && preArr[0] == arr[0] && preArr[1] == arr[1] && preArr[2] == arr[2] {
			continue
		}
		res2 = append(res2, arr)
		preArr = arr
	}
	return res2
}

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{}
	//nums := []int{0, 0, 0, 0}
	//nums := []int{3, 0, -2, -1, 1, 2}
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
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
