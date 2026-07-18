//Approach 2 : Hash Table
func groupAnagrams(strs []string) [][]string {

    //declare slice that can store key 
    slice := make(map[[26]int][]string)

    //time complexity O(m * n)
    for _, s := range strs{ 
        var count [26]int
        for _ , char := range s {// char - 'a' represent index of []int
            count[char - 'a']++
        }

        slice[count] = append(slice[count], s)

    }

    var result [][]string
    //loop slice[] to get every []string 
    for _ , strArr := range slice {

        result = append(result, strArr)
    }

    return result


}
