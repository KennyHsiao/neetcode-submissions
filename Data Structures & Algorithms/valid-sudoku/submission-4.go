func isValidSudoku(board [][]byte) bool {
    
	rows := make([]map[byte]bool , 9)
	cols := make([]map[byte]bool , 9)
	squares := make([]map[byte]bool , 9)
   
	//every map needs to be initialized because zero value of map is nil
	//or it'll panic if put value in map
	for i:= 0 ; i<9 ;i++{
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	//use one loop to validate every rows , cols , squares 
	for r:=0 ; r < 9 ; r++{
		for c:= 0 ; c <9 ; c++{

			if board[r][c] == '.'{
				continue
			}
			// (r/3)*3 + c/3
			//r=0 
			//c=0		 c=1		c=2		c=3	        c=4	     c=5	   c=6		 c=7	  c=8 
			//sq0+0      sq0+0     sq0+0    sq0+1      sq0+1     sq0+1.   sq0+2      sq0+2     sq0+2  
			
			//r=1 
			//c=0 			
			//sq0+0      sq0+0     sq0+0    sq0+1      sq0+1     sq0+1.   sq0+2      sq0+2     sq0+2  
			//...
			value := board[r][c]
			squareIndex := (r/3) *3 + c/3

			if rows[r][value] || cols[c][value] {
				return false
			}
			if squares[squareIndex][value] {
				return false
			}

			rows[r][value] = true
			cols[c][value] = true
			squares[squareIndex][value] = true
		}
	}
	return true

}
