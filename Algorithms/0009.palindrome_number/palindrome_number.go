package _009_palindrome_number

/* Approach 1: Using simple push and pop(my first answer) */
/* Time complexity : O(log10(n)) Space complexity: O(1) */
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	var rev int
	dividedx := x

	for dividedx != 0 {
		//pop
		pop := dividedx % 10
		dividedx /= 10
		//push
		rev = rev * 10 + pop
	}

	return x == rev
}

/* Approach 2: Revert half of the number */
/* Time complexity : O(log10(n)) Space complexity: O(1) */
func isPalindrome2(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	var rev int
	for x > rev {
		//pop and push
		rev = rev * 10 + x % 10
		x /= 10
	}
	// Considering about both odd and even cases.
	return x == rev || x  == rev / 10
}
