package main

func main() {
	
}

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool)

	for _, val := range nums {
		if _, exists := numsMap[val]; exists {
			return true
		}
		numsMap[val] = true
	}
	return false
}