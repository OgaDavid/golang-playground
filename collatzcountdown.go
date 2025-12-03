package main

// Write a function, CollatzCountdown, that returns the
// number of steps necessary to reach 1 using the collatz countdown.

// Thinking process 💡
// start with a number
// if the number is even - divide by 2
// if the number is odd - multiply by 3 & add 1
// return the number of steps to get to 1
// we'd most likely use recursion or while loop

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	steps := 0

	for start != 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = 3*start + 1
		}
		steps++
	}

	return steps
}
