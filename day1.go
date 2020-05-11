// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3316/
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

package main

import "fmt"

func main() {
	fmt.Println("output: ", firstBadVersion(10000000))
}

func firstBadVersionBS(n int) int {
	goal := 1
	for i := 1; i <= n; i++ {
		if isBadVersion(i) == true && isBadVersion(i-1) == false {
			goal = i
		}
	}

	return goal
}

func firstBadVersionBSV2(n int) int {
	goal := 0
	divideBy2 := n / 2

	for i := 1; i <= divideBy2; i++ {
		if isBadVersion(i * 2) {
			if isBadVersion(i*2-1) == false {
				goal = i * 2
				break
			} else {
				goal = i*2 - 1
				break
			}
		}
	}

	if goal == 0 {
		goal = n
	}

	return goal
}

func firstBadVersionV3(n int) int {
	left := 1
	right := n

	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(n int) bool {
	if n < 8990 {
		return false
	} else {
		return true
	}
}
