/*
1.6
String Compression: Implement a method to perform basic string compression using the counts of repeated characters.
For example, the string aabcccccaaa would become a2b1c5a3.
If the "compressed" string would not become smaller than the original string, your method should return the original string.
You can assume the string has only uppercase and lowercase letters (a-z).
*/

package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Printf("aabccccccaaa -> %v \n",stringCompression("aabcccccaaa"))
	fmt.Printf("aabcaaa -> %v \n",stringCompression("aabcaaa"))
}

func stringCompression(str string) string {
	previousChar := str[0]
	count := 1
	finalStr := string(previousChar)

	for i:=1;i<len(str);i++{
		if str[i] != previousChar {
			finalStr += strconv.Itoa(count) + string(str[i])
			count = 1
			previousChar = str[i]
		} else {
			count ++
		}
	}
	finalStr += strconv.Itoa(count)

	if len(finalStr) < len(str) {
		return finalStr
	}
	return str
}