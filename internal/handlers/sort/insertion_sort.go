package sort

import (
	"fmt"
	"crypto-simulator/internal/models"
)

func InsertionSortByName(ascending bool) {
	for i := 1; i < models.JumlahCrypto; i++ {
		key := models.CryptoList[i]
		j := i - 1

		if ascending {
			for j >= 0 && models.CryptoList[j].Nama > key.Nama {
				models.CryptoList[j+1] = models.CryptoList[j]
				j--
			}
		} else {
			for j >= 0 && models.CryptoList[j].Nama < key.Nama {
				models.CryptoList[j+1] = models.CryptoList[j]
				j--
			}
		}
		models.CryptoList[j+1] = key
	}

	if ascending {
		fmt.Println("Data telah diurutkan berdasarkan nama (ascending)!")
	} else {
		fmt.Println("Data telah diurutkan berdasarkan nama (descending)!")
	}
}
