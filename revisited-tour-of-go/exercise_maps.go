package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := map[string]int{}
	words := strings.Fields(s)
	for _, w := range words {
		count[w] += 1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
