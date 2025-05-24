package crypto

import (
	"fmt"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
)

func editCryptoMenu() {
	utils.ClearScreen()
	TampilkanDaftarCypto(false)

	if models.JumlahCrypto == 0 {
		utils.TungguUser()
		return
	}

	fmt.Print("\nPilih nomor Crypto yang akan diedit: ")
	var index int
	fmt.Scanln(&index)
	index--

	if index < 0 || index >= models.JumlahCrypto {
		fmt.Println("Nomor tidak valid!")
		utils.TungguUser()
		return
	}

	var nama, simbol string
	var harga, kapitalisasi float64

	fmt.Printf("Nama (%s): ", models.CryptoList[index].Nama)
	fmt.Scanln(&nama)
	fmt.Printf("Simbol (%s): ", models.CryptoList[index].Simbol)
	fmt.Scanln(&simbol)
	fmt.Printf("Harga (%.2f): ", models.CryptoList[index].Harga)
	fmt.Scanln(&harga)
	fmt.Printf("Kapitalisasi (%.2f): ", models.CryptoList[index].Kapitalisasi)
	fmt.Scanln(&kapitalisasi)

	if nama != "" {
		models.CryptoList[index].Nama = nama
	}
	if simbol != "" {
		models.CryptoList[index].Simbol = simbol
	}
	if harga > 0 {
		models.CryptoList[index].Harga = harga
	}
	if kapitalisasi > 0 {
		models.CryptoList[index].Kapitalisasi = kapitalisasi
	}

	fmt.Println("Data berhasil diupdate!")
	utils.TungguUser()
}
