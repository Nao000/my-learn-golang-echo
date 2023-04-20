package main

import (
	_"fmt"
	"golang.org/x/tour/pic"
)
// go mod init <モジュール名>
// go get golang.org/x/tour/pic@latest
func Pic(dx, dy int) [][]uint8 {

	board := make([][]uint8, dy)
	for i:=0; i < dy; i++{
		board[i] = make([]uint8, dx)
		for j:=0; j < dx; j++{
			board[i][j] = uint8((j + i) / 2)
		}
	}

	return board
}

func main() {
	pic.Show(Pic)
}