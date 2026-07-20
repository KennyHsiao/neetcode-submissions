//division
//Time Complexity : O(n)
//Space Complexity : O(n)
func productExceptSelf(nums []int) []int {

	prod := 1
	zeroCount := 0
	for _, num := range nums {
		if num != 0 {
			prod *= num
		}else{
			zeroCount++
		}
	}

	result := make([]int , len(nums))
	if zeroCount > 1 {
		return result
	}


	for i, num :=range nums {

		if zeroCount > 0 {
			if num == 0 {
				result[i] = prod
			}else {
				result[i] = 0
			}

		}else {
			result[i] = prod / num

		}

	}

	return result
}
