package search

import (
	"crypto-simulator/internal/utils"
	"fmt"
)

func CariMenu() {
	for {
		fmt.Println("=== CARI CRYPTO ===")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println()
			SequentialSearch()
		case 2:
			fmt.Println()
			BinarySearch()
		case 0:
			fmt.Print("\n\n\n")
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}
