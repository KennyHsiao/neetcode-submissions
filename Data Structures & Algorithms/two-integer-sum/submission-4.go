//U:
//Given an array of interger and find out target and return these indexs of nums
//M
//brute force : 
//loop nums with i , j = 0 , i+1  check i j sum are equal to target
//Time Complexity O(n ^ 2)

//Hash Map :
//create hashMap [num : index]
// loop nums :
// check  diff = target - nums[i] is exist in hashMap , if yes , return index  
// if not , store indices[num] = i
//P
//   
//I
//R
//[3,4,5,6]
//E
//Time Complexity O(n)
//Space Complexity O(n)
func twoSum(nums []int, target int) []int {
    
	numsMap := make(map[int]int)

	for i, num := range nums{

		diff := target - num //4, 3
		if j, ok := numsMap[diff] ; ok {
			return []int{j , i} //[0,1]
		}
		numsMap[num] = i //[3:0]
		

	}

	return []int{}

}
