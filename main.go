package main

import (
	"fmt"
	"strings"
	"os"
)

var N int = 4
var reverse [][2]int

func show_board(board [][]int) {
	fmt.Print(" ")
	for i := 0; i < N; i++ {
		fmt.Print(" ")
		fmt.Print(i)
	}
	fmt.Print("\n")
	line := " +" + strings.Repeat("-", 2*N-1) + "+"
    fmt.Println(line)
	for i := 0; i < N; i++ {
		fmt.Print(i)
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

func is_on_board(x int, y int) bool{
    if x >= 0 && x <= N-1 {
		if y >= 0 && y <= N-1 {
            return true
		}else{
			return false
		}
	}else{
		return false 
	}
}

func isMyColorAppearAgain(board [][]int, x int, y int, i int, j int, bante int) {
	var addreverse [][2]int
	var temp [2] int
	for{
		x += i
		y += j
		if is_on_board(x, y) == false {
            return
		}else{
			if board[x][y] == bante{
				for i := 0; i < len(addreverse); i++{
					reverse = append(reverse, addreverse[i])
				}
                return 
			}else if board[x][y] == 0{
				return
			}else{
				temp[0] = x
				temp[1] = y
				addreverse = append(addreverse, temp)
				continue
			}
		}
	}
}

func reversible(board [][]int, x int, y int, bante int) bool{
	if board[x][y] != 0 {
		fmt.Println("すでに置かれています")
		return false
	}else{
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {		
                if is_on_board(x+i, y+j) == false {
                    continue
				}else{
					if board[x+i][y+j] != (3 - bante) {
                        continue
					}else{
						isMyColorAppearAgain(board, x, y, i, j, bante)
					}
				}
			}
		}
		if reverse == nil {
			return false
		}else{
			return true
		}
	}
}

func updateboard(board [][]int, x int, y int, bante int) [][]int{
    for i := 0; i < len(reverse); i++{
		board[reverse[i][0]][reverse[i][1]] = bante
	}
	board[x][y] = bante
	reverse = nil
	return board
}

func show_win_lose(board [][]int){
	var white int = 0
	var black int = 0
	for i :=0; i < N; i += N-1{
		for j := 0; j < N; j += N-1{
			if board[i][j] == 1{
				black += 1
			}else if board[i][j] == 2{
				white += 1 
			}
		}
	} 
	if white==black{
		fmt.Println("引き分けです")
		os.Exit(0)
	}else if white > black{
		fmt.Println("白の勝ちです")
		os.Exit(0)
	}else{
		fmt.Println("黒の勝ちです")
		os.Exit(0)
	}
}

func main() {
	// 0: 空白, 1: 黒, 2: 白
	var number int
	var board [][]int
	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Println("このオセロは角を多く取った方が勝ちです。角を取った数が同数の場合、引き分けとなります。")
	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Println("盤面のサイズを決めてください")
	fmt.Println("(1)4×4, (2)6×6 (3)8×8")
	fmt.Println("1から3以外の数字を入力した場合は8×8になります")
	fmt.Scan(&number)
	if number == 1 {
		board = [][]int{
		{0, 0, 0, 0},
		{0, 2, 1, 0},
		{0, 1, 2, 0},
		{0, 0, 0, 0},
		}
	}else if number == 2 {
		board = [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 2, 1, 0, 0},
		{0, 0, 1, 2, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		}
		N = 6
	}else{
		board = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 1, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	    }
		N = 8
	} 
	var x int
	var y int
	bante := 1
	var flag bool = false

	show_board(board)
    
    for{
		can_put := false
		is_reversible := false
		for i:=0; i < N; i++ {
			for j:=0; j < N; j++ {
				if board[i][j] == 0 {
					is_reversible = reversible(board, i, j, bante)
				}
				if is_reversible == true {
					can_put = true
					break
				}
			}
			if can_put == true {
				break
			}
		}
		reverse = nil
		if can_put == false {
			if bante == 1{
				fmt.Println("黒はスキップされました")
			}else{
				fmt.Println("白はスキップされました")
			}
			if flag == true {
				break
			}
			bante = 3 - bante
			flag = true
			continue
		}
		if bante == 1{
			fmt.Println("現在は黒の番です")
		}else{
			fmt.Println("現在は白の番です")
		}
		fmt.Println("縦と横の数字を順に入力してください")
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			fmt.Println("入力エラー:", err)
			continue
		}
		if is_on_board(x, y) == false{
			fmt.Println("範囲外です")
		}else{
			is_reversible := reversible(board, x, y, bante)
			if (is_reversible == true) {
            	board = updateboard(board,x, y, bante)
				flag = false
				show_board(board)
				if bante == 1 {
					bante = 2
				}else{
					bante = 1
				}
			}else{
				fmt.Println("挟めません")
			}
		}
	}
	show_win_lose(board)
}
