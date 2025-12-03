package main

import (
	"fmt"
	"strings"
)

func makeItemsSlice(str string) []string {
	var result []string
	word := ""

	for _, v := range str {
		if v == ' ' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(v)
		}
	}

	// Append last word if not empty
	if word != "" {
		result = append(result, word)
	}

	return result
}

func ShoppingSummaryCounter(str string) map[string]int {

	itemsSlice := makeItemsSlice(str)
	// option 2-------------
	itemsSlice2 := strings.Fields(str)
	fmt.Println(itemsSlice)
	//-----------------
	fmt.Println(itemsSlice2)
	items := make(map[string]int)

	for _, item := range itemsSlice {
		items[item]++
	}

	return items
}
