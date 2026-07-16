//Approach1 :Sorting 
//Time Complexity : O(n * klogk)
//Space Complexity : O(n * k)
func groupAnagrams(strs []string) [][]string {
    

        //slice to store (key,value) string: []string
        slice := make(map[string][]string)

        for _, s := range strs {

            sortedS := sortString(s)
            slice[sortedS] = append(slice[sortedS],s)

        }
    
        var result [][]string
        for _, group := range slice {

            result = append(result, group)
        } 

        return result

}


func sortString(s string) string {

    chars := []rune(s)
    sort.Slice(chars, func(i, j int) bool{
        return chars[i] < chars[j]
    })
    
    return string(chars)

}
