package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var currentCount int = 0

const nPlayer int = 128

type player struct {
	id         byte
	name       string
	hits       int
	numOfTurns int
	difficulty int
	numOfFail  int
}

type tabPlayer [nPlayer]player

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func loadInitialPlayers() tabPlayer {
	var tp tabPlayer
	rand.Seed(time.Now().UnixNano())

	var names = []string{
		"Aidan", "Ben", "Clay", "Doni", "Erlan", "Fredy", "Geral", "Henry", "Ian", "John",
		"Ken", "Liam", "Marcello", "Nathan", "Orland", "Pier", "Que", "Ron", "Steven", "Thor",
		"Ulric", "Vyel", "William", "Xavier", "Yhe", "Zod",
	}

	//mengrandom urutan nama
	for i := len(names) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		names[i], names[j] = names[j], names[i]
	}

	//random numOfTurns
	for i := 0; i < len(names); i++ {
		diffFile := rand.Intn(2) + 1 // 1 = Easy, 2 = Hard
		var maxTurn int
		if diffFile == 1 {
			maxTurn = 9
		} else {
			maxTurn = 16
		}

		tp[i] = player{
			id:         byte(i),
			name:       names[i],
			hits:       4,
			numOfTurns: rand.Intn(maxTurn-3) + 4, // minimal 4 turn
			difficulty: diffFile,
			numOfFail:  rand.Intn(2), // 0 atau 1 gagal
		}
	}
	currentCount = len(names)
	return tp
}

