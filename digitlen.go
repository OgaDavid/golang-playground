package main

func DigitLen(n, base int) int {
	if n < 0 {
		n = -n
	}

	if base < 2 || base > 36 {
		return -1
	}

	count := 0

	for n > 0 {
		n /= base
		count++
	}

	return count
}
