package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

type sortRunes []rune
type charMap struct {
	Key   string
	Value int
}
type charMapList []charMap

func Main() {
	exampleInput1 := StripSpace("Aabb")

	m := strMap(exampleInput1)
	for key, element := range m {
		fmt.Println(key, element)
	}
	fmt.Println(m)

	keys, values := SliceKeysAndValues(m)

	s := []rune(exampleInput1)
	sort.Sort(sortRunes(s))
	fmt.Println(string(s))

	fmt.Println(keys, values)

	srtd := rankByWordCount(m)
	fmt.Println(srtd)

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

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SliceKeysAndValues(stringmap map[string]int) ([]string, []int) {
	keys := make([]string, 0, len(stringmap))
	values := make([]int, 0, len(stringmap))

	for k, v := range stringmap {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func StripSpace(inputString string) string {
	return strings.Join(strings.Fields(inputString), "")
}

func strMap(inputString string) map[string]int {
	runifiedInput := []rune(inputString)
	m := make(map[string]int)
	fmt.Println(m)
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
