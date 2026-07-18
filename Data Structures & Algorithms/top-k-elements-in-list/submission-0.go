// n = len(nums)
// m = is the number of unique element
// time complexity O(n + mlogm) . worst case : nlogn
// space complexity O(m). worst case : O(b)
func topKFrequent(nums []int, k int) []int {

    //declare a map to get frequency from integer array
    mapFreq := make(map[int]int)

    for _, num := range nums {
        mapFreq[num]++  
    }

    pairs := make([][2]int, 0 , len(mapFreq))

    for num, count := range mapFreq {
        pairs = append(pairs, [2]int{num , count}) 
    }

    sort.Slice(pairs , func(i , j int)bool{
        return pairs[i][1] > pairs[j][1]
    })

    var result []int
    for i:= 0 ; i < k ; i++ {
        result = append(result, pairs[i][0])
    }

    return result
}
