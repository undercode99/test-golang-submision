package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupingAnagrams(strs []string) [][]string {
	mapWord := make(map[string][]string)

	for _, word := range strs {
		key := sortStringAscending(word)
		mapWord[key] = append(mapWord[key], word)
	}

	result := make([][]string, 0)
	for _, v := range mapWord {
		result = append(result, v)
	}

	return result
}

func sortStringAscending(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	fmt.Println(groupingAnagrams([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}))
}
