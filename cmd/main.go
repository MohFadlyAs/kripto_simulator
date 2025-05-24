package main

import (
	"fmt"
	"crypto-simulator/internal/handlers/crypto"
	"crypto-simulator/internal/handlers/portfolio"
	"crypto-simulator/internal/handlers/riwayat"
	"crypto-simulator/internal/handlers/search"
	"crypto-simulator/internal/handlers/sort"
	"crypto-simulator/internal/handlers/transaksi"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
)

func main() {
	crypto.InisialisasiData()

	for {
		utils.ClearScreen()
		fmt.Println("=== APLIKASI SIMULASI CRYPTO ===")
		fmt.Printf("Saldo: $%.2f\n\n", models.Saldo)
		fmt.Println("1. Kelola Crypto")
		fmt.Println("2. Transaksi")
		fmt.Println("3. Portofolio")
		fmt.Println("4. Riwayat Transaksi")
		fmt.Println("5. Cari Crypto")
		fmt.Println("6. Urutkan Crypto")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			crypto.KelolaCryptoMenu()
		case 2:
			transaksi.TransaksiMenu()
		case 3:
			portfolio.PortofolioMenu()
		case 4:
			riwayat.RiwayatMenu()
		case 5:
			search.CariMenu()
		case 6:
			sort.UrutkanMenu()
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}
