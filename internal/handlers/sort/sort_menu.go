package sort

import (
	"crypto-simulator/internal/handlers/crypto"
	"crypto-simulator/internal/utils"
	"fmt"
)

func UrutkanMenu() {
	for {
		fmt.Println("=== URUTKAN CRYPTO ===")
		fmt.Println("1. Selection Sort (Harga)")
		fmt.Println("2. Insertion Sort (Nama - Ascending)")
		fmt.Println("3. Insertion Sort (Nama - Descending)")
		fmt.Println("4. Selection Sort (Kapitalisasi)")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println()
			SelectionSortByHarga()
			crypto.TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 2:
			fmt.Println()
			InsertionSortByName(true)
			crypto.TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 3:
			fmt.Println()
			InsertionSortByName(false)
			crypto.TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 4:
			fmt.Println()
			SelectionSortByMarketCap()
			crypto.TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 0:
			fmt.Print("\n\n\n")
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}
