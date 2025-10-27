package namelist

import "sort"

func increaseMap(m map[string]int, s []string) map[string]int {
	for _, v := range s {
		_, exist := m[v]
		if !exist {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	return m
}

func merge(s1, s2 []string) []string {
	unifiqueSlice := countRepeated(s1, s2)
	resultSlice := newStringSlice(unifiqueSlice)
	sort.Strings(resultSlice)
	return resultSlice
}

func newStringSlice(m map[string]int) []string {
	var result []string

	for key := range m {
		result = append(result, key)
	}
	return result
}

func countRepeated(s1, s2 []string) map[string]int {
	result := make(map[string]int)
	result = increaseMap(result, s1)
	result = increaseMap(result, s2)

	return result
}
