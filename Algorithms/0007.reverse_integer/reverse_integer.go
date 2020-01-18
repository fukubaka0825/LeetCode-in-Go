package _07_reverse_integer

import "math"

/* [Go] Approach 1: Pop and Push Digits & Check before Overflow */
//MaxInt32 == 32767
func reverse(x int) int {
	var rev int
	for x != 0 {
		//1.pop operation
		pop := x % 10
		x /= 10

		//2.check beforehand whether or not appending push ope would cause overflow.
		if ( rev > math.MaxInt32 / 10 || (rev == math.MaxInt32/10 && pop > 7)){
			return 0
		}
		if ( rev < math.MinInt32 / 10 || (rev == math.MinInt32/10 && pop < -8)){
			return 0
		}

		//3. push operation
		rev = rev * 10 + pop
	}
	return rev
}