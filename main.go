package main

import "fmt"

func main(){
	nums :=  []int {2,7,11,15}
	target := 9
	nums = twoSum(nums,target)
	fmt.Println (nums)
}

