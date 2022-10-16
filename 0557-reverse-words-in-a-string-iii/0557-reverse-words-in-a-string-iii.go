func reverseWords(s string) string {
    ss := []byte(s)
    
    firstIndexAfterSpace := 0
    for i:=0; i<= len(ss); i++ {
        if len(s) == i || s[i] == ' ' {
            reverseWord(ss, firstIndexAfterSpace, i-1)
            firstIndexAfterSpace = i+1
        }
    }
    
    return string(ss)
}

func reverseWord(s []byte, l,r int) {
    var tmp byte
    for l<r {
        tmp = s[l]
        s[l] = s[r]
        s[r] = tmp
        
        l++
        r--
    }
    
} 