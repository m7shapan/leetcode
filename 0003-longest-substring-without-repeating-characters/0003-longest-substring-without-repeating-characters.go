func lengthOfLongestSubstring(s string) int {
    substring := make(map[byte]struct{})
    
    l :=0
    max := 0
    for i:=0;i<len(s); {
        _, found := substring[s[i]]
        if !found {
            substring[s[i]] = struct{}{}
            i++
        } else {
            delete(substring, s[l])
            l++
        }
        
        if len(substring) > max {
            max = len(substring)
        }
    }
    
    return max
}

// "aab"
// a
// 