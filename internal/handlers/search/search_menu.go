package search

import (
	"fmt"
	"crypto-simulator/internal/utils"
)

func CariMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("=== CARI CRYPTO ===")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			SequentialSearch()
		case 2:
			BinarySearch()
		case 0:
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}
