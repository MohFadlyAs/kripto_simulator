package transaksi

import (
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
	"fmt"
)

func TransaksiMenu() {
	for {
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
			fmt.Println()
			BeliCrypto()
		case 2:
			fmt.Println()
			JualCrypto()
		case 0:
			fmt.Print("\n\n\n")
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}
