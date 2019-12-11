package main

import "fmt"

func threeSum(nums []int) [][]int {
	var res [][]int
	hashMap := map[int]int{}
	for idx, dst := range nums {
		for key, _ := range hashMap {
			_, exist := hashMap[-dst-key]
			if exist {
				arr := make([]int, 3)
				arr = append(arr, key)
				arr = append(arr, -dst-key)
				arr = append(arr, dst)
				res = append(res, arr)
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println("[")
	for _, arr := range res {
		fmt.Println("[")
		for _, val := range arr {
			fmt.Print(val, ",")
		}
		fmt.Println("],")
	}
	fmt.Println("]")
}
