package main

func isAlpha(v rune) bool {
	if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') {
		return false
	}
	return true
}

func AlphabetIndex(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a')
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 'A')
	} else {
		return -1
	}
}

func RepeatAlpha(s string) string {
	result := []rune{}
	for _, v := range s {
		if isAlpha(v) {
			for i := 0; i <= AlphabetIndex(v); i++ {
				result = append(result, v)
			}
		}
	}

	return string(result)
}
