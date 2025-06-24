package main

import (
	"fmt"
	"strings"
)

const (
	contentWidth           = 101
	topBottomBorder        = "+=====================================================================================================+"
	leaderboardTitleBorder = "+====================================== LEADERBOARD ==================================================+"
)

// fungsi pembantu untuk membuat baris dengan teks di tengah, diapit oleh pipa
func formatCenteredLine(text string) string {
	// Memastikan teks tidak lebih panjang dari lebar konten, jika demikian, potong (penanganan dasar)
	if len(text) > contentWidth {
		text = text[:contentWidth-3] + "..." // menunjukkan bahwa teks terpotong
	}
	padding := (contentWidth - len(text)) / 2
	remainder := (contentWidth - len(text)) % 2 // untuk menangani ganjil/genap panjang perbedaan untuk perfect centering
	return fmt.Sprintf("|%s%s%s|", strings.Repeat(" ", padding), text, strings.Repeat(" ", padding+remainder))
}

func entryBanner() {
	fmt.Println()
	fmt.Println(topBottomBorder)
	fmt.Println("|                                                                                                     |")
	fmt.Println("|                             █▄▄ ▄▀█ ▀█▀ ▀█▀ █░░ █▀▀ █▀ █░█ █ █▀█                                    │")
	fmt.Println("|                             █▄█ █▀█ ░█░ ░█░ █▄▄ ██▄ ▄█ █▀█ █ █▀▀                                    │")
	fmt.Println("|                                                                                                     |")
	fmt.Println("|                                           Created by :                                              |")
	fmt.Println("|                                                                                                     |")
	fmt.Println("|                     1. William Peter Vanx Najoan                                                    |")
	fmt.Println("|                     2. Muhammad Faris Dhiyayl Haq Sarbini                                           |")
	fmt.Println("|                                                                                                     |")
}

func victoryBanner() {
	fmt.Println(topBottomBorder)
	fmt.Println("|                                                                                                     |")
	fmt.Println("|                                   █░█ █ █▀▀ ▀█▀ █▀█ █▀█ █▄█ █                                       |")
	fmt.Println("|                                   ▀▄▀ █ █▄▄ ░█░ █▄█ █▀▄ ░█░ ▄                                       |")
	fmt.Println("|                                                                                                     |")
}

func defeatBanner() {
	fmt.Println()
	fmt.Println(topBottomBorder)
	fmt.Println("|                                                                                                     |")
	fmt.Println("|                          ▓█████▄ ▓█████   █████▒▓█████ ▄▄▄     ▄▄▄█████▓ ▐██▌                       │")
	fmt.Println("|                          ▒██▀ ██▌▓█   ▀ ▓██   ▒ ▓█   ▀▒████▄   ▓  ██▒ ▓▒ ▐██▌                       │")
	fmt.Println("|                          ░██   █▌▒███   ▒████ ░ ▒███  ▒██  ▀█▄ ▒ ▓██░ ▒░ ▐██▌                       │")
	fmt.Println("|                          ░▓█▄   ▌▒▓█  ▄ ░▓█▒  ░ ▒▓█  ▄░██▄▄▄▄██░ ▓██▓ ░  ▓██▒                       │")
	fmt.Println("|                          ░▒████▓ ░▒████▒░▒█░    ░▒████▒▓█   ▓██▒ ▒██▒ ░  ▒▄▄                        │")
	fmt.Println("|                           ▒▒▓  ▒ ░░ ▒░ ░ ▒ ░    ░░ ▒░ ░▒▒   ▓▒█░ ▒ ░░    ░▀▀▒                       │")
	fmt.Println("|                           ░ ▒  ▒  ░ ░  ░ ░       ░ ░  ░ ▒   ▒▒ ░   ░     ░  ░                       │")
	fmt.Println("|                           ░ ░  ░    ░    ░ ░       ░    ░   ▒    ░          ░                       │")
	fmt.Println("|                             ░       ░  ░           ░  ░     ░  ░         ░                          │")
	fmt.Println("|                                                                                                     |")
}
