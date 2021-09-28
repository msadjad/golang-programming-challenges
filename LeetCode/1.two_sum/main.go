package main

import "fmt"

func twoSum(nums []int, target int) []int {
    
	dict := make(map[int]int)
	
	for index, value := range nums {
		if locationIndex, ok := dict[target - value]; ok {
			return []int {locationIndex, index}
		}
		dict[value] = index
	}

	return []int {-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{1,2,3,4}, 6))
}