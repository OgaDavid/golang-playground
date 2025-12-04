package main

func UnMatch(elems []int) int {
	for _, el := range elems {
		count := 0
		for _, v := range elems {
			if v == el {
				count++
			}
		}
		if count%2 != 0 {
			return el
		}
	}
	return -1
}
