//bucket sort
//time complexity : O(n)
//space complexity : O(n)
func topKFrequent(nums []int, k int) []int {

    count := make(map[int]int) 

    for _, num := range nums {
        count[num]++ 
    }


     freq := make([][]int, len(nums)+ 1)
     
     for num, c := range count {
        freq[c] = append(freq[c], num)  //  [[0] , [1] , [2] , [3] , [0] ,[0] ,[0]]
     }                                //index 0.    1     2.    3.    4.   5.   6

     
     var result []int

     for i := len(freq) - 1 ; i > 0 ; i--{ // i : how many time appear  freq[i] : num

            for _ ,num := range freq[i]{
                
                result = append(result, num) 
                if len(result) == k {
                    return result
                }
            }
     }
    
     return []int{}
}
