package models

type Crypto struct {
	Nama         string
	Simbol       string
	Harga        float64
	Kapitalisasi float64
	Jumlah       float64
}

type Transaksi struct {
	ID     int
	Jenis  string // "beli" atau "jual"
	Nama   string
	Jumlah float64
	Harga  float64
	Total  float64
	Waktu  string
}

var (
	CryptoList      [100]Crypto
	TransaksiList   [200]Transaksi
	JumlahCrypto    int
	JumlahTransaksi int
	Saldo           float64 = 10000 // Saldo awal
)
