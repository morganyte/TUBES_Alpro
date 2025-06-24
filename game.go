package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type gameInfo struct {
	row, column int
	ships       [][]int
}

func initGame(difficulty int) gameInfo {
	var boardSize int = 3
	if difficulty == 2 {
		boardSize = 4
	}

	ships := make([][]int, boardSize)
	for i := range ships {
		ships[i] = make([]int, boardSize)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	placeShips(ships, r)

	return gameInfo{
		ships: ships,
	}
}

func playGame(user *player, data *gameInfo) {
	reader := bufio.NewReader(os.Stdin)

	// Bersihkan buffer leftover jika ada (original line)
	reader.ReadString('\n')

	for user.hits < 4 && user.numOfFail < 2 {
		fmt.Println(topBottomBorder)
		fmt.Println(formatCenteredLine("Pilih Koordinat (format: x y)"))
		fmt.Println(topBottomBorder)
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		fmt.Println()

		if len(parts) < 2 {
			fmt.Println(topBottomBorder)
			fmt.Println(formatCenteredLine("Input kurang lengkap! Masukkan 2 angka."))
			fmt.Println(topBottomBorder)
			user.numOfFail++
		} else {
			var errX, errY error
			data.row, errX = strconv.Atoi(parts[0])
			data.column, errY = strconv.Atoi(parts[1])

			if errX != nil || errY != nil {
				fmt.Println(topBottomBorder)
				fmt.Println(formatCenteredLine("Input tidak valid! Masukkan angka."))
				fmt.Println(topBottomBorder)
				user.numOfFail++
			} else {
				if data.row < 0 || data.row >= len(data.ships) || data.column < 0 || data.column >= len(data.ships[0]) {
					fmt.Println(topBottomBorder)
					fmt.Println(formatCenteredLine("Koordinat di luar batas! Coba lagi."))
					fmt.Println(topBottomBorder)
					user.numOfFail++
				} else {
					if data.ships[data.row][data.column] == 2 {
						fmt.Println(topBottomBorder)
						fmt.Println(formatCenteredLine("Kamu sudah menembaknya"))
						user.numOfTurns++
					} else {
						if data.ships[data.row][data.column] == 1 {
							data.ships[data.row][data.column] = 2
							user.hits++
							user.numOfTurns++
							fmt.Println(topBottomBorder)
							fmt.Println(formatCenteredLine(fmt.Sprintf("KENA!!! Tersisa %d target lagi!", 4-user.hits)))
						} else {
							data.ships[data.row][data.column] = 2
							fmt.Println(topBottomBorder)
							fmt.Println(formatCenteredLine("Tidak Kena... hehehe... coba lagi"))
							user.numOfTurns++
						}
					}
				}
			}
		}
	}

	if user.numOfFail >= 2 {
		defeatBanner()
		fmt.Println(topBottomBorder)
		fmt.Println(formatCenteredLine("Gagal terlalu banyak. Game over!"))
		fmt.Println(topBottomBorder)
	} else {
		victoryBanner()
		fmt.Println(topBottomBorder)
		fmt.Println(formatCenteredLine(fmt.Sprintf("Selamat %s, Anda menang dalam %d kali tembak!", user.name, user.numOfTurns)))
		fmt.Println(topBottomBorder)
	}
}

func placeShips(board [][]int, r *rand.Rand) { // r bagian dari *rand.Rand sebagai initGame
	placed := 0
	for placed < 4 {
		x := r.Intn(len(board))
		y := r.Intn(len(board[0]))
		if board[x][y] == 0 {
			board[x][y] = 1
			placed++
		}
	}
}

func printLeaderboard(players tabPlayer) {
	var diffStr string

	fmt.Println(leaderboardTitleBorder) // Special border untuk judul leaderboard

	// Header format: bagian konten 101 chars
	// "| %-30s | %-30s | %-33s |" -> 1(space)+30(col1)+3( | )+30(col2)+3( | )+33(col3)+1(space) = 101
	headerContent := fmt.Sprintf(" %-30s | %-30s | %-33s ", "Nama", "Difficulty", "Turns")
	fmt.Printf("|%s|\n", headerContent)
	fmt.Println(topBottomBorder) // Separator line

	// Mengulang seluruh iterasi dan memeriksa nama yang tidak kosong
	for i := 0; i < len(players); i++ {
		if players[i].name != "" { // Original check
			if players[i].difficulty == 2 {
				diffStr = "Hard"
			} else {
				diffStr = "Easy"
			}
			// Format baris data, mencocokkan struktur header
			dataContent := fmt.Sprintf(" %-30s | %-30s | %-33d ", players[i].name, diffStr, players[i].numOfTurns)
			fmt.Printf("|%s|\n", dataContent)
		}
	}
	fmt.Println(topBottomBorder)
}

func printErrorScreen(msg string) {
	clearScreen() // memakai original clearScreen() jika ada error
	fmt.Println(topBottomBorder)
	fmt.Println(formatCenteredLine("ERROR: " + msg + " :ERROR"))
	fmt.Println(topBottomBorder)
	os.Exit(1) // keluar dari program (original line)
}
