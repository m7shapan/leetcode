func sortedSquares(nums []int) []int {
    result := make([]int, len(nums))
    l := 0
    r := len(nums)-1
    for i:=len(result)-1;i>=0;i-- {
        if nums[l]*nums[l] > nums[r]*nums[r] {
            result[i] = nums[l]*nums[l]
            l++
        } else {
            result[i] = nums[r]*nums[r]
            r--
        }
    }
    
    return result
}
