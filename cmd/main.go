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
			fmt.Println()
			crypto.KelolaCryptoMenu()
		case 2:
			fmt.Println()
			transaksi.TransaksiMenu()
		case 3:
			fmt.Println()
			portfolio.PortofolioMenu()
		case 4:
			fmt.Println()
			riwayat.RiwayatMenu()
		case 5:
			fmt.Println()
			search.CariMenu()
		case 6:
			fmt.Println()
			sort.UrutkanMenu()
		case 0:
			fmt.Println()
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Input salah!")
			utils.TungguUser()
		}
	}
}
