//Optimized two pointer (not use space complexity O(n))
//Time Complexity O(n)
//Space Complexity O(1)
func isPalindrome(s string) bool {

	left, right:= 0, len(s) - 1

	for left < right{

		if !isAlphanumeric(s[left]){
			left++
			continue
		}

		if !isAlphanumeric(s[right]){
			right--
			continue
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		
		left++
		right--

	}

	return true
	
}


func toLower(c byte) byte{
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}


func isAlphanumeric(c byte) bool{

	return (c >= 'a' && c <= 'z') || 
	(c >= 'A' && c <= 'Z') || 
	(c >= '0' && c <= '9')
}