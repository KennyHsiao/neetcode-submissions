//brute force hashset
//time complexity : O(n^2)
//space complexity : O(n). seen
func isValidSudoku(board [][]byte) bool {

	for col := 0 ; col < 9 ; col++{
		 seen := make(map[byte]bool)
		 for row := 0 ; row < 9 ; row++{
			 if board[row][col] == '.' {
				continue
			 }

			 if seen[board[row][col]] {
				return false
			 }

			seen[board[row][col]] = true
		 }
	}

	for row := 0 ; row < 9 ; row++{
		 seen := make(map[byte]bool)
		 for col := 0 ; col < 9 ; col++{
			 if board[row][col] == '.' {
				continue
			 }

			 if seen[board[row][col]] {
				return false
			 }

			seen[board[row][col]] = true
		 }
	}

	//handle 3*3 squares
	for squares := 0 ; squares < 9 ; squares++ {
		seen := make(map[byte]bool)
		for i:= 0 ; i < 3 ; i++{
			for j := 0 ; j < 3 ; j++ {
				
				row := (squares/3) * 3 + i
				col := (squares%3) * 3 + j

				if board[row][col] == '.'{
					continue
				}

				if seen[board[row][col]] {
					return false
				}
				seen[board[row][col]] = true

			}

		}
	}


	return true
}
