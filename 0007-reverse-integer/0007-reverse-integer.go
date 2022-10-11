func reverse(x int) int {
    var isNegative bool
    if x < 0 {
        isNegative = true
        x *= -1
    }
    
    var answer int64
    for x > 9 {
        answer += int64(x%10)
        answer *= 10
        x = (x - (x%10)) / 10
    }

    answer += int64(x)
    
    if math.MaxInt32 < answer {
        return 0
    }
    
    if isNegative == true {
        answer *= -1
    }
    
    return int(answer)
}