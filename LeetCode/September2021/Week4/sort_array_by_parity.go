package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7, 9, 6}))
}

func sortArrayByParityII(nums []int) []int {
	end := len(nums)

	for i, j := 0, 1; i < end-1 && j < end; {
		if nums[j]%2 == 0 && nums[i]%2 == 1 {
			temp := nums[j]
			nums[j] = nums[i]
			nums[i] = temp
		}

		if nums[i]%2 == 0 {
			i += 2
		}

		if nums[j]%2 == 1 {
			j += 2
		}
	}

	return nums
}
