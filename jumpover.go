package main

// Write a function JumpOver() that takes a string and 

// returns another string with every third character.

// Prints the output followed by newline \n.

// If the string is empty, return newline \n.

// If there is no third character, return newline \n.

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	var result []rune
	count := 0

	for _, r := range str {
		count++
		if count%3 == 0 {
			result = append(result, r)
		}
	}

	if len(result) == 0 {
		return "\n"
	}

	return string(result) + "\n"
}
