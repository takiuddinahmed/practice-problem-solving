/* Palindrome Permutation: Given a string,
write a function to check if it is a permutation of a palin drome.
 A palindrome is a word or phrase that is the same forwards and backwards.
 A permutation is a rearrangement of letters.
 The palindrome does not need to be limited to just dictionary words.
*/

package main

import (
	"fmt"
	"strings"
)


func main(){
	fmt.Printf("'Tact Coa' -> %v \n",palindromPermutation("Tact Coa"))
	fmt.Printf("'Tact Coaa' -> %v \n",palindromPermutation("Tact Coaa"))
}

func palindromPermutation(value string) bool{
	// lower case and remove space
	value = strings.ReplaceAll(strings.ToLower(value)," ","")

	// count all chars
	charMap := make(map[rune]int)
	for _,v := range value {
		if _,ok := charMap[v] ; ok {
			charMap[v] ++
		} else {
			charMap[v] = 1
		}
	}

	// check odd counts
	odd := 0

	for _,v := range charMap{
		if v % 2 != 0 {
			odd ++
		} 
	}

	// for even length string, all char count should be  even
	if len(value) % 2 == 0 {
		return odd == 0
	}
	// for odd length string, just 1 char should be odd for middle position
	return odd == 1
}


