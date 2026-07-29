//brute force
//Time Complexity O(n^2)
//space Complexity O(n)
func longestConsecutive(nums []int) int {
	
	maxLength := 0
	hashSet := make(map[int]struct{})
	for _, num := range nums {
		hashSet[num] = struct{}{} 
	}
	
	//loop hashSet 

	//1 check if num in nums is exist in hashSet 
	//2 if exist then num++ currLength++ and keep check num++ in next loop
	//3 if not exist , go on next num in nums array

	for _, num:= range nums {
		length,currentNum := 0 , num
		
		
		//method 3 
		for {
			if _, ok := hashSet[currentNum]; !ok{
				break
			}
			currentNum++
			length++
		}


		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
