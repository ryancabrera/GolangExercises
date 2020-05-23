package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

type charMap struct {
	Key   string
	Value int
}
type charMapList []charMap

func Main() {
	//exampleInput1 := StripSpace("Aabb")
	exampleInput1 := "AAaaabbb  b a  baba asd"
	m := strMap(exampleInput1)
	srtd := rankByWordCount(m)

	uncatOutput := []string{}

	for _, element := range srtd {
		for i := 0; i < element.Value; i++ {
			uncatOutput = append(uncatOutput, element.Key)
		}
	}
	fmt.Println(strings.Join(uncatOutput, ""))
}

func (p charMapList) Len() int           { return len(p) }
func (p charMapList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p charMapList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func rankByWordCount(wordFrequencies map[string]int) charMapList {
	pl := make(charMapList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = charMap{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func StripSpace(inputString string) string {
	return strings.Join(strings.Fields(inputString), "")
}

func strMap(inputString string) map[string]int {
	runifiedInput := []rune(inputString)
	m := make(map[string]int)
	for i := 0; i < len(runifiedInput); i++ {
		_, key := m[string(runifiedInput[i])]
		if key {
			m[string(runifiedInput[i])] += 1
		} else {
			m[string(runifiedInput[i])] = 1
		}
	}
	return m
}
