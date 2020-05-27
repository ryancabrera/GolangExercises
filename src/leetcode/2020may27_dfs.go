package leetcode

import (
	"fmt"
)

func possibleBipartition(N int, dislikes [][]int) bool {
	colorMap := make(map[int]int)
	edges := make(map[int][]int)

	for _, dislike := range dislikes {
		edges[dislike[0]] = append(edges[dislike[0]], dislike[1])
		edges[dislike[1]] = append(edges[dislike[1]], dislike[0])
	}

	for node := 1; node <= N; node++ {
		if _, ok := colorMap[node]; !ok {
			if !dfs(node, 0, &colorMap, &edges) {
				return false
			}
		}
	}
	return true
}

func dfs(node int, color int, colorMap *map[int]int, edges *map[int][]int) bool {
	if val, ok := (*colorMap)[node]; ok {
		return val == color
	} else {
		(*colorMap)[node] = color
	}

	for _, neigh := range (*edges)[node] {
		if !dfs(neigh, color^1, colorMap, edges) {
			return false
		}
	}
	return true
}

func PossiblePartition() {
	input1 := [][]int{{1, 2}, {1, 3}, {2, 4}}
	output1 := possibleBipartition(4, input1)
	fmt.Println(output1)

	input2 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	output2 := possibleBipartition(3, input2)
	fmt.Println(output2)

	input3 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	output3 := possibleBipartition(5, input3)
	fmt.Println(output3)
}
