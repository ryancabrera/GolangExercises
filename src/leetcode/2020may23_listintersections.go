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
*/

package leetcode

import (
	"fmt"
)

func intervalIntersection(A [][]int, B [][]int) [][]int {
	for index, element := range A {
		fmt.Println("Index: ", index, "\tElement: ", element)
		fmt.Println(element[0])
	}
	return A
}

func Main23() {
	intervalA := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	intervalB := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	output := intervalIntersection(intervalA, intervalB)
	fmt.Println(output)
}
