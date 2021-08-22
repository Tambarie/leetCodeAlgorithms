package main

import "fmt"

//Given n non-negative integers a1, a2, …, an , where each represents a point at coordinate (i, ai). n vertical lines are drawn
//such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container,
//such that the container contains the most water.
//
//Note: You may not slant the container and n is at least 2.


//The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
//In this case, the max area of water (blue section) the container can contain is 49.
//Example 1:
//
//
//Input: [1,8,6,2,5,4,8,3,7]
//Output: 49


func maxArea(height []int) int {
	max, start , end := 0,0,len(height)-1

	for start < end {
		width  := start - end
		high := 0

		if height[start] < height[end]{
			high = height[start]- height[end]
			start++
		}else{
			high = height[end] - height[start]
			end--
		}
		temp := width * high
		if temp > max{
			max = temp
		}
	}
	return max

}







func main()  {
height := []int{11,8,6,2,5,4,8,3,7}
fmt.Println(maxArea(height))

}