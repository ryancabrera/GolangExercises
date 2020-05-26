//Credit from user Leaf Peng from leetcode

package leetcode

import "fmt"

func intersections(A [][]int, B [][]int) [][]int {
	a, b := 0, 0
	ans := [][]int{}
	for a < len(A) && b < len(B) {
		a1, a2, b1, b2 := A[a][0], A[a][1], B[b][0], B[b][1]
		if a1 <= b2 && b1 <= a2 {
			ans = append(ans, []int{ma(a1, b1), mi(a2, b2)})
		}
		if a2 > b2 {
			b++
		} else {
			a++
		}
	}
	return ans
}

func ma(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mi(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Main25() {
	intervalA := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	intervalB := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}

	output := intersections(intervalA, intervalB)
	fmt.Println(output)

}
