package main

import (
	"fmt"
	"strings"
)

var N int = 4

func show_board(board [4][4]int) {
	line := "+" + strings.Repeat("-", 2*N-1) + "+"
    fmt.Println(line)
	for i := 0; i < N; i++ {
		fmt.Print("|")
		for j := 0; j < N; j++ {
			if board[i][j] == 0 {
				fmt.Print(" |")
			} else if board[i][j] == 1 {
				fmt.Print("●|")
			} else {
				fmt.Print("○|")
			}
		}
		fmt.Println()
	}
	fmt.Println(line)
}

func is_on_board(input [2]int) bool{
    if input[0] >= 0 && input[0] <= N-1 {
		if input[1] >= 0 && input[1] <= N-1 {
            return true
		}else{
			fmt.Println("範囲外です")
			return false
		}
	}else{
		fmt.Println("範囲外です")
		return false 
	}
}

func main() {
	// 0: 空白, 1: 黒, 2: 白
	board := [4][4]int{
		{0, 0, 0, 0},
		{0, 1, 2, 0},
		{0, 2, 1, 0},
		{0, 0, 0, 0},
	} 
	var input [2]int

	show_board(board)
    
    for{
		_, err := fmt.Scan(&input[0], &input[1])
		if err != nil {
			fmt.Println("入力エラー:", err)
			continue
		}
		if is_on_board(input) == true{
			fmt.Println("入力成功!")
			break
		}
	}
	
}
