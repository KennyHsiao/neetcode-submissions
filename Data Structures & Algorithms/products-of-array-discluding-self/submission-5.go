//Time Complexity O(n)
//Space Complexity O(n)
func productExceptSelf(nums []int) []int {
	
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	prefix[0],suffix[len(nums)-1]= 1,1

	for i,j := 1,len(nums) -2 ; i< len(nums) && j >= 0 ; i,j = i + 1 ,j - 1 {

		prefix[i] = prefix[i - 1] * nums[i - 1] //1 * 1 , 1 * 2 , 2 * 4  => [1,1,2,8]
		suffix[j] = suffix[j + 1] * nums[j + 1] // 24* 1 , 6 * 4 , 1 * 6 = > [48,24,6,1]

	}	
	result := make([]int, len(nums))
	for i := 0 ; i< len(nums)  ; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}
