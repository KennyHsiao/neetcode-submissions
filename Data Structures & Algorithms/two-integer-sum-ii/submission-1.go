//time complexity O(n)
//space complexity O(1)
func twoSum(numbers []int, target int) []int {

	if len(numbers) == 0 {
		return []int{}
	}

	left, right := 0 , len(numbers) - 1
	for left < right {

		if numbers[left] + numbers[right] < target {
			left++
		}else if numbers[left] + numbers[right] > target {
			right--
		}else {
			return []int{left + 1,right + 1}
		}

	}
	return []int{}

}
