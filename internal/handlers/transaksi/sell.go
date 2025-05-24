package transaksi

import (
	"fmt"
	"crypto-simulator/internal/handlers/crypto"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
)

func JualCrypto() {
	utils.ClearScreen()
	fmt.Println("=== JUAL CRYPTO ===")
	crypto.TampilkanDaftarCypto(true)

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

	if models.CryptoList[index].Jumlah <= 0 {
		fmt.Println("Anda tidak memiliki crypto ini!")
		utils.TungguUser()
		return
	}

	fmt.Printf("Anda memiliki %.4f %s\n", models.CryptoList[index].Jumlah, models.CryptoList[index].Nama)
	fmt.Print("Jumlah yang dijual: ")
	var jumlah float64
	fmt.Scanln(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Jumlah harus positif!")
		utils.TungguUser()
		return
	}

	if jumlah > models.CryptoList[index].Jumlah {
		fmt.Println("Jumlah melebihi kepemilikan!")
		utils.TungguUser()
		return
	}

	total := jumlah * models.CryptoList[index].Harga
	models.Saldo += total
	models.CryptoList[index].Jumlah -= jumlah
	TambahTransaksi("jual", models.CryptoList[index].Nama, jumlah, models.CryptoList[index].Harga, total)

	fmt.Printf("Berhasil menjual %.2f %s seharga $%.2f\n", jumlah, models.CryptoList[index].Nama, total)
	utils.TungguUser()
}
