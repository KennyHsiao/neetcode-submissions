//time complexity O(n)
//space complexity O(1)
func maxArea(heights []int) int {
	//min(heights[i], heights[j]) * abs(heights[i] - heights[j])

	product:= 0
	i,j := 0,len(heights) -1
	for i<j {
		
		temp := min(heights[i] , heights[j]) * (j - i)
		if temp > product {
			product = temp
		}


		if heights[i] > heights[j] {
			j--
		}else if heights[i] <= heights[j] {
			i++
		}
	}

	return product

}


func min(a ,b int) int{
	if a > b {
		return b
	}
	return a
}

