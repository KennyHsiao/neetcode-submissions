//HashSet
//Time complexity O(n)
//Space complexity O(n)
func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]struct{},len(nums))
	//store in numSet 
	for _, num := range nums{
		numSet[num] = struct{}{}
	}

	maxLength := 1


	for num := range numSet{
		//start from the begin of this consecutive sequence
		if _, found := numSet[num - 1]; found{
			continue
		}

		currentNum := num
		currentLength := 1

		for {
			if _,exist := numSet[currentNum + 1] ;!exist{
				break
			}
			
			currentNum++
			currentLength++
		

		}

		if currentLength > maxLength{
				maxLength = currentLength
			}
	}
	return maxLength

}
