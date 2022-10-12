/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
 
    l := 0
    h := n
    
    for (h-l) > 1 {
        m := (h+l)/2
        if !isBadVersion(m) {
            l = m+1
        } else {
            h = m
        }
    }
    
    var bad int
    if isBadVersion(l) {
        bad = l
    } else {
        bad = h
    }
    
    if bad > 1 && isBadVersion(bad-1) {
        return bad-1
    }
    
    return bad
}