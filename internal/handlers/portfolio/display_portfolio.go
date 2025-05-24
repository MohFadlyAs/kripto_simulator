package portfolio

import (
	"fmt"
	"crypto-simulator/internal/models"
)

func tampilkanPortofolio() {
	fmt.Println("CRYPTO YANG DIMILIKI:")
	fmt.Println("==========================================================")
	fmt.Printf("%-5s %-12s %-10s %-15s %-15s\n", "No", "Crypto", "Jumlah", "Harga Satuan", "Total Nilai")
	fmt.Println("==========================================================")

	no := 1
	for i := 0; i < models.JumlahCrypto; i++ {
		if models.CryptoList[i].Jumlah > 0 {
			total := models.CryptoList[i].Jumlah * models.CryptoList[i].Harga
			fmt.Printf("%-5d %-12s %-10.4f $%-15.2f $%-15.2f\n",
				no,
				models.CryptoList[i].Nama,
				models.CryptoList[i].Jumlah,
				models.CryptoList[i].Harga,
				total)
			no++
		}
	}

	if no == 1 {
		fmt.Println(" Anda belum memiliki aset crypto")
	}
	fmt.Println("==========================================================")
}
