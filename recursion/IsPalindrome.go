package recursion

// Write a recursive function called IsPalindrome which returns true
// if the string passed to it is a palindrome (reads the same forward and backward).
// Otherwise it returns false.
func IsPalindrome(s string) bool {

	// If length is zero or not even, return false
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	if string(s[0]) == string(s[len(s)-1]) {
		return true
	} else {
		return IsPalindrome(string(s[1 : len(s)-1]))
	}

}
