package leetcode

import (
	"fmt"
)

func iis(A [][]int, B [][]int) [][]int {

	maxArraySize := getMaxArraySize(A, B)
	arrA := generateBooleanArray(A, maxArraySize)
	arrB := generateBooleanArray(B, maxArraySize)

	collisionArray := make([]bool, maxArraySize)

	for i := range arrA {
		if arrA[i] && arrB[i] == true {
			collisionArray[i] = true
		}
	}
	fmt.Println("arrA", arrA)
	fmt.Println("arrB", arrB)

	fmt.Println(collisionArray)
	intervals := getIntervals(collisionArray)
	fmt.Println(intervals)

	return intervals
}

func generateBooleanArray(arr1 [][]int, maxSize int) []bool {
	arrLen := maxSize
	ArrayOfLenOfArr1 := make([]bool, arrLen+1)

	for _, j := range arr1 {
		upperBound := j[1]
		lowerBound := j[0]
		for lowerBound <= upperBound {
			fmt.Println("Lower: ", lowerBound, "\tUpper: ", upperBound)
			ArrayOfLenOfArr1[lowerBound] = true
			lowerBound++
		}
	}
	return ArrayOfLenOfArr1
}

func getMaxArraySize(Array1 [][]int, Array2 [][]int) int {
	indexArr1 := len(Array1) - 1
	indexArr2 := len(Array2) - 1

	maxValArray1 := Array1[indexArr1][1]
	maxValArray2 := Array2[indexArr2][1]
	if maxValArray1 > maxValArray2 {
		return maxValArray1
	} else {
		return maxValArray2
	}
}

func getIntervals(ca []bool) [][]int {
	beginInterval := 0
	var intervals [][]int
	var started bool

	for index := 0; index < len(ca); index++ {
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

	output := iis(intervalA, intervalB)
	fmt.Println(output)
}
