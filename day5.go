// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3320/
/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}

func firstUniqChar(s string) int {
	tmp := make(map[int32]int)

	for _, v := range s {
		c := tmp[v]
		tmp[v] = c + 1
	}

	for i, v := range s {
		if tmp[v] == 1 {
			return i
		}
	}

	return -1
}
