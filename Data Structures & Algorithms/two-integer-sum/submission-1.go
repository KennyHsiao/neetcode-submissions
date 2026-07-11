//Sorting with two pointer
func twoSum(nums []int, target int) []int {
    //declare slice with length nums, where each element is 
    //a fixed-size array of two integers
    A := make([][2]int, len(nums))
    
    //loop through the input array and build a slice of pairs
    //Each pair contains the number and its original index
    for i, num := range nums {
        A[i] = [2]int{num, i}
    }
    //sorting slice A by the first element of each pair,which is original number value
    sort.Slice(A, func(i , j int)bool{
        return A[i][0] < A[j][0]
    })

    left, right := 0,len(nums) - 1

    for left < right{
        cur:= A[left][0] + A[right][0]
        if cur == target {
            if A[left][1] < A[right][1]{
                return []int{ A[left][1], A[right][1]}
            }else {
                return []int{ A[right][1], A[left][1]}
            }
        }else if cur < target {
            left++
        }else {
            right--
        }
    }

    return []int{}
}
