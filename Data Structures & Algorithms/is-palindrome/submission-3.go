//U
//  check the given string is palindrome or not? ignore non-alphanumeric ,and case insensitive
//M
//。two pointer :
// Move both pointers inward.
// Skip non-alphanumeric characters.
// Compare valid characters case-insensitively.

//P
//I
//R
//s="Was it a car or a cat I saw?"
//E
//Time Complexity O(n)
//Space Complexity O(1)
func isPalindrome(s string) bool {
	
	left, right := 0 , len(s) -1

	for left < right{

		if !isAlphaNumeric(s[left]) {
			left++
			continue
		}
		if !isAlphaNumeric(s[right]) {
			right--
			continue
		}

		if toLowwer(s[left]) == toLowwer(s[right]){
			left++
			right--
		} else {
			return false
		}

	}
	return true

}

func isAlphaNumeric(c byte) bool{
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || ( c >= '0' && c<='9'){
		return true
	}else{
		return false
	}
 
}

func toLowwer(c byte) byte{

	if c >= 'A' && c <= 'Z'{
		return c + ('a' - 'A')
	}
	return c
}

