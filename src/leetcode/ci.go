package leetcode

import "fmt"

func sol(nums []int) int {
	arr := make([]int, 2*len(nums)+1)

	for index := range arr {
		arr[index] = -2
	}

	arr[len(nums)] = -1

	maxlen := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		var j int
		if nums[i] == 0 {
			j = -1
		} else {
			j = 1
		}
		count = count + j
		if arr[count+len(nums)] >= -1 {
			maxlen = ma(maxlen, i-arr[count+len(nums)])
		} else {
			arr[count+len(nums)] = 1
		}
	}
	return maxlen
}

func Sol() {
	example2 := []int{0, 1}
	output2 := sol(example2)
	fmt.Println(output2)

	example3 := []int{0, 1, 0}
	output3 := sol(example3)
	fmt.Println(output3)

	example4 := []int{1, 1, 0}
	output4 := sol(example4)
	fmt.Println(output4)

	example5 := []int{0, 1, 0, 1}
	output5 := sol(example5)
	fmt.Println(output5)

	example6 := []int{1, 0, 1, 0}
	output6 := sol(example6)
	fmt.Println(output6)

	example7 := []int{0, 1, 0, 1, 1, 0, 1}
	output7 := sol(example7)
	fmt.Println(output7)
}
