package recursion

// Write a recursive function called reverse
// which accepts a string and returns a new string in reverse.
func Reverse(s string) string {
	if len(s) <= 1 {
		return s
	}
	return string(s[len(s)-1]) + Reverse(string(s[:len(s)-1]))
}
