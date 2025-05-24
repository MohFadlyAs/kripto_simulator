package crypto

import (
	"fmt"
	"crypto-simulator/internal/models"
)

func TampilkanDetailCrypto(index int) {
	fmt.Println("\nDETAIL CRYPTO:")
	fmt.Println("====================================")
	fmt.Printf("Nama: %s\n", models.CryptoList[index].Nama)
	fmt.Printf("Simbol: %s\n", models.CryptoList[index].Simbol)
	fmt.Printf("Harga: $%.2f\n", models.CryptoList[index].Harga)
	fmt.Printf("Kapitalisasi Pasar: $%.2f\n", models.CryptoList[index].Kapitalisasi)
	fmt.Printf("Jumlah Dimiliki: %.4f\n", models.CryptoList[index].Jumlah)
	fmt.Println("====================================")
}

func TampilkanDaftarCypto(onlyOwned bool) {
	fmt.Println("\nDAFTAR CRYPTO:")
	fmt.Println("================================================")
	fmt.Printf("%-4s %-10s %-6s %-12s %-15s %-10s\n", "No", "Nama", "Simbol", "Harga", "Kapitalisasi", "Dimiliki")
	fmt.Println("================================================")

	for i := 0; i < models.JumlahCrypto; i++ {
		if !onlyOwned || models.CryptoList[i].Jumlah > 0 {
			fmt.Printf("%-4d %-10s %-6s $%-12.2f $%-15.2f %-10.4f\n",
				i+1,
				models.CryptoList[i].Nama,
				models.CryptoList[i].Simbol,
				models.CryptoList[i].Harga,
				models.CryptoList[i].Kapitalisasi,
				models.CryptoList[i].Jumlah)
		}
	}
	fmt.Println("================================================")
}
