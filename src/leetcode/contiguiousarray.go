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

Another solution might be to split the input array into a slice of slices wherever a 0,0 or a 1,1 is encountered, then
iterate over each element and count the number, this should be an O(n) solution as well (O(2n))
*/

package leetcode

import (
	"fmt"
)

func minmax(array []int) (int, int) {
	var max = array[0]
	var min = array[0]
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

	convertedArray := transformArray(nums)
	for i := 0; i <= arrayBound; i++ {
		if i <= arrayBound-1 {
			if convertedArray[i]+convertedArray[i+1] != 0 {
				countList = append(countList, counter)
				counter = 0
			} else {
				counter += 1
			}
		}
		if i == arrayBound && counter != 0 {
			if convertedArray[i] != convertedArray[i-1] {
				counter += 1
				countList = append(countList)
			}
		}
	}

	/*
		if len(nums)-1 % 2 != 0 {
			if convertedArray[arrayBound-1] + convertedArray[arrayBound] == 0 {
				counter += 1
			}
		}*/

	countList = append(countList, counter)

	_, largestSubValue := minmax(countList)
	return largestSubValue
}

func FindMaxLength() {
	/*
		example1 := []int{0, 0, 1, 0, 1, 0, 1, 1}
		output1 := findMaxLength(example1)
		fmt.Println(output1)


		example10 := []int{1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0,1,0,1,0,1,0,1,0}
		output10 := findMaxLength(example10)
		fmt.Println(output10)

	*/
	example2 := []int{0, 1}
	output2 := findMaxLength(example2)
	fmt.Println(output2)

	example3 := []int{0, 1, 0}
	output3 := findMaxLength(example3)
	fmt.Println(output3)

	example4 := []int{1, 1, 0}
	output4 := findMaxLength(example4)
	fmt.Println(output4)

	example5 := []int{0, 1, 0, 1}
	output5 := findMaxLength(example5)
	fmt.Println(output5)

	example6 := []int{1, 0, 1, 0}
	output6 := findMaxLength(example6)
	fmt.Println(output6)

	example7 := []int{0, 1, 0, 1, 1, 0, 1}
	output7 := findMaxLength(example7)
	fmt.Println(output7)
	//begins with 0, odd number, 6 = Correct
	//begins with 1, odd number, 5 = Correct
}
