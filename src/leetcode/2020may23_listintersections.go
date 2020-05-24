package leetcode

import (
	"fmt"
)

func intervalIntersection(A [][]int, B [][]int) [][]int {
	arrA := generateBooleanArray(A)
	arrB := generateBooleanArray(B)
	/*
		for i := range arrA {
			fmt.Println("Num: ", i, " val: ", arrA[i])

		}*/

	maxArraySize, boundary := getMaxAndMin(arrA, arrB)
	collisionArray := make([]bool, maxArraySize)
	//fmt.Println("Smaller Array has: ", maxArraySize, " Smaller Array has:  ", boundary)
	for index := 0; index < boundary; index++ {
		if arrA[index] && arrB[index] == true {
			collisionArray[index] = true
		}
	}
	intervals := getIntervals(collisionArray)
	fmt.Println(intervals)
	return intervals
}

func generateBooleanArray(arr1 [][]int) []bool {
	arrLen := arr1[len(arr1)-1][1]
	ArrayOfLenOfArr1 := make([]bool, arrLen)

	for index := range arr1 {
		beginning := arr1[index][0]
		end := arr1[index][1]
		if beginning == arrLen-1 {
			break
		}
		for ; beginning <= end; beginning++ {
			ArrayOfLenOfArr1[beginning] = true
		}
	}
	return ArrayOfLenOfArr1
}

func getMaxAndMin(ArrayOfLenOfArr1 []bool, ArrayOfLenOfArr2 []bool) (int, int) {
	max1 := len(ArrayOfLenOfArr1) - 1
	max2 := len(ArrayOfLenOfArr2) - 1
	if max1 > max2 {
		return max1, max2
	} else {
		return max2, max1
	}
}

func getIntervals(ca []bool) [][]int {
	beginInterval := 0
	var intervals [][]int
	var started bool

	for index := 0; index < len(ca); index++ {
		//fmt.Println(index, ca[index])
		if !started && ca[index] {
			started = true
			beginInterval = index
		} else if started && !ca[index] {
			started = false
			innerInterval := [][]int{{beginInterval, index - 1}}
			intervals = append(intervals, innerInterval...)
		}
	}
	if started {
		innerInterval := [][]int{{beginInterval, len(ca) - 1}}
		intervals = append(intervals, innerInterval...)
	}
	return intervals
}

func Main23() {
	intervalA := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	intervalB := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}

	output := intervalIntersection(intervalA, intervalB)
	fmt.Println(output)
}