func menuAwal(filePlayer *tabPlayer) *player {
	var nameInput string
	var validName bool = false
	var userChoice string
	var diff int
	var finished bool = false

	for !finished {
		fmt.Println(topBottomBorder)
		fmt.Println("|                                    Pilih Pilihan Anda  (1/2/3/0)                                    |")
		fmt.Println("|   1. Play                                                                                           |")
		fmt.Println("|   2. LeaderBoard                                                                                    |")
		fmt.Println("|   3. Lihat skor anda                                                                                |")
		fmt.Println("|   0. Exit                                                                                           |")
		fmt.Println("|                                                                                                     |")
		fmt.Println(topBottomBorder)
		fmt.Print("Pilih: ")
		fmt.Scan(&userChoice)

		switch userChoice {
		case "1":
			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')

			for !validName {
				var nameStatus bool = false

				for !nameStatus {
					fmt.Print("Masukkan Nama: ")
					nameInput, _ = reader.ReadString('\n')
					nameInput = strings.TrimSpace(nameInput)

					if nameInput != "" && nameInput != " " {
						nameStatus = true
						fmt.Println()
					} else {
						fmt.Println(topBottomBorder)
						fmt.Println(formatCenteredLine("Nama tidak boleh kosong"))
						fmt.Println(topBottomBorder)
						fmt.Print("\n")
					}
				}
				index := findPlayer(nameInput, *filePlayer)
				if index != -1 {
					fmt.Println(topBottomBorder)
					fmt.Println(formatCenteredLine(fmt.Sprintf("Nama '%s' sudah ada. Lanjut timpa? (y/n):", nameInput)))
					fmt.Println(topBottomBorder)
					fmt.Print("Pilih: ")
					confirm, _ := reader.ReadString('\n')
					confirm = strings.TrimSpace(strings.ToLower(confirm))

					if confirm == "y" {
						fmt.Println(topBottomBorder)
						fmt.Println(formatCenteredLine("Data akan di-overwrite."))
						fmt.Println(topBottomBorder)
						// Lanjut ke permainan dengan overwrite data lama
						(*filePlayer)[index] = player{name: nameInput, difficulty: -1}
						validName = true
					} else if confirm == "n" {
						fmt.Println(topBottomBorder)
						fmt.Println(formatCenteredLine("Silakan masukkan nama lain."))
						fmt.Println(topBottomBorder)
					} else {
						fmt.Println(topBottomBorder)
						fmt.Println(formatCenteredLine("Pilihan tidak valid, Default ke pilihan 'n'"))
						fmt.Println(topBottomBorder)
						confirm = "n"
					}
				} else {
					fmt.Println(topBottomBorder)
					fmt.Println(formatCenteredLine(fmt.Sprintf("Silakan bermain %s", nameInput)))
					fmt.Println(topBottomBorder)
					validName = true
				}
			}

			fmt.Println(formatCenteredLine("Pilih difficulty:"))
			fmt.Println(formatCenteredLine("1. Easy (3x3)"))
			fmt.Println(formatCenteredLine("2. Hard (4x4)"))
			fmt.Println(topBottomBorder)
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&diff)
			fmt.Println()

			if diff != 1 && diff != 2 {
				fmt.Println(topBottomBorder)
				fmt.Println(formatCenteredLine("Pilihan tidak valid, Default ke Easy. Selamat Bermain!"))
				diff = 1
			}

			// Tambah player baru ke array dan berikan ID berikutnya
			if currentCount < nPlayer {
				filePlayer[currentCount] = player{
					id:         byte(currentCount),
					name:       nameInput,
					hits:       0,
					numOfTurns: 0,
					difficulty: diff,
				}
				currentCount++
				return &filePlayer[currentCount-1]
			} else {
				fmt.Println(topBottomBorder)
				fmt.Println(formatCenteredLine("Maksimum player sudah tercapai!"))
				fmt.Println(topBottomBorder)
				return &player{name: "", difficulty: -1}
			}

		case "2":
			var exitLeaderboard bool = false
			for !exitLeaderboard {
				fmt.Println(topBottomBorder)
				fmt.Println(formatCenteredLine("LEADERBOARD MENU"))
				fmt.Println(formatCenteredLine("")) // Space
				fmt.Println("|   1. Tampilkan Leaderboard (original)                                                               |")
				fmt.Println("|   2. Urutkan berdasarkan Nama (ascending)                                                           |")
				fmt.Println("|   3. Urutkan berdasarkan Turns (decending)                                                          |")
				fmt.Println("|   0. Kembali ke menu utama                                                                          |")
				fmt.Println(formatCenteredLine("")) // Space
				fmt.Println(topBottomBorder)

				var sortChoice string
				fmt.Print("Pilih (1/2/3/0): ")
				fmt.Scan(&sortChoice)
				fmt.Println()

				if sortChoice == "1" {
					printLeaderboard(*filePlayer)
				} else if sortChoice == "2" || sortChoice == "3" {
					temp := copyPlayerList(*filePlayer)
					if sortChoice == "2" {
						sortByName(&temp)
					} else {
						sortByTurns(&temp)
					}
					printLeaderboard(temp)

					var subChoice string
					var exitSubMenu bool = false

					for !exitSubMenu {
						fmt.Println(topBottomBorder)
						fmt.Println(formatCenteredLine("LEADERBOARD MENU"))
						fmt.Println(formatCenteredLine(""))
						fmt.Println("|   1. Cari Player berdasarkan Nama                                                                   |")
						fmt.Println("|   2. Hapus Player berdasarkan Nama                                                                  |")
						fmt.Println("|   0. Kembali ke menu Leaderboard                                                                    |")
						fmt.Println(formatCenteredLine(""))
						fmt.Println(topBottomBorder)
						fmt.Print("Pilih (1/2/0): ")
						fmt.Scan(&subChoice)
						fmt.Println()

						switch subChoice {
						case "1":
							fmt.Print("Masukkan Nama player yang ingin dicari: ")
							var nameInput string
							fmt.Scan(&nameInput)

							sortByName(&temp)
							idx := findPlayerByNameBinary(nameInput, temp)
							if idx != -1 {
								p := temp[idx]
								fmt.Println(topBottomBorder)
								fmt.Println(formatCenteredLine("Player ditemukan:"))
								fmt.Println(topBottomBorder)

								// Header dengan kolom tetap
								header := fmt.Sprintf(" %-10s | %-30s | %-10s | %-10s ", "ID", "Nama", "Turns", "Difficulty")
								fmt.Printf("|%s|\n", header)
								fmt.Println(topBottomBorder)

								// Isi data dengan padding sesuai kolom
								diffStr := "Easy"
								if p.difficulty == 2 {
									diffStr = "Hard"
								}
								row := fmt.Sprintf(" %-10s | %-30s | %-10d | %-10s ", fmt.Sprintf("%08b", p.id), p.name, p.numOfTurns, diffStr)
								fmt.Printf("|%s|\n", row)
								fmt.Println(topBottomBorder)
							} else {
								fmt.Println(topBottomBorder)
								fmt.Println(formatCenteredLine("Player dengan nama tersebut tidak ditemukan."))
								fmt.Println(topBottomBorder)
							}
						case "2":
							fmt.Print("Masukkan Nama player yang ingin dihapus: ")
							var nameInput string
							fmt.Scan(&nameInput)
							sortByName(filePlayer) //tidak memakai temp, karena ingin menghapus file array asli
							idx := findPlayerByNameBinary(nameInput, *filePlayer)
							if idx != -1 {
								deletePlayer(idx, filePlayer)
								fmt.Println(topBottomBorder)
								fmt.Println(formatCenteredLine("Player berhasil dihapus."))
								fmt.Println(topBottomBorder)
							} else {
								fmt.Println(topBottomBorder)
								fmt.Println(formatCenteredLine("Player dengan nama tersebut tidak ditemukan."))
								fmt.Println(topBottomBorder)
							}
							// update temp supaya sama dengan data terbaru
							temp = copyPlayerList(*filePlayer)
							if sortChoice == "2" {
								sortByTurns(&temp)
							} else {
								sortByName(&temp)
							}
							printLeaderboard(temp)
						case "0":
							exitSubMenu = true
						default:
							fmt.Println("Pilihan tidak valid.")
							fmt.Println()
						}
					}
				} else if sortChoice == "0" {
					exitLeaderboard = true
					fmt.Println("Kembali ke Menu Awal...")
					entryBanner()
				} else {
					fmt.Println("Pilihan tidak valid.")
					fmt.Println()
				}
			}
		case "3":
			fmt.Print("Masukkan Nama Anda: ")
			fmt.Scan(&nameInput)
			idx := findPlayer(nameInput, *filePlayer)
			if idx != -1 {
				p := filePlayer[idx]
				if p.difficulty == 1 {
					scoreMessage := fmt.Sprintf("Nama: %s | Turns: %d | Difficulty: Easy", p.name, p.numOfTurns)
					fmt.Println(topBottomBorder)
					fmt.Println(formatCenteredLine(scoreMessage))
					fmt.Println(topBottomBorder)
					fmt.Println("Kembali ke Menu Awal...")
					entryBanner()
				} else {
					scoreMessage := fmt.Sprintf("Nama: %s | Turns: %d | Difficulty: Hard", p.name, p.numOfTurns)
					fmt.Println(topBottomBorder)
					fmt.Println(formatCenteredLine(scoreMessage))
					fmt.Println(topBottomBorder)
					fmt.Println("Kembali ke Menu Awal...")
					entryBanner()
				}
			} else {
				fmt.Println(topBottomBorder)
				fmt.Println(formatCenteredLine("Nama tidak ditemukan."))
				fmt.Println(topBottomBorder)
				fmt.Println("Kembali ke Menu Awal...")
				entryBanner()
			}

		case "0":
			finished = true

		default:
			fmt.Println("Pilihan tidak valid. Silahkan pilih Pilihan sesuai instruksi.")
			fmt.Println()
			entryBanner()
		}
	}

	// Kembalikan player kosong jika belum main
	return &player{name: "", difficulty: -1}
}

