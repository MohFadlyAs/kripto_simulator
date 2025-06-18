package portfolio

import (
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
	"fmt"
)

func PortofolioMenu() {
	fmt.Println("=== PORTOFOLIO CRYPTO ===")
	fmt.Printf("Saldo Virtual: $%.2f\n\n", models.Saldo)

	totalNilai := hitungTotalPortofolio()
	fmt.Printf("Total Nilai Portofolio: $%.2f\n", totalNilai)
	fmt.Printf("Total Aset (Saldo + Portofolio): $%.2f\n\n", models.Saldo+totalNilai)

	tampilkanPortofolio()
	utils.TungguUser()
}
