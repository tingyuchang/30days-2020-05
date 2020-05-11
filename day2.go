// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3317/
/*
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  
Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. 
Letters are case sensitive, so "a" is considered a different type of stone from "A".
S and J will consist of letters and have length at most 50.
The characters in J are distinct.
*/
package main 

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numJewelsInStones("asSWH", "ahshrvwqcviwbtrbiyiqocynqwocqcdnahdhhdwqhq"))
}

func numJewelsInStones(J string, S string) int {
	count := 0
	stones := strings.Split(S, "")
	// jewels := strings.Split(J, "")
    for _, v  :=  range stones  {
		if strings.Contains(J, v) {
			count++
		}
	}
	return count
}

func  numJewelsInStonesV2(J string, S string) int {
	count := 0
	jewels := make(map[int32]bool)

	for _, v := range J {
		jewels[v] = true
	}  

	for _, v := range S {
		if jewels[v] {
			count++
		}
	}
	return count
}
