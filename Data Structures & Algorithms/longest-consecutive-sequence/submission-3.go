//Time Complexity O(nlogn)
//Space Complexity O(n)
//Sorting 
func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	maxLength := 1

	for i,curLen:= 1,1 ; i< len(nums); i++{
		
		if nums[i] == nums[i-1]{
			continue
		}
		if nums[i] == nums[i-1] + 1{
			curLen++
		}else {
			curLen = 1
		}
	    
		if curLen > maxLength {
			maxLength = curLen
		}
		
	}
	return maxLength
	
}
