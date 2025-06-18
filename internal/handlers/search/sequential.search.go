package search

import (
	"crypto-simulator/internal/handlers/crypto"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
	"fmt"
)

func SequentialSearch() {
	fmt.Print("Masukkan nama crypto: ")
	var nama string
	fmt.Scanln(&nama)

	found := false
	for i := 0; i < models.JumlahCrypto; i++ {
		if models.CryptoList[i].Nama == nama {
			crypto.TampilkanDetailCrypto(i)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan!")
	}
	utils.TungguUser()
}
