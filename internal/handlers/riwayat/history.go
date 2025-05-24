package riwayat

import (
	"fmt"
	"crypto-simulator/internal/models"
	"crypto-simulator/internal/utils"
)

func RiwayatMenu() {
	utils.ClearScreen()
	fmt.Println("=== RIWAYAT TRANSAKSI ===")

	if models.JumlahTransaksi == 0 {
		fmt.Println("\nBelum ada transaksi!")
		utils.TungguUser()
		return
	}

	fmt.Println("\nDAFTAR TRANSAKSI:")
	fmt.Println("====================================================================")
	fmt.Printf("%-4s %-8s %-12s %-10s %-12s %-12s %-10s\n", "ID", "Jenis", "Nama", "Jumlah", "Harga", "Total", "Waktu")
	fmt.Println("====================================================================")

	for i := 0; i < models.JumlahTransaksi; i++ {
		fmt.Printf("%-4d %-8s %-12s %-10.4f $%-12.2f $%-12.2f %-10s\n",
			models.TransaksiList[i].ID,
			models.TransaksiList[i].Jenis,
			models.TransaksiList[i].Nama,
			models.TransaksiList[i].Jumlah,
			models.TransaksiList[i].Harga,
			models.TransaksiList[i].Total,
			models.TransaksiList[i].Waktu)
	}
	fmt.Println("====================================================================")
	utils.TungguUser()
}
