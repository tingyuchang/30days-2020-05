// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3318/
/*
Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aaa", "aaabb"))
}

func canConstruct(ransomNote string, magazine string) bool {
    notes := make(map[int32]int)
    count := len(ransomNote)
    for _,v := range magazine {
        notes[v] += 1
    } 
    
    for  _,v := range ransomNote {
        if notes[v] > 0 {
            count--
            notes[v] -= 1
        }
    }
    
    return (count==0)
} 