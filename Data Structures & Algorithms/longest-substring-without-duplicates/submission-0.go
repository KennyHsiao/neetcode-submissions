// Understand
// → 需要找 longest substring without duplicate characters

// Match
// → 這是 variable-size sliding window
// → HashMap 可以記 last seen index

// Plan
// → 我會用 left/right 維護合法 window
// → right expand
// → duplicate 時調整 left
// → 更新 lastIndex
// → 更新 max length

// Implement
// → 開始 coding
// Time Complexity O(n)
// Space Complexity O(m)
// m <= n
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	//store [char : index]
	strMap := make(map[byte]int)

	maxLength := 1
	l,r := 0,1
	strMap[s[l]] = l
	
	for r < len(s){
		if idx , ok := strMap[s[r]] ; ok && idx >= l { //因為map 是無序的需要多判別idx 
			l = idx + 1 
		}else {
			if r - l + 1> maxLength {
				maxLength = r - l + 1
			}
		}
		strMap[s[r]] = r
		r++ 
	}

	return maxLength
}
