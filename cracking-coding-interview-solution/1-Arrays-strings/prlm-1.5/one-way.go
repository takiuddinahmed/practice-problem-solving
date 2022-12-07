/*


 */

package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Printf("pale, vble -> %v \n",checkOneWay("pale","vble"))
	fmt.Printf("pale, pble -> %v \n",checkOneWay("pale","pble"))
	fmt.Printf("pale, vvle -> %v \n",checkOneWay("pale","vvle"))
	fmt.Printf("pale, ale -> %v \n",checkOneWay("pale","ale"))
	fmt.Printf("pale, vle -> %v \n",checkOneWay("pale","vle"))
	fmt.Printf("pale, vpa -> %v \n",checkOneWay("pale","pa"))
	fmt.Printf("pale, pavle -> %v \n",checkOneWay("pale","pavle"))
	fmt.Printf("pale, padele -> %v \n",checkOneWay("pale","padele"))
	fmt.Printf("pale, padee -> %v \n",checkOneWay("pale","padee"))

}

func checkOneWay (originalStr string, editedStr string) bool {
	// zero edit
	if strings.Compare(originalStr, editedStr) == 0 {
		fmt.Println("Matched")
		return true
	}

	originalStrLenth := len(originalStr)
	editedStrLenth := len(editedStr)

	// insert
	if originalStrLenth + 1 == editedStrLenth {
		return checkInsertOrRemove(editedStr,originalStr)
	}

	// remove 
	if (originalStrLenth == 1+ editedStrLenth) {
		return checkInsertOrRemove(originalStr,editedStr)
	}

	if (originalStrLenth == editedStrLenth) {
		return checkReplace(originalStr, editedStr, originalStrLenth)
	}
	return false
	
}

// pale,  pble
func checkReplace(originalStr string, editedStr string, length int) bool {
	notMatched := false
	for i:=0;i<length;i++{
		if originalStr[i] != editedStr[i]{
			if(notMatched) {
				return false
			}
			notMatched = true
		}
	}
	return true
}

// pale, pae
func checkInsertOrRemove(largeStr string, smallerStr string) bool {
	extraCharCount := 0
	lenth := len(smallerStr)

	i:=0

	for ;i<lenth;{
		if largeStr[i+extraCharCount] != smallerStr[i] {
			if extraCharCount > 0 {
				return false
			}
			extraCharCount = 1
		}else 
		{i++}
	}
	return true
}