// cari nama pada saat input nama, sequential search
func findPlayer(playerX string, A tabPlayer) int {
	var idx int = -1
	var i int

	for i = 0; i < nPlayer; i++ {
		if A[i].name == playerX {
			idx = i
		}
	}

	return idx
}

// menggunakan algoritma selection sort
func sortByTurns(filePlayer *tabPlayer) {
	var n = len(filePlayer)

	for pass := 0; pass < n; pass++ {
		var min = pass
		for j := pass; j < n; j++ {
			if filePlayer[j].numOfTurns > filePlayer[min].numOfTurns {
				min = j
			}
		}
		filePlayer[pass], filePlayer[min] = filePlayer[min], filePlayer[pass]
	}
}

// sorting leaderboard secara ascending, selection sort
func sortByName(filePlayer *tabPlayer) {
	n := currentCount
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if filePlayer[j].name < filePlayer[min].name {
				min = j
			}
		}
		filePlayer[i], filePlayer[min] = filePlayer[min], filePlayer[i]
	}
}

// mencari nama dalam leaderboard, binary search
func findPlayerByNameBinary(name string, filePlayer tabPlayer) int {
	low := 0
	high := currentCount - 1
	idx := -1

	for low <= high && idx == -1 {
		mid := (low + high) / 2
		if filePlayer[mid].name == name {
			idx = mid
		} else if filePlayer[mid].name < name {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

// delete player yang di leaderboard
func deletePlayer(idx int, filePlayer *tabPlayer) {
	for i := idx; i < currentCount-1; i++ {
		filePlayer[i] = filePlayer[i+1]
		filePlayer[i].id = byte(i)
	}
	filePlayer[currentCount-1] = player{}
	currentCount--
}

// copy file nama
func copyPlayerList(src tabPlayer) tabPlayer {
	var dst tabPlayer
	copy(dst[:], src[:])
	return dst
}
