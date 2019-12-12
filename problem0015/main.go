package main

import "fmt"

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Sort(sort.IntSlice(nums))
	for idx, mid := range nums {
		if idx > 0 && nums[idx] != nums[idx-1] {
			left, right := 0, len(nums)-1
			for left < idx && right > idx {
				addRes := nums[left] + nums[right]
				if addRes == -mid {
					var arr []int
					arr = append(arr, nums[left])
					arr = append(arr, mid)
					arr = append(arr, nums[right])
					res = append(res, arr)
					for left+1 < idx && nums[left] == nums[left+1] {
						left++
					}
					for right-1 > idx && nums[right] == nums[right-1] {
						right--
					}
					left, right = left+1, right-1
				} else if addRes < -mid {
					left++
				} else if addRes > -mid {
					right--
				}
			}
		}
	}

	return res
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
