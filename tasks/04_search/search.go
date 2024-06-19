package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

var input = []string{"пятка", "пятак", "листок", "тяпка", "слиток", "столик", "ааббвв", "ббаавв", "ввббаа"}

func main() {
	s := searchAnagram(&input)
	for k, v := range *s {
		fmt.Println(k, *v)
	}
}

func searchAnagram(words *[]string) *map[string]*[]string {
	mapOfAnagrams := make(map[string]string)
	result := make(map[string]*[]string)
	for _, v := range *words {
		v := strings.ToLower(v)
		sortedString := getSortedString(v)
		firstWord, ok := mapOfAnagrams[sortedString]
		if ok {
			*result[firstWord] = append(*result[firstWord], v)
		} else {
			mapOfAnagrams[sortedString] = v
			result[v] = &[]string{v}
		}
	}

	for k, v := range result {
		if len(*v) < 2 {
			delete(result, k)
			continue
		}
		cmp := func(a, b string) int {
			return strings.Compare(a, b) * -1
		}
		slices.SortStableFunc(*v, cmp)
	}
	mapOfAnagrams = nil
	return &result
}

func getSortedString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
