package sort

import (
	"fmt"
	"crypto-simulator/internal/models"
)

func SelectionSortByMarketCap() {
	for i := 0; i < models.JumlahCrypto-1; i++ {
		maxIndex := i
		for j := i + 1; j < models.JumlahCrypto; j++ {
			if models.CryptoList[j].Kapitalisasi > models.CryptoList[maxIndex].Kapitalisasi {
				maxIndex = j
			}
		}
		models.CryptoList[i], models.CryptoList[maxIndex] = models.CryptoList[maxIndex], models.CryptoList[i]
	}
	fmt.Println("Data telah diurutkan berdasarkan kapitalisasi pasar (descending)!")
}

func SelectionSortByHarga() {
	for i := 0; i < models.JumlahCrypto-1; i++ {
		minIndex := i
		for j := i + 1; j < models.JumlahCrypto; j++ {
			if models.CryptoList[j].Harga < models.CryptoList[minIndex].Harga {
				minIndex = j
			}
		}
		models.CryptoList[i], models.CryptoList[minIndex] = models.CryptoList[minIndex], models.CryptoList[i]
	}
	fmt.Println("Data telah diurutkan berdasarkan harga (ascending)!")
}
