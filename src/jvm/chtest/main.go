package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
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
