func moveZeroes(nums []int)  {
    ai := 0
    zeroesNum := 0
    for i:=0;i<len(nums);i++{
        if nums[i] == 0 {
            zeroesNum++
            continue
        }
        
        nums[ai] = nums[i]
        ai++
    }
    
    for i:=len(nums)-zeroesNum;i<len(nums);i++{
        nums[i]=0
    }
}