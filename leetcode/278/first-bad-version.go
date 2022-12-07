package main

import "fmt"

func main(){
	fmt.Println(firstBadVersion(3))
}

func isBadVersion(num int) bool{
	return num >= 1 
}

func firstBadVersion(n int) int {
    
	left := 1
	right := n
	middle := int((left+right)/2)

	for left <= right {
			if(isBadVersion(middle)){
					right = middle -1
				 
			} else {
					left = middle + 1
			}
			middle = int((left+right)/2)
			
	}
	return left
}