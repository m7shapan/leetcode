func sortedSquares(nums []int) []int {
    
    for i:=0;i<len(nums); i++ {
        nums[i] = nums[i]*nums[i]
    }
    
    for i:=0;i<len(nums); i++ {
        for j:=i+1;j<len(nums); j++ {
            if nums[i] > nums[j] {
                tmp := nums[j]
                nums[j] = nums[i]
                nums[i] = tmp
            }
        }
    }
    
    return nums
}
