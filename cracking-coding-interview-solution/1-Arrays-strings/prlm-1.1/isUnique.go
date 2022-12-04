/*
Problem statement: Implement an algorithm to determine if a string has all unique characters.
*/
package main

import "fmt"

func main(){
	fmt.Printf("abc -> %v \n",IsUniqueChars("abc") )
	fmt.Printf("abca -> %v \n",IsUniqueChars("abca") )
}

func IsUniqueChars(value string) bool{
	valMap := make(map[rune]bool)
	for _, val := range value{
		if _, exist := valMap[val]; exist {
			return false
		}
		valMap[val] = true
	}
	return true
}