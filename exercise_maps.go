package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	substrings := strings.Fields(s)
	for _, word := range substrings {
		_, ok := count[word]
		if ok {
			count[word]++
		} else {
			count[word] = 1
		}
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
