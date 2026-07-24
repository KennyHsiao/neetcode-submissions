//Time Complexity O(nlogn)
//Space Complexity O(n)
//Sorting 
func longestConsecutive(nums []int) int {
	//hashMap store every unique num
	hashSet := make(map[int]struct{})

	for _, num := range nums {
		hashSet[num]=struct{}{}
	}
	//sorting slice
	sortedNums := make([]int,0,len(hashSet))
	for key:= range hashSet {
		sortedNums = append(sortedNums,key)
	}
	
	if len(sortedNums) == 0{
		return 0
	}

	sort.Ints(sortedNums)

	//count longest consecutive sequence of elements
	longestCount := 1
	currentCount := 1
	for i:= 0 ; i< len(sortedNums)-1 ; i++ {
		
		if sortedNums[i+1] - sortedNums[i] == 1 {
			currentCount++
		}else {
			if currentCount > longestCount {
				longestCount = currentCount
			}
			currentCount = 1
		}
	}	
	if currentCount > longestCount {
		longestCount = currentCount
	}
	return 	longestCount

}
