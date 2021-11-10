package main

import (
	"fmt"
)

func main() {
	var (
		board [8][8]string
		black string = "kqrbnp"
		white string = "KQRBNP"
	)

	board[0][0] = "r"
	board[0][7] = "r"

	for i := range black {
		board[0][i+1] = string(black[i])
		board[7][i+1] = string(white[i])
	}

	func() {
		for i := 0; i < len(board[1]); i++ {
			fmt.Println(board[:][i])
		}
	}()
}
