package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("aba"))
	// fmt.Println(isPalindrome("aa"))
}

func isPalindrome(s string) bool {
	str := ""
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') || (v >= 'A' && v<= 'Z') {
			str += strings.ToLower(string(v))
		}
	}
	lenth := len(str)

	for i:=0;i<int(math.Floor(float64(lenth)/2.0));i++{
		if(str[i] != str[lenth-1-i]){
			return false
		}
	}
	return true
}
