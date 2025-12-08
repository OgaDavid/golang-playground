package main

// Write a function that returns the median of five int arguments.

func sortNumbers(a, b, c, d, e int) []int {
	arr := []int{a, b, c, d, e}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	return arr
}

func Abort(a, b, c, d, e int) int {
	numbers := sortNumbers(a, b, c, d, e)
	return numbers[2]
}
