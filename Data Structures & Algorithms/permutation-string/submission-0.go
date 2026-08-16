//U:
// check s2 if contains permutation of s1 , reture true otherwise return false

//M:
//  Since s2 contains continuous permutation of s1 , it come out to match sliding window with fixed size and frequency count 

// sliding windows ：
// In sliding window state:  len(s1) == len(windowSize) 
// 
//. 
// invalid : 每次r++ 去判斷freqCount s1 是否等於 freqCount windowSize


// P
// 
//.  build FreqCount of s1   

	 //expand right through s2
//.  add s2[right] to windowFreq
//.  
//.  if len(windowSize) > len(s1) 
// 		remove windowSize[left]
//.  	left++
    //if len(windowSize)== len(s1) 
    //  if freqWindow == freqTarget
	    // return true
// return false

//I
//R
//s1="ab"
//s2="lecabee"

//E
//Time Complexity O(n)
//Space Complexity O(1)

func checkInclusion(s1 string, s2 string) bool {

    var freqTarget [26]int //[s1[i] - 'a' : freq]
	var freqWindow [26]int
	for _ , ch := range s1 {
		freqTarget[ch - 'a']++
	}
	left,right := 0 , 0

	for right < len(s2) {

		freqWindow[s2[right]- 'a']++
		if right - left + 1 > len(s1){
			freqWindow[s2[left] - 'a']--
			left++
		}

		if right - left + 1 == len(s1){
			if freqTarget == freqWindow{
				return true
			}
		}
		right++
	}
	return false

}
