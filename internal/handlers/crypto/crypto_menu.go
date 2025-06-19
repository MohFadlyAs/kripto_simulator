package crypto

import (
	"crypto-simulator/internal/utils"
	"fmt"
)

func KelolaCryptoMenu() {
	for {
		fmt.Println("=== KELOLA CRYPTO ===")
		fmt.Println("1. Tambah Crypto")
		fmt.Println("2. Edit Crypto")
		fmt.Println("3. Hapus Crypto")
		fmt.Println("4. Lihat Daftar")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println()
			tambahCryptoMenu()
		case 2:
			fmt.Println()
			editCryptoMenu()
		case 3:
			fmt.Println()
			hapusCryptoMenu()
		case 4:
			fmt.Println()
			TampilkanDaftarCypto(false)
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
