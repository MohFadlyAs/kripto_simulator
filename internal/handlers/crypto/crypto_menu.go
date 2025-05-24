package crypto

import (
	"fmt"
	"crypto-simulator/internal/utils"
)

func KelolaCryptoMenu() {
	for {
		utils.ClearScreen()
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
			tambahCryptoMenu()
		case 2:
			editCryptoMenu()
		case 3:
			hapusCryptoMenu()
		case 4:
			TampilkanDaftarCypto(false)
			utils.TungguUser()
		case 0:
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}
