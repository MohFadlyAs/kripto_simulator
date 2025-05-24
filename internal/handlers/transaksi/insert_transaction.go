package transaksi

import (
	"crypto-simulator/internal/models"
	"time"
)

func TambahTransaksi(jenis, nama string, jumlah, harga, total float64) {
	if models.JumlahTransaksi >= len(models.TransaksiList) {
		for i := 0; i < len(models.TransaksiList)-1; i++ {
			models.TransaksiList[i] = models.TransaksiList[i+1]
		}
		models.JumlahTransaksi = len(models.TransaksiList) - 1
	}

	models.TransaksiList[models.JumlahTransaksi] = models.Transaksi{
		ID:     models.JumlahTransaksi + 1,
		Jenis:  jenis,
		Nama:   nama,
		Jumlah: jumlah,
		Harga:  harga,
		Total:  total,
		Waktu:  time.Now().Format("2006-01-02 15:04:05"),
	}
	models.JumlahTransaksi++
}
