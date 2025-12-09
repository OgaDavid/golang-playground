package main

// Write a function that takes a string and return a string
// containing its first word, followed by a newline ('\n').
// A word is a sequence of characters delimited by spaces
// or by the start/end of the argument.

func FirstWord(s string) string {
	result := []rune{}
	started := false

	for _, v := range s {
		if v == ' ' || v == '\t' {
			if started {
				break
			}
			continue
		}
		started = true
		result = append(result, v)
	}
	return string(result) + "\n"
}
