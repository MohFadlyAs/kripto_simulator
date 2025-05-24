package crypto

import (
	"fmt"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
)

func hapusCryptoMenu() {
	utils.ClearScreen()
	TampilkanDaftarCypto(false)

	if models.JumlahCrypto == 0 {
		utils.TungguUser()
		return
	}

	fmt.Print("\nPilih nomor crypto yang akan dihapus: ")
	var index int
	fmt.Scanln(&index)
	index--

	if index < 0 || index >= models.JumlahCrypto {
		fmt.Println("Nomor tidak valid!")
		utils.TungguUser()
		return
	}

	for i := index; i < models.JumlahCrypto-1; i++ {
		models.CryptoList[i] = models.CryptoList[i+1]
	}
	models.JumlahCrypto--

	fmt.Println("Crypto berhasil dihapus!")
	utils.TungguUser()
}
