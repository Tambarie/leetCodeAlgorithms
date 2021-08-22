package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		checker := target - nums[i]
		if _,ok := m[checker];ok{
			return []int{m[checker],i}
		}
		m[nums[i]] = i
	}

	return nil
}



func main() {
	var givenNums = []int{2,7,11,15,12}
	target := 18
	fmt.Println(twoSum(givenNums,target))
}