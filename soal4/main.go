package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var listWords = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	var groupWords = make(map[string][]string)

	for _, word := range listWords {
		splWord := strings.Split(word, "")
		sort.Strings(splWord)
		sortedWord := strings.Join(splWord, "")
		groupWords[sortedWord] = append(groupWords[sortedWord], word)
	}

	for _, words := range groupWords {
		fmt.Println(words)
	}
}
