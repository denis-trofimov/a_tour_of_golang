package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	substrings := strings.Fields(s)
	for _, word := range substrings {
		count[word] += 1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
