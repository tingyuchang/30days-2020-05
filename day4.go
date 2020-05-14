// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3319/
/*
Given a positive integer num, output its complement number. The complement strategy is to flip the bits of its binary representation.



Example 1:

Input: num = 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
Example 2:

Input: num = 1
Output: 0
Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findComplement(987))
	fmt.Println(findComplement2(987))
}

func findComplement(num int) int {
	x := 0.0
	for int(math.Pow(2, x)) <= num {
		x++
	}

	return int(math.Pow(2, x)) - 1 - num
}

func findComplement2(num int) int {
	tmp := num
	mask := 0
	for tmp > 0 {
		tmp = tmp >> 1
		mask = mask << 1
		mask = mask | 1
	}
	fmt.Println(num ^ mask)

	return 9999
}
