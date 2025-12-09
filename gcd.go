package main

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	// Euclid's Algorithm
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
