package main

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0

	// Check sign
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	n := 0

	for i := start; i < len(s); i++ {
		ch := s[i]

		// If not a digit → invalid
		if ch < '0' || ch > '9' {
			return 0
		}

		// Convert ASCII to number:
		// '0' → 0, '1' → 1, ..., '9' → 9
		digit := int(ch - '0')

		n = n*10 + digit
	}

	return n * sign
}
