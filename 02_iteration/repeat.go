package iteration

import "strings"

// Repeat takes charater and time to count, return the character concatenate by time to count
func Repeat(character string, repeatCount int) string {
	repeated := strings.Repeat(character, repeatCount)
	return repeated
}
