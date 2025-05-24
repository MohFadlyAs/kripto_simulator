package transaksi

import (
	"fmt"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
)

func TransaksiMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("=== TRANSAKSI ===")
		fmt.Printf("Saldo: $%.2f\n\n", models.Saldo)
		fmt.Println("1. Beli")
		fmt.Println("2. Jual")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			BeliCrypto()
		case 2:
			JualCrypto()
		case 0:
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}