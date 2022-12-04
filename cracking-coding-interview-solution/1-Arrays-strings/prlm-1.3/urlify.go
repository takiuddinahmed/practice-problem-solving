/*
	URLify: Write a method to replace all spaces in a string with ' .

You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the "true" length of the string.
(Note: If implementing in Java, please use a character array so that you can perform this operation in place.)
*/
package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(URLify("Mr John Smith       "))
}

func URLify(value string)string{
	trimedValue := strings.TrimSpace(value);
	return strings.ReplaceAll(trimedValue," ","%20")
}