package main

// import "fmt"

// Write a function Compact that takes a pointer to a
// slice of strings as the argument. This function must:
// Return the number of elements with non-zero value.
// Compact, i.e., delete the elements with zero-values in the slice.

func Compact(ptr *[]string) int {
	count := 0
	var compacted []string
	for _, v := range *ptr {
		if v != "" {
			count++
			compacted = append(compacted, v)
		}
	}

	*ptr = compacted

	return count
}
