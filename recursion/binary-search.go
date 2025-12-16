package main

func BinarySearch(arr []int, n int) int {
	arrlen := len(arr)
	mid := arrlen / 2

	var tempArr []int

	if arrlen == 0 {
		return -1
	}

	if arr[mid] != n {
		if arr[mid] < n {
			tempArr = arr[mid+1:]
			result := BinarySearch(tempArr, n)
			if result == -1 {
				return -1
			}
			return (mid + 1) + BinarySearch(tempArr, n)
		} else {
			tempArr = arr[:mid]
			return BinarySearch(tempArr, n)
		}
	}

	return mid
}
