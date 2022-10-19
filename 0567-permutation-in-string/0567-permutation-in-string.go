func checkInclusion(s1 string, s2 string) bool {                                                   
    var isPermutation bool
    permutations := make(map[byte]int)
    
    for i:=0;i < len(s1); i++ {
        v, _ := permutations[s1[i]]
        permutations[s1[i]] = v+1
    }
    
    j:=0
    for i:=0;i < len(s2);{
        _, found := permutations[s2[i]]
        if !found && !isPermutation {
            i++
        } else if found && !isPermutation {
            isPermutation = true
            
            if v, found := permutations[s2[i]]; found && v > 1 {
                permutations[s2[i]] = v-1
            } else {
                delete(permutations, s2[i])
            }
            j = i
            i++
        } else if !found && isPermutation && i != j {
            v, _ := permutations[s2[j]]
            permutations[s2[j]] = v+1
            j++
        } else if found && isPermutation {
            if v, found := permutations[s2[i]]; found && v > 1 {
                permutations[s2[i]] = v-1
            } else {
                delete(permutations, s2[i])
            }
            i++
        } else {
            isPermutation = false
            i++
            j++
        }
        fmt.Println(permutations)
        
        if len(permutations) == 0 {
           return true
        }
    }
    
    return false
}

// ooolleoooleh
// F F
// F F
// F F
// T 