package leetcode

import "fmt"

func findMaxLen(nums []int) int {
	mp := make(map[int]int)
	mp[0] = -1
	maxlen, count := 0, 0
	for i := 0; i < len(nums); i++ {
		var j int
		if nums[i] == 1 {
			j = 1
		} else {
			j = -1
		}
		count = count + j
		if _, ok := mp[count]; ok {
			maxlen = max(maxlen, i-mp[count])
		} else {
			mp[count] = i
		}
	}
	return maxlen
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func FindMaxLen() {
	example2 := []int{0, 1}
	output2 := findMaxLen(example2)
	fmt.Println(output2)

	example3 := []int{0, 1, 0}
	output3 := findMaxLen(example3)
	fmt.Println(output3)

	example4 := []int{1, 1, 0}
	output4 := findMaxLen(example4)
	fmt.Println(output4)

	example5 := []int{0, 1, 0, 1}
	output5 := findMaxLen(example5)
	fmt.Println(output5)

	example6 := []int{1, 0, 1, 0}
	output6 := findMaxLen(example6)
	fmt.Println(output6)

	example7 := []int{0, 1, 0, 1, 1, 0, 1}
	output7 := findMaxLen(example7)
	fmt.Println(output7)

	example8 := []int{0, 1, 1, 0, 1, 1, 1, 0}
	output8 := findMaxLen(example8)
	fmt.Println(output8)
}
