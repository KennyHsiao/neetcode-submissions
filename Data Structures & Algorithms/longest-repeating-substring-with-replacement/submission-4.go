//Understand:
// Given a string s and interger k
// I can replace  at most k characters and find the longest substring that can be changed into all the same characters.

//Match: 
//variable-size sliding windows + Frequency count


//Plan:
// right ++ expand
// update new char frequency
// know maxFreq
// windows size - maxFreq <= k
// if <= k  valid window , update longest
// else invalid window , left++ , frequency count--

//Review :
//s="XYYX"
//k=2
//s="AABABBA"
//k=1

//Evaluation
//Time Complexity O(n)
//Space Complexity O(1)
func characterReplacement(s string, k int) int {

	if len(s) == 1 {
		return 1
	}
	freqCount := make(map[byte]int)
	maxFreq := 0
	longestLen := 0

	l,r := 0 , 0
	
	for r < len(s) { //r = 2 l = 0
		freqCount[s[r]]++  //[A: 2][B: 1]

		maxFreq = max(maxFreq, freqCount[s[r]]) //maxFreq : 2

		if r - l  + 1 - maxFreq <= k { //k = 1 check valid window 

			longestLen = max(longestLen , r - l + 1) // longestLen = 2
			r++

		}else{
			freqCount[s[l]]-- 
			l++
			r++
		}
	}

	return longestLen
}


func max(a, b int) int{

	if a > b {
		return a
	}
	return b

} 
