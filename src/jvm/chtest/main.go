package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("接雨水")
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))

	fmt.Println("合并区间")
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}

func trap(height []int) int {
	leftIndex := 0
	rightIndex := len(height) - 1
	totalVolume := 0
	leftMax := height[leftIndex]
	rightMax := height[rightIndex]
	for leftIndex < rightIndex {
		leftMax = max(leftMax, height[leftIndex])
		rightMax = max(rightMax, height[rightIndex])
		if leftMax <= rightMax {
			totalVolume += leftMax - height[leftIndex]
			leftIndex++
		} else {
			totalVolume += rightMax - height[rightIndex]
			rightIndex--
		}
	}
	return totalVolume
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		newStart := intervals[i][0]
		if newStart > end {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		} else {
			end = max(end, intervals[i][1])
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}
