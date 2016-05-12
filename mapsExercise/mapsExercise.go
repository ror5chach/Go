package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	tokens := []string(strings.Fields(s))
	
	myMap := make(map[string]int)
	
	for _,value := range tokens {
		_, ok := myMap[value]
		if (ok) {
			myMap[value]++
		} else {
			myMap[value] = 1	
		}
	}
	
	return myMap
}

func main() {
	wc.Test(WordCount)
}
