//two pointers
// Use two pointers from both ends.
// Maintain the maximum height seen from the left and right.
// Process the side with the smaller boundary.
// If the current height is lower than its side's max,
// the difference is trapped water.
//left < right
//     ↓
// 比較 height[left] 和 height[right]
//     ↓
// 較矮的那一側可以先處理
//     ↓
// 更新該側的 max，或累加 water
//     ↓
// 只移動該側 pointer
func trap(height []int) int {

	if len(height) == 0{
		return 0
	}

	trapWater := 0
	left,right := 0 , len(height) - 1
	leftMax, rightMax := height[left] , height[right]
	for left < right {
		
		if height[left] <= height[right]{
			left++
			
			if height[left] > leftMax{
				leftMax = height[left] //當大於leftMax trapWater = 0 所以不用計算
			}else {
				trapWater += leftMax - height[left]
			}

		}else {
			right--
			
			if height[right] > rightMax {
				rightMax = 	height[right] //當大於rightMax trapWater = 0 所以不用計算
			}else{
				trapWater += rightMax - height[right]
			}			

		}


	}

	return trapWater
}
