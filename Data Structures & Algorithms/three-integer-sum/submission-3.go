//two pointers
//time complexity O(n^2)
//space complexity O(1)
func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	result := make([][]int, 0)
	for i:= 0 ; i< len(nums); i++{ //i = 4
		//skip duplicated target
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		left , right := i+1 , len(nums) -1 //4,4
		
		for left < right {
			threeSum := nums[i] + nums[left] + nums[right] 
			//skip duplicated left, right from i
			if threeSum > 0 {
				right-- //3
			}else if threeSum < 0 {
				left++
			}else {
				result = append(result, []int{nums[i], nums[left] , nums[right]})
				left++ // 2
				right-- // 3
				for left < right &&nums[left] == nums[left - 1]{
					left++
				}

				for left < right && nums[right+ 1] == nums[right]{
					right--
				} 
				
				
			}
		}
	}

	return result

}
