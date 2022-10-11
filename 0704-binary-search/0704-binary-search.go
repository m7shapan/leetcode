func search(nums []int, target int) int {
    low := 0
    high := len(nums)-1
    
    for (high-low) > 1 {
        mid := (low+high)/2
        if nums[mid] < target {
            low = mid+1
        } else {
            high = mid
        }
    }
    
    if nums[high] == target {
        return high
    } else if nums[low] == target {
        return low
    }
    
    return -1
}
