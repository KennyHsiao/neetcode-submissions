//hash map
//Time Complexity O(n ^ 2)
//Space complexity O(n)
func threeSum(nums []int) [][]int {
	
	//sort nums
	sort.Ints(nums)
	//[-4,-1,-1,0,1,2]
	count := make(map[int]int)
	//store nums set in hashmap
	for _,num:= range nums {
		count[num]++
	}

	var result [][]int
	//[[-4:1],[-1:2],[0:1],[1:1],[2:1]]
	//loop
	for i:= 0 ; i< len(nums) ; i++{
		count[nums[i]]--
		//avoid duplicated elementLeft 
		if i>0 && nums[i] == nums[i-1]{
			continue
		}

		
		for j:= i+1 ; j< len(nums) ; j++{
			
			//Ensure avoiding duplicated triplet in third element
			count[nums[j]]--

			if j > i + 1 && nums[j] == nums[j-1]{
				continue
			}
			
			target := -(nums[i] + nums[j])
			//find third element in triplet
			if count[target] > 0 {
				 result = append(result, []int{nums[i], nums[j], target})
			}

		}
		
		for j := i + 1; j < len(nums) ; j++{
			count[nums[j]]++
		}


	}
	
	return result

}
