package sort

import (
	"fmt"
	"crypto-simulator/internal/handlers/crypto"
	"crypto-simulator/internal/utils"
)

func UrutkanMenu() {
	for {
		utils.ClearScreen()
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
			SelectionSortByHarga()
			crypto.TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 2:
			InsertionSortByName(true)
			crypto.TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 3:
			InsertionSortByName(false)
			crypto.TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 4:
			SelectionSortByMarketCap()
			crypto.TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 0:
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}
