package main

func Itoa(n int) string {
	// Special case: zero
	if n == 0 {
		return "0"
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n // make positive for processing
	}

	var runes []rune

	// Extract digits from right to left
	for n > 0 {
		digit := n % 10
		runes = append(runes, rune(digit)+'0')
		n /= 10
	}

	// Add '-' sign if needed
	if negative {
		runes = append(runes, '-')
	}

	// runes are reversed → reverse them to correct order
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
