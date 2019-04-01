/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.22%)
 * Total Accepted:    643.7K
 * Total Submissions: 2.6M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */
func maxBound() int {
	maxi := uint(1)<<31 - 1
	return int(maxi)
}

func reverse(x int) int {
	bound := maxBound()
	maxBoundD := bound / 10
	maxBoundM := bound % 10
	minBoundD := -maxBoundD
	minBoundM := -maxBoundM - 1

	factor := 1
	if x < 0 {
		factor = -1
	}

	s := []int{x % 10}

	for x = x / 10; x != 0; x = x / 10 {
		s = append(s, x%10)
	}

	r := 0
	for _, v := range s {
		// skip leading 0
		if v == 0 && r == 0 {
			continue
		}
		if factor == 1 {
			// over bound
			if r > maxBoundD {
				return 0
			}
			// over bound
			if r == maxBoundD && v > maxBoundM {
				return 0
			}
		}

		if factor == -1 {
			if r < minBoundD {
				return 0
			}

			if r == minBoundD && v < minBoundM {
				return 0
			}
		}

		r = r*10 + v
	}

	return r
}


