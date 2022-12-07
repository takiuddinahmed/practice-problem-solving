package main

func main(){

}

func searchInsert(nums []int, target int) int {
	if(target < nums[0]) {
			return 0
	}
	length := len(nums)
	if (target > nums[length-1]) {
			return length
	}
	left := 0
	right := length -1
	middle := int((left+right)/2)
	for left <= right {
			if nums[middle] == target {
					return middle
			}
			if nums[middle] < target {
					left = middle+1
			} else {
					right = middle -1
			}
			
			middle = int((left+right)/2)
	
	}
	if target > nums[middle] {
			return middle +1
	}
	return middle
}