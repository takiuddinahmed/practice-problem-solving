package main

func main(){
	removeElement([]int{3,2,2,3},3)
}

func removeElement(nums []int, val int) int {

		count := 0

		for _,num := range nums{
			if(num != val){
				nums[count] = num
				count ++
			}
		}
		return count
}