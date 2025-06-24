package main

import "fmt"

func main() {
	var restart bool = false
	if nPlayer < 26 {
		printErrorScreen("Sesuatu Tidak Beres")
	}

	entryBanner()
	filePlayer := loadInitialPlayers()

	for !restart {
		user := menuAwal(&filePlayer)
		if user.name == "" {
			fmt.Println("Program selesai. Terima kasih sudah bermain!")
			fmt.Println()
			restart = true
		} else {
			data := initGame(user.difficulty)

			// Reset statistik pemain sebelum bermain
			user.hits = 0
			user.numOfTurns = 0
			user.numOfFail = 0

			playGame(user, &data)

			fmt.Print("\nMau main lagi? (y/n): ")
			var again string
			fmt.Scan(&again)
			if again == "n" && again != "N" {
				fmt.Println("Terima kasih sudah bermain!")
				restart = true
			} else if again == "y" && again != "Y" {
				fmt.Println("Kembali ke Menu Awal...")
				fmt.Println()
				entryBanner()
			} else {
				fmt.Println("Pilihan tidak valid, Default ke pilihan 'y'")
				again = "y"
				fmt.Println()
				entryBanner()
			}
		}

	}
}
