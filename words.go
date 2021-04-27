package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	wordsMap := make(map[string]int)

	for i := 0; i < len(words); i++ {
		_, ok := wordsMap[words[i]]
		if !ok {
			wordsMap[words[i]] = 1
		} else {
			wordsMap[words[i]]++
		}
	}

	return wordsMap

}

func count() {
	wc.Test(WordCount)
}
