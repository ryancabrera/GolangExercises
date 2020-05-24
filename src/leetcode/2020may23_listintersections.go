/*

Ax1-Bx1
Ax1-Bx2
Ax2-Bx1
Ax2-Bx2
Bx1-Ax1
Bx1-Ax2
Bx2-Ax1
Bx2-Ax2

Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
Reminder: The inputs and the desired output are lists of Interval objects, and not arrays or lists.
     0     2          1    5
A = [X_1, X_2], B = [X_1, X_2]
A.X_1 < B.X_1 AND A.X_2 < B.x_2
0-1=-1 YES
0-5=-5
2-1=1 YES
2-5=-3
1-0=1 YES
1-2=-2 YES
5-0=5
5-2=3

     5     10          8    12
A = [X_1, X_2], B = [X_1, X_2]
A.X_1 < B.X_1 AND A.X_2 < B.x_2
5-8=-3
5-12=-7
10-8=2
10-12=-2
8-5=4
8-10=-2
12-5=-7
12-10=2


     13     23          15    24
A = [X_1, X_2], B = [X_1, X_2]
A.X_1 < B.X_1 AND A.X_2 < B.x_2


     24     25          25    26
A = [X_1, X_2], B = [X_1, X_2]
A.X_1 < B.X_1 AND A.X_2 < B.x_2
8,10

Take the Greater of the X1s ands the lesser of the X2s, that is the range IF the condition is met?

X1 if N Condition1,

Above won't work, instead , for another naive approach,
I'll try making two new arrays with N elements where N is the max integer found in either array
then iterate over array_1 and fill where value at index exists in first or second array

Might also be worth creating a map for indexes I've already filled for a future optimized problem

//Declare bool array with N elements where N is the last element in Array[X2][X2]
//Might be able to use fancy slicing functions too optimize in future
ArrayOfLenOfArr1 := [arr1[len(arr1)-1][1]]bool{}
for pair in arr1:
	beginning = arr1[0]
	end = arr1 [1]
	for index=beginning; index > end; index++:
		ArrayOfLenOfArr1[index] = true

Repeat with second set of coordinates
ArrayOfLenOfArr2 := [arr2[len(arr1)-1][1]]bool{}
for pair in arr2:
	beginning = arr2[0]
	end = arr2[1]
	for index=beginning; index > end; index++:
		ArrayOfLenOfArr2[index] = true


//I can get the max values of both previous arrays for boundary checking
getMaxAndMin(ArrayOfLenOfArr1, ArrayOfLenOfArr2) (int, int):
	max1 := ArrayOfLenOfArr1[len(ArrayOfLenOfArr1)-a]
	max2 := ArrayOfLenOfArr2[len(ArrayOfLenOfArr2)-a]
	if max1 > max2:
		return max1, max2
	else:
		return max2, max1

maxArraySize, boundary := getMax(getMax(ArrayOfLenOfArr1, ArrayOfLenOfArr)
collisionArray := [maxCollisionArraySize]bool{}
for index=0; index > boundary; index++:
		if ArrayOfLenOfArr1[index] && ArrayOfLenOfArr1[index] == true:
			collisionArray[index] = true


intervals := getIntervals(collisionArray)
return intervals

getIntervals(ca []bool{}) [][]int{} {
	var beginInterval int
	var endInterval int
	var interval [0]int{}
	var intervals[][]int{}

	for index := 0; index > len(ca); index++
	//Maybe print index for debugging
		if ca[index] == true:
			//set beginInterval back to false after end is found in inner loop
			beginInterval := index
			//begin inner loop
			for endOfInterval := index; endOfInterval > end endOfInterval++:
				//set index to end of interval at end of function
				if ca[endOfInterval] == false:
					endInterval = endOfInterval-1
					interval[0] = {beginInterval, endInterval}
					intervals = apppend(intervals, interval)
					index = endOfInterval
					break
	return intervals
}
Check last element of either of new arrays, create a third array of N elements where N is max of len of
either of the previous two, iterate over arrays but be careful to check bounds since 3rd array may be
larger than one of the previous two

Then we'll need to return an array of arrays which will represent intervals, we can create declare the intervals array
and append to it as needed since we won't know how many intervals there are, then we call a function which iterates over
the collision array and sets the beginning of an interval when it encounters the first value of true, then continues over
a second iteration which breaks when it encounters a false, once it encounters a false, it decrements the index and marks
the end of the inner loop, which the outer loop should jump to the index found by the inner loop, and appends the interval
to the interval array
*/

package leetcode

import (
	"fmt"
)

func intervalIntersection(A [][]int, B [][]int) [][]int {
	/*
		for index, element := range A {
			fmt.Println("Index: ", index, "\tElement: ", element)
			fmt.Println(element[0])
		}*/

	arrA := generateBooleanArray(A)
	arrB := generateBooleanArray(B)

	for i := range arrA {
		fmt.Println("Num: ", i, " val: ", arrA[i])

	}

	maxArraySize, boundary := getMaxAndMin(arrA, arrB)
	collisionArray := make([]bool, maxArraySize)
	fmt.Println("Smaller Array has: ", maxArraySize, " Smaller Array has:  ", boundary)
	for index:=0; index < boundary; index++ {
		if arrA[index] && arrB[index] == true {
			collisionArray[index] = true
		}
	}

	fmt.Println(collisionArray)

	intervals := getIntervals(collisionArray)
	fmt.Println(len(arrA))
	fmt.Println(len(arrB))
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
			//fmt.Println("Beginning: ", beginning, " End: ", end)
			break
		}
		for ; beginning <= end; beginning++ {
			//fmt.Println("Beginning: ", beginning, " End: ", end)
			ArrayOfLenOfArr1[beginning] = true
		}
	}
	return ArrayOfLenOfArr1
}

