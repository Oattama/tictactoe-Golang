package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func main() {
	var board = [][] string {
		{" ","|"," ","|"," "},{"-","+","-","+","-"},
		{" ","|"," ","|"," "},{"-","+","-","+","-"},
		{" ","|"," ","|"," "}}

	menu(board)
}

func menu(board [][] string) {
	var ans string

	var stop bool = false

	fmt.Println("Start tictactoe Yes/No?")

	fmt.Scanln(&ans)

	if(strings.ToLower(ans) == "yes") {
		
		printGameBoard(board)

		for stop != true {

			playerPosi(board)
			ranCPU(board)
			fmt.Println("")
			//Check gameOver
			gameStop := checkWinning(board)
			if(gameStop > 0) {
				stop = true
			}
		}
	}
}

// Player turn
func playerPosi (board [][] string) {
	var pos int

	fmt.Scanln(&pos)
	
	if(pos < 10 && pos >= 1) {
		position(board, pos, "player")
		fmt.Println("")

		return
	} else {
		fmt.Println("Please Enter Number 1-9!!!")
		playerPosi(board)

		return
	}

	return
}

// CPU turn
func ranCPU (board [][] string) {
	var pos int

	rand.Seed(time.Now().UnixNano())
	pos = rand.Intn(10-1)+1
	position(board, pos, "CPU")
	fmt.Println("")

	return
}

func position(board [][] string, pos int, user string) {

	var symbol string

	if(user == "player") {
		symbol = "X"
	} else {
		symbol = "O"
	}

	switch pos{
	case 1:
		board = checkPositionExist(board, 0, 0, symbol)
		break
	case 2:
		board = checkPositionExist(board, 0, 2, symbol)
		break
	case 3:
		board = checkPositionExist(board, 0, 4, symbol)
		break
	case 4:
		board = checkPositionExist(board, 2, 0, symbol)
		break
	case 5:
		board = checkPositionExist(board, 2, 2, symbol)
		break
	case 6:
		board = checkPositionExist(board, 2, 4, symbol)
		break
	case 7:
		board = checkPositionExist(board, 4, 0, symbol)
		break
	case 8:
		board = checkPositionExist(board, 4, 2, symbol)
		break
	case 9:
		board = checkPositionExist(board, 4, 4, symbol)
		break
	}

	printGameBoard(board)

}

func checkPositionExist (board [][] string, x int, y int, symbol string) [][] string {
	if(board[x][y] == " ") {
		board[x][y] = symbol
	} else {
		if(symbol != "X") {
			ranCPU(board)
		} else {
			fmt.Println("This position is already exist!!!")
			playerPosi(board)
		}
	}

	return board
}

func checkWinning (board [][] string) int {

	wins := [8]string{"", "", "", "", "", "", "", ""}

	wins[0] = board[0][0] + board[0][2] + board[0][4]
	wins[1] = board[2][0] + board[2][2] + board[2][4]
	wins[2] = board[4][0] + board[4][2] + board[4][4]
	wins[3] = board[0][0] + board[2][0] + board[4][0]
	wins[4] = board[0][2] + board[2][2] + board[4][2]
	wins[5] = board[0][4] + board[2][4] + board[4][4]
	wins[6] = board[0][0] + board[2][2] + board[4][4]
	wins[7] = board[0][4] + board[2][2] + board[4][0]

	for idx, _ := range wins {
		
		if(wins[idx] == "XXX") {

			fmt.Println("Player wins")

			return 1

		} else if (wins[idx] == "OOO") {

			fmt.Println("CPU wins")

			return 1

		} 
	}
	return 0
}

func printGameBoard(board [][] string) {

	for i := range board {
		fmt.Println(board[i])
	}

}
