package main

func CheckBrackets(input string) bool {
	var stack []rune

	pairs := map[rune]rune{
		'>': '<',
		'}': '{',
		']': '[',
	}

	for _, char := range input {
		switch char {
		case '<', '{', '[':
			stack = append(stack, char)
		case '>', '}', ']':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			expected, exists := pairs[last]
			if !exists || expected != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