func getMaxAndMin(ArrayOfLenOfArr1 []bool, ArrayOfLenOfArr2 []bool) (int, int){
	max1 := len(ArrayOfLenOfArr1)-1
	max2 := len(ArrayOfLenOfArr2)-1
	if max1 > max2 {
		return max1, max2
	} else {
		return max2, max1
	}
}

/*
func intersection(a []bool, b []bool) (inter []string) {
	// interacting on the smallest list first can potentailly be faster...but not by much, worse case is the same
	low, high := a, b
	if len(a) > len(b) {
		low = b
		high = a
	}

	done := false
	for i, l := range low {
		for j, h := range high {
			// get future index values
			f1 := i + 1
			f2 := j + 1
			if l == h {
				inter = append(inter, h)
				if f1 < len(low) && f2 < len(high) {
					// if the future values aren't the same then that's the end of the intersection
					if low[f1] != high[f2] {
						done = true
					}
				}
				// we don't want to interate on the entire list everytime, so remove the parts we already looped on will make it faster each pass
				high = high[:j+copy(high[j:], high[j+1:])]
				break
			}
		}
		// nothing in the future so we are done
		if done {
			break
		}
	}
	return
}*/

func getIntervals(ca []bool) [][]int {
	beginInterval := 0
	endInterval := 0
	interval := make([]int,2)
	var intervals [][]int
	//var started bool

	for index := 0; index < len(ca)-1; index++ {
		started := index
		for ca[jindex] == true &&
	}
		for index := 0; index < len(ca)-1; index++ {
		if ca[index] == true {
			beginInterval += 1
			for jindex := index; jindex < len(ca)-1; jindex++{
				if ca[jindex] == true{
					continue
				} else if ca[jindex] == false {
					endInterval = jindex - 1
					interval[0], interval[1] = beginInterval, endInterval
					intervals = append(intervals, interval)
					break
				}
			}
		}
	}
	return intervals

		/*
		for index := 0; index < len(ca)-1; index++ {
			if ca[index] == true && started == false && index != len(ca)-1 {
				started = true
				beginInterval += 1
				endInterval += 1
			} else if ca[index] == true && started == false && index == len(ca)-1{
				endInterval = index
				interval[0], interval[1] = beginInterval, endInterval
				intervals = append(intervals, interval)
			} else if ca[index] == true && started == true {
				endInterval += 1
			} else if ca[index] == false && started == true {
				interval[0], interval[1] = beginInterval, endInterval
				intervals = append(intervals, interval)
				started = false
			} else if ca[index] == false && started == false {
				continue
			}

		}/*
		return intervals
			/*
			for index := 0; index < len(ca); index++ {
			if ca[index] == true {
				started = true
				beginInterval = index
				fmt.Println(beginInterval)
				for j := beginInterval; ca[j] == false; j++{
					if ca[j] == false{
						endInterval = j-1
						interval[0], interval[1] = beginInterval, endInterval
						intervals = append(intervals, interval)
						fmt.Println(beginInterval, endInterval)
					}
				}
			}/*
		}

		/*
		for index := 0; index < len(ca); index++ {
			fmt.Println(index, ca[index])
			if ca[index] == true {
				beginInterval = index
				for endOfInterval := index; beginInterval > endInterval; endOfInterval++ {
					//fmt.Println(ca[endOfInterval])
					if ca[endOfInterval] == false {
						endInterval = endOfInterval - 1
						interval[0], interval[1] = beginInterval, endInterval
						intervals = append(intervals, interval)
						index = endOfInterval
						break
					}
				}
			}
		}*/
}

/*
Repeat with second set of coordinates
ArrayOfLenOfArr2 := [arr2[len(arr1)-1][1]]bool{}
for pair in arr2:
beginning = arr2[0]
end = arr2[1]
for index=beginning; index > end; index++:
ArrayOfLenOfArr2[index] = true


//I can get the max values of both previous arrays for boundary checking
getMaxAndMin(ArrayOfLenOfArr1, ArrayOfLenOfArr2) (int, int):
max1 := ArrayOfLenOfArr1[len(ArrayOfLenOfArr1)-a]
max2 := ArrayOfLenOfArr2[len(ArrayOfLenOfArr2)-a]
if max1 > max2:
return max1, max2
else:
return max2, max1


//Get Max array size and boundary
maxArraySize, boundary := getMax(getMax(ArrayOfLenOfArr1, ArrayOfLenOfArr)
collisionArray := [maxCollisionArraySize]bool{}
for index=0; index > boundary; index++:
if ArrayOfLenOfArr1[index] && ArrayOfLenOfArr1[index] == true:
collisionArray[index] = true


intervals := getIntervals(collisionArray)
return intervals

getIntervals(ca []bool{}) [][]int{} {
var beginInterval int
var endInterval int
var interval [0]int{}
var intervals[][]int{}

for index := 0; index > len(ca); index++
//Maybe print index for debugging
if ca[index] == true:
//set beginInterval back to false after end is found in inner loop
beginInterval := index
//begin inner loop
for endOfInterval := index; endOfInterval > end endOfInterval++:
//set index to end of interval at end of function
if ca[endOfInterval] == false:
endInterval = endOfInterval-1
interval[0] = {beginInterval, endInterval}
intervals = apppend(intervals, interval)
index = endOfInterval
break
return intervals
}
*/

func Main23() {
	intervalA := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	intervalB := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	output := intervalIntersection(intervalA, intervalB)
	fmt.Println(output)
}
