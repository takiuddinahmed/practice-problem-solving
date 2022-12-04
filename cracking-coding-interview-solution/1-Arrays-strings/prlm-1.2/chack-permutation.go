// Given two strings, write a method to decide if one is a permutaion of the other

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	fmt.Printf(" ab, ba -> %v \n",CheckPermutation("ab","ba"))
	fmt.Printf(" ab, bc -> %v \n",CheckPermutation("ab","bc"))
}

func CheckPermutation(str1 string, str2 string) bool {
	return strings.Compare(SortString(str1),SortString(str2)) == 0;
}

func SortString(str string) string{
	strArr := strings.Split(str, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}