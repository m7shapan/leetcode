func searchInsert(nums []int, target int) int {
    l := 0 
    h := len(nums)-1
    
    for (h-l) > 1 {
        mid := (h+l)/2
        
        if target > nums[mid] {
            l = mid+1
        } else {
            h = mid
        }
    }
    
    if target <= nums[l] {
        return l
    } else if l != h && nums[h] >= target{
        return h
    } else {
        return h+1
    }
    
    return h
}

// l=0 h=2 m=1
// 3 > 3 
// l=2 h=2
// i1 
