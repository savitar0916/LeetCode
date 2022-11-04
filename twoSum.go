package main



func twoSum(nums []int, target int) []int {
	// 依據進來的 nums 跑槽狀迴圈 來找出對比中哪兩個數字是我要的目標數
	// 我的暴力破解法
	// for i :=0;i < len(nums);i++{
    //     for j := 1 ;j < len(nums); j++{
    //         if target == nums[i]+nums[j] && i != j{
	// 			result := []int {nums[i],nums[j]}
	// 			return result
    //         }
    //     }
    // }
	// 時間複雜度O(N)的解法
    hashMap := map[int]int{}
    for i, x := range nums {
        if p, ok := hashMap[target-x]; ok {
            return []int{p, i}
        }
        hashMap[x] = i
    }
    return nil

}