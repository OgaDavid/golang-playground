package main

// A ROT14 string is a string encoded using the ROT14 cipher,
// which is simply a Caesar cipher that shifts each alphabet
// letter by 14 positions.

func isUpper(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func Rot14(str string) string {

	result := make([]rune, len(str))

	for i, v := range str {
		if isLower(v) {
			newChar := 'a' + (v-'a'+14)%26
			result[i] = newChar
		} else if isUpper(v) {
			newChar := 'A' + (v-'A'+14)%26
			result[i] = newChar
		} else {
			result[i] = v
		}
	}

	return string(result)
}
