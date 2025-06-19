package transaksi

import (
	"crypto-simulator/internal/handlers/crypto"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
	"fmt"
)

func BeliCrypto() {
	fmt.Println("=== BELI CRYPTO ===")
	crypto.TampilkanDaftarCypto(false)

	if models.JumlahCrypto == 0 {
		utils.TungguUser()
		return
	}

	fmt.Print("\nPilih nomor crypto: ")
	var index int
	fmt.Scanln(&index)
	index--

	if index < 0 || index >= models.JumlahCrypto {
		fmt.Println("Nomor tidak valid!")
		utils.TungguUser()
		return
	}

	fmt.Printf("Harga %s: $%.2f\n", models.CryptoList[index].Nama, models.CryptoList[index].Harga)
	fmt.Print("Jumlah yang dibeli: ")
	var jumlah float64
	fmt.Scanln(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Jumlah harus positif!")
		utils.TungguUser()
		return
	}

	total := jumlah * models.CryptoList[index].Harga
	if total > models.Saldo {
		fmt.Println("Saldo tidak cukup!")
		utils.TungguUser()
		return
	}

	models.Saldo -= total
	models.CryptoList[index].Jumlah += jumlah
	TambahTransaksi("beli", models.CryptoList[index].Nama, jumlah, models.CryptoList[index].Harga, total)

	fmt.Printf("Berhasil membeli %.2f %s seharga $%.2f\n", jumlah, models.CryptoList[index].Nama, total)
	utils.TungguUser()
}
