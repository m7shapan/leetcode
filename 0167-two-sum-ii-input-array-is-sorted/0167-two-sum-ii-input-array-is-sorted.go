func twoSum(numbers []int, target int) []int {
    result := make([]int, 2)
    for i:=0;i<len(numbers);i++ {
        x := target - numbers[i]
        found := false
        j:=i
        for ; j+1 < len(numbers); {
            j++
            if numbers[j] > x {
                break
            }
            
            if numbers[j] == x {
                found = true
                break
            }
        }
        
        if found {
            result[0] = i+1
            result[1] = j+1
        }
    }
    
    return result
}