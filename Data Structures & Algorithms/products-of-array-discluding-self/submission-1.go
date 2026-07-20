//prefix & Suffix
//Time Complexity : O(n)
//Space Complexity : O(n)
func productExceptSelf(nums []int) []int {

	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	prefix[0],suffix[len(nums) - 1] = 1 ,1

	for i:= 1 ; i < len(nums) ; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	for j := len(nums) - 2 ; j >= 0 ; j-- {
		suffix[j] = nums[j+1] * suffix[j+1]
	}

	prod := make([]int, len(nums))
	for i := 0; i< len(nums) ; i++{
		prod[i] = prefix[i] * suffix[i]
	}

	return prod
 
}
