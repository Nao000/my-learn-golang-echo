package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	list := strings.Fields(s)

	for i := 0; i < len(list); i++ {
		_, ok := m[list[i]]

		if ok {
			m[list[i]]++
		} else {
			m[list[i]] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}