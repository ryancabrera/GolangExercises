/*
Example 1:
	Input: [0,1]
	Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.

Example 2:
	Input: [0,1,0]
	Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

Note: The length of the given binary array will not exceed 50,000.


My Solution:

An O(n) solution:

Iterate over an array, with an index starting at 0, check for the initial valid substring of ordered 0,1, when a valid
substring is discovered, use the index of the initial valid substring to create an inner loop where i=outerindex, j=i+1,
and a counter. For every iteration where the condition does not break, increment the counter. When the inner loop breaks
pass the index to the outer loop and move the index of the outer loop. Proceed until the array has been parsed
*/

package leetcode

import (
	"fmt"
)

func findMaxLength(nums []int) int {
	arrayBound := len(nums) - 1
	countList := []int{}
	counter := 0

	if arrayBound == 1 {
		if nums[0] == nums[1] {
			return 0
		} else {
			return 2
		}
	}
	for index, val := range nums {
		if index == arrayBound {
			break
		} else {
			if val != nums[index+1] {
				counter++
			} else {
				//countList = append(countList, counter)
				countList = append(countList, counter)
				counter = 0
			}
		}
	}

	//sort.Ints(countList)
	max, _ := MinMax(countList)
	return max
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func transformArray(nums []int) []int {
	retarr := []int{}
	for _, value := range nums {
		if value == 0 {
			retarr = append(retarr, -1)
		} else {
			retarr = append(retarr, value)
		}
	}
	return retarr
}

func altApproach(nums []int) int {
	arrayBound := len(nums) - 1
	countList := []int{}
	counter := 0

	if arrayBound == 1 {
		if nums[0] == nums[1] {
			return 0
		} else {
			return 2
		}
	}

	convertedArray := transformArray(nums)
	fmt.Println(convertedArray)
	for i := 0; i < arrayBound-1; i++ {
		if convertedArray[i]+convertedArray[i+1] != 0 {
			countList = append(countList, counter)
			counter = 0
		} else {
			counter += 1
		}
	}
	countList = append(countList, counter)
	fmt.Println(countList)
	return countList[0]
}

func FindMaxLength(nums []int) {
	example1 := []int{0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
	//example3 := altApproach(example2)
	//fmt.Println(example3)

	output2 := altApproach(example1)
	fmt.Println(output2)
	//output1, output2 := findMaxLength(example1), findMaxLength(example3)
	//output3 := findMaxLength(example2)

	//fmt.Println(output1, "\n", output2, "\n", output3)

}
