func hasDuplicate(nums []int) bool {
    
	hashSet := make(map[int]struct{})

	for _, num := range nums{
		hashSet[num] = struct{}{}
	}

	return len(hashSet) < len(nums)
}
