package search

import (
	"crypto-simulator/internal/handlers/crypto"
	"crypto-simulator/internal/handlers/sort"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
	"fmt"
)

func BinarySearch() {
	sort.InsertionSortByName(true)

	fmt.Print("Masukkan nama Crypto: ")
	var nama string
	fmt.Scanln(&nama)

	low := 0
	high := models.JumlahCrypto - 1
	found := -1

	for low <= high {
		mid := (low + high) / 2
		if models.CryptoList[mid].Nama == nama {
			found = mid
			break
		} else if models.CryptoList[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found != -1 {
		crypto.TampilkanDetailCrypto(found)
	} else {
		fmt.Println("Tidak ditemukan!")
	}
	utils.TungguUser()
}
