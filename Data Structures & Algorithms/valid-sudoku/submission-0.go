//brute force
func isValidSudoku(board [][]byte) bool {
	
	for col:= 0 ; col < 9 ; col++ {
		checkCol := make(map[byte]bool)
		for i:= 0 ; i < 9 ; i++ {
			 if board[i][col] == '.' {
				continue
			 }else if checkCol[board[i][col]] {
				return false 
			 }else{
				checkCol[board[i][col]] = true
			 }
		}	
		
	}

	for row:= 0 ; row < 9 ; row++ {
		checkRow := make(map[byte]bool)
		for i:= 0 ; i < 9 ; i++ {
			 if board[row][i] == '.' {
				continue
			 }else if checkRow[board[row][i]] {
				return false 
			 }else{
				checkRow[board[row][i]] = true
			 }
		}	
	}


	//spares 
	for square := 0 ; square < 9 ; square++ {

		checkSquare := make(map[byte]bool)	

		for i:= 0 ; i<3 ; i++{
			for j:= 0 ; j<3 ; j++{
                        //第(square/3) square 的 第i row
				row := (square / 3) * 3+ i 
						//第(square%3) square 的 第j col
 				col := (square % 3)  *3 + j

			if board[row][col] == '.' {
				continue
			 } else if checkSquare[board[row][col]] {
				return false 
			 }

			 checkSquare[board[row][col]] = true

			}
		}

	}
	return true


}
