func rotate(nums []int, k int)  {
    if len(nums) < 2 {
        return
    }
    
    if k > len(nums) {
        k = k%len(nums)
    }
    
    result := append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
    
    for i:=0;i<len(nums);i++{
        nums[i] = result[i]
    }
}
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// nums[len(nums)-3:]
// nums[3:len(nums)-3]