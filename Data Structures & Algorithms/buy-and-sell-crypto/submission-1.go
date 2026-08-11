//[10,1,5,6,7,1]
//l = 0 , r = 1 
// l = 0 , r = 2 ...
//key pointer : 當prices[l] >  prices[r] 就可以省略後續的prices[r] - prices[l]
//              因為直接把l 從 r 開始，prices[r] - pricesp[l] 去算  
//time complexity O(n)
//space complexity O(1)
func maxProfit(prices []int) int {
	
	
	l,r := 0,1
	maxProfit := 0
	for r < len(prices) {

		if prices[l] > prices[r] {
			l = r
			
		}else{
			if prices[r] - prices[l] > maxProfit{
				maxProfit = prices[r] - prices[l]
			}
		} 
		r++

	}

	return maxProfit
}
