package crypto

import (
	"fmt"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
)

func tambahCryptoMenu() {
	utils.ClearScreen()
	var nama, simbol string
	var harga, kapitalisasi float64

	fmt.Println("=== TAMBAH CRYPTO ===")
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Simbol: ")
	fmt.Scanln(&simbol)
	fmt.Print("Harga: ")
	fmt.Scanln(&harga)
	fmt.Print("Kapitalisasi: ")
	fmt.Scanln(&kapitalisasi)

	TambahCrypto(nama, simbol, harga, kapitalisasi)
	fmt.Println("Berhasil ditambahkan!")
	utils.TungguUser()
}

func TambahCrypto(nama, simbol string, harga, kapitalisasi float64) {
	if models.JumlahCrypto >= len(models.CryptoList) {
		fmt.Println("Error: Daftar penuh!")
		return
	}
	models.CryptoList[models.JumlahCrypto] = models.Crypto{
		Nama:         nama,
		Simbol:       simbol,
		Harga:        harga,
		Kapitalisasi: kapitalisasi,
	}
	models.JumlahCrypto++
}
