package portfolio

import (
	"crypto-simulator/internal/models"
)

func hitungTotalPortofolio() float64 {
	var total float64
	for i := 0; i < models.JumlahCrypto; i++ {
		if models.CryptoList[i].Jumlah > 0 {
			total += models.CryptoList[i].Jumlah * models.CryptoList[i].Harga
		}
	}
	return total
}
