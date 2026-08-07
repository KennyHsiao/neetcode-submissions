//prefix and suffix 
//leftMax[i] = max(height[i], leftMax[i-1]) , start from i = 1
	//leftMax[i] = [0,2,2,3,3,3,3,3,3,3]
	//			   [0,2,0,3,1,0,1,3,2,1]
	//rightMax[i] = max(height[i], rightMax[i+1]]) ,start from i = len(height) -2
	//rightMax[i]= [2,2,3,3,1,1,3,3,2,1]
	//how much water can be traped 
	//min(leftMax[i] , rightMax[i]) - height[i]
// Time complexity O(n)
// Space compleixty O(n)
func trap(height []int) int {
	
	if len(height) == 0 {
		return 0
	}

	leftMax := make([]int, len(height))
	rightMax := make([]int ,  len(height))

	for i:= 0 ; i < len(height) ; i++{
		if i == 0 {
			leftMax[i] = height[i]
		}else {
			leftMax[i] = max(leftMax[i - 1] , height[i])
		}

	}


	for i:= len(height) - 1 ; i > -1 ; i--{ 
		if i == len(height) - 1 {
			rightMax[i] = height[i]
		}else {
			rightMax[i] = max(rightMax[i+1] , height[i])
		}

    }	

	trapWater := 0
	for i := 0 ; i < len(height) ; i++{
		
		trapWater += min(leftMax[i] ,rightMax[i]) - height[i]

	}

	return trapWater
}


func min(a , b int) int{
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}