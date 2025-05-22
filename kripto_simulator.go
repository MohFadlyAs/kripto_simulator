package main

import "fmt"

// ======================== STRUCT & VARIABEL GLOBAL ========================
type Kripto struct {
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
    kriptoList      [100]Kripto
    transaksiList   [200]Transaksi
    jumlahKripto    int
    jumlahTransaksi int
    saldo           float64 = 10000 // Saldo awal
)

// ======================== MAIN & MENU UTAMA ========================
func main() {
    inisialisasiData()
    
    for {
        clearScreen()
        fmt.Println("=== APLIKASI SIMULASI KRIPTO ===")
        fmt.Printf("Saldo: $%.2f\n\n", saldo)
        fmt.Println("1. Kelola Kripto")
        fmt.Println("2. Transaksi")
        fmt.Println("3. Portofolio")
		fmt.Println("4. Riwayat Transaksi")
        fmt.Println("5. Cari Kripto")
        fmt.Println("6. Urutkan Kripto")
        fmt.Println("0. Keluar")
        fmt.Print("Pilih menu: ")
        
        var pilihan int
        fmt.Scanln(&pilihan)
        
        switch pilihan {
        case 1: menuKelolaKripto()
        case 2: menuTransaksi()
		case 3: menuPortofolio()
        case 4: menuRiwayat()
        case 5: menuCariKripto()
        case 6: menuUrutkanKripto()
        case 0: 
            fmt.Println("Terima kasih!")
            return
        default:
            fmt.Println("Input salah!")
            tungguUser()
        }
    }
}

func inisialisasiData() {
    // Data dummy
    tambahKripto("Bitcoin", "BTC", 50000, 950000000000)
    tambahKripto("Ethereum", "ETH", 3000, 360000000000)
    tambahKripto("Cardano", "ADA", 1.5, 50000000000)
}

// ======================== FUNGSI UTILITAS ========================
func clearScreen() {
    fmt.Print("\033[H\033[2J")
}

func tungguUser() {
    fmt.Print("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
}

// ======================== MODUL KELOLA KRIPTO ========================
func menuKelolaKripto() {
    for {
        clearScreen()
        fmt.Println("=== KELOLA KRIPTO ===")
        fmt.Println("1. Tambah Kripto")
        fmt.Println("2. Edit Kripto")
        fmt.Println("3. Hapus Kripto")
        fmt.Println("4. Lihat Daftar")
        fmt.Println("0. Kembali")
        fmt.Print("Pilih: ")
        
        var pilihan int
        fmt.Scanln(&pilihan)
        
        switch pilihan {
        case 1: tambahKriptoMenu()
        case 2: editKriptoMenu()
        case 3: hapusKriptoMenu()
        case 4: 
            tampilkanDaftarKripto(false)
            tungguUser()
        case 0: return
        default: 
            fmt.Println("Input salah!")
            tungguUser()
        }
    }
}

func tambahKriptoMenu() {
    clearScreen()
    var nama, simbol string
    var harga, kapitalisasi float64
    
    fmt.Println("=== TAMBAH KRIPTO ===")
    fmt.Print("Nama: ")
    fmt.Scanln(&nama)
    fmt.Print("Simbol: ")
    fmt.Scanln(&simbol)
    fmt.Print("Harga: ")
    fmt.Scanln(&harga)
    fmt.Print("Kapitalisasi: ")
    fmt.Scanln(&kapitalisasi)
    
    tambahKripto(nama, simbol, harga, kapitalisasi)
    fmt.Println("Berhasil ditambahkan!")
    tungguUser()
}

func tambahKripto(nama, simbol string, harga, kapitalisasi float64) {
    if jumlahKripto >= len(kriptoList) {
        fmt.Println("Error: Daftar penuh!")
        return
    }
    kriptoList[jumlahKripto] = Kripto{
        Nama: nama, 
        Simbol: simbol, 
        Harga: harga, 
        Kapitalisasi: kapitalisasi,
    }
    jumlahKripto++
}

func editKriptoMenu() {
    clearScreen()
    tampilkanDaftarKripto(false)
    
    if jumlahKripto == 0 {
        tungguUser()
        return
    }
    
    fmt.Print("\nPilih nomor kripto yang akan diedit: ")
    var index int
    fmt.Scanln(&index)
    index--
    
    if index < 0 || index >= jumlahKripto {
        fmt.Println("Nomor tidak valid!")
        tungguUser()
        return
    }
    
    var nama, simbol string
    var harga, kapitalisasi float64
    
    fmt.Printf("Nama (%s): ", kriptoList[index].Nama)
    fmt.Scanln(&nama)
    fmt.Printf("Simbol (%s): ", kriptoList[index].Simbol)
    fmt.Scanln(&simbol)
    fmt.Printf("Harga (%.2f): ", kriptoList[index].Harga)
    fmt.Scanln(&harga)
    fmt.Printf("Kapitalisasi (%.2f): ", kriptoList[index].Kapitalisasi)
    fmt.Scanln(&kapitalisasi)
    
    // Update data
    if nama != "" {
        kriptoList[index].Nama = nama
    }
    if simbol != "" {
        kriptoList[index].Simbol = simbol
    }
    if harga > 0 {
        kriptoList[index].Harga = harga
    }
    if kapitalisasi > 0 {
        kriptoList[index].Kapitalisasi = kapitalisasi
    }
    
    fmt.Println("Data berhasil diupdate!")
    tungguUser()
}

func hapusKriptoMenu() {
    clearScreen()
    tampilkanDaftarKripto(false)
    
    if jumlahKripto == 0 {
        tungguUser()
        return
    }
    
    fmt.Print("\nPilih nomor kripto yang akan dihapus: ")
    var index int
    fmt.Scanln(&index)
    index--
    
    if index < 0 || index >= jumlahKripto {
        fmt.Println("Nomor tidak valid!")
        tungguUser()
        return
    }
    
    // Geser elemen array
    for i := index; i < jumlahKripto-1; i++ {
        kriptoList[i] = kriptoList[i+1]
    }
    jumlahKripto--
    
    fmt.Println("Kripto berhasil dihapus!")
    tungguUser()
}

// ======================== MODUL TRANSAKSI ========================
func menuTransaksi() {
    for {
        clearScreen()
        fmt.Println("=== TRANSAKSI ===")
        fmt.Printf("Saldo: $%.2f\n\n", saldo)
        fmt.Println("1. Beli")
        fmt.Println("2. Jual")
        fmt.Println("0. Kembali")
        fmt.Print("Pilih: ")
        
        var pilihan int
        fmt.Scanln(&pilihan)
        
        switch pilihan {
        case 1: beliKripto()
        case 2: jualKripto()
        case 0: return
        default:
            fmt.Println("Input salah!")
            tungguUser()
        }
    }
}

func beliKripto() {
    clearScreen()
    fmt.Println("=== BELI KRIPTO ===")
    tampilkanDaftarKripto(false)
    
    if jumlahKripto == 0 {
        tungguUser()
        return
    }
    
    fmt.Print("\nPilih nomor kripto: ")
    var index int
    fmt.Scanln(&index)
    index--
    
    if index < 0 || index >= jumlahKripto {
        fmt.Println("Nomor tidak valid!")
        tungguUser()
        return
    }
    
    fmt.Printf("Harga %s: $%.2f\n", kriptoList[index].Nama, kriptoList[index].Harga)
    fmt.Print("Jumlah yang dibeli: ")
    var jumlah float64
    fmt.Scanln(&jumlah)
    
    if jumlah <= 0 {
        fmt.Println("Jumlah harus positif!")
        tungguUser()
        return
    }
    
    total := jumlah * kriptoList[index].Harga
    if total > saldo {
        fmt.Println("Saldo tidak cukup!")
        tungguUser()
        return
    }
    
    saldo -= total
    kriptoList[index].Jumlah += jumlah
    tambahTransaksi("beli", kriptoList[index].Nama, jumlah, kriptoList[index].Harga, total)
    
    fmt.Printf("Berhasil membeli %.2f %s seharga $%.2f\n", jumlah, kriptoList[index].Nama, total)
    tungguUser()
}

func jualKripto() {
    clearScreen()
    fmt.Println("=== JUAL KRIPTO ===")
    tampilkanDaftarKripto(true) // Hanya tampilkan yang dimiliki
    
    if jumlahKripto == 0 {
        tungguUser()
        return
    }
    
    fmt.Print("\nPilih nomor kripto: ")
    var index int
    fmt.Scanln(&index)
    index--
    
    if index < 0 || index >= jumlahKripto {
        fmt.Println("Nomor tidak valid!")
        tungguUser()
        return
    }
    
    if kriptoList[index].Jumlah <= 0 {
        fmt.Println("Anda tidak memiliki kripto ini!")
        tungguUser()
        return
    }
    
    fmt.Printf("Anda memiliki %.4f %s\n", kriptoList[index].Jumlah, kriptoList[index].Nama)
    fmt.Print("Jumlah yang dijual: ")
    var jumlah float64
    fmt.Scanln(&jumlah)
    
    if jumlah <= 0 {
        fmt.Println("Jumlah harus positif!")
        tungguUser()
        return
    }
    
    if jumlah > kriptoList[index].Jumlah {
        fmt.Println("Jumlah melebihi kepemilikan!")
        tungguUser()
        return
    }
    
    total := jumlah * kriptoList[index].Harga
    saldo += total
    kriptoList[index].Jumlah -= jumlah
    tambahTransaksi("jual", kriptoList[index].Nama, jumlah, kriptoList[index].Harga, total)
    
    fmt.Printf("Berhasil menjual %.2f %s seharga $%.2f\n", jumlah, kriptoList[index].Nama, total)
    tungguUser()
}

func tambahTransaksi(jenis, nama string, jumlah, harga, total float64) {
    if jumlahTransaksi >= len(transaksiList) {
        // Geser transaksi terlama (FIFO)
        for i := 0; i < len(transaksiList)-1; i++ {
            transaksiList[i] = transaksiList[i+1]
        }
        jumlahTransaksi = len(transaksiList) - 1
    }
    
    transaksiList[jumlahTransaksi] = Transaksi{
        ID:       jumlahTransaksi + 1,
        Jenis:    jenis,
        Nama:     nama,
        Jumlah:   jumlah,
        Harga:    harga,
        Total:    total,
        Waktu:    "2023-11-15", // Format sederhana
    }
    jumlahTransaksi++
}

// ======================== MODUL PORTOFOLIO ========================
func menuPortofolio() {
    clearScreen()
    fmt.Println("=== PORTOFOLIO KRIPTO ===")
    fmt.Printf("Saldo Virtual: $%.2f\n\n", saldo)
    
    // Hitung total nilai portofolio
    totalNilaiPortofolio := hitungTotalPortofolio()
    fmt.Printf("Total Nilai Portofolio: $%.2f\n", totalNilaiPortofolio)
    fmt.Printf("Total Aset (Saldo + Portofolio): $%.2f\n\n", saldo + totalNilaiPortofolio)
    
    // Tampilkan daftar kripto yang dimiliki
    tampilkanPortofolio()
    tungguUser()
}

func hitungTotalPortofolio() float64 {
    total := 0.0
    for i := 0; i < jumlahKripto; i++ {
        if kriptoList[i].Jumlah > 0 {
            total += kriptoList[i].Jumlah * kriptoList[i].Harga
        }
    }
    return total
}

func tampilkanPortofolio() {
    fmt.Println("KRIPTO YANG DIMILIKI:")
    fmt.Println("==========================================================")
    fmt.Printf("%-5s %-12s %-10s %-15s %-15s\n", "No", "Kripto", "Jumlah", "Harga Satuan", "Total Nilai")
    fmt.Println("==========================================================")
    
    no := 1
    for i := 0; i < jumlahKripto; i++ {
        if kriptoList[i].Jumlah > 0 {
            total := kriptoList[i].Jumlah * kriptoList[i].Harga
            fmt.Printf("%-5d %-12s %-10.4f $%-15.2f $%-15.2f\n", 
                no,
                kriptoList[i].Nama,
                kriptoList[i].Jumlah,
                kriptoList[i].Harga,
                total)
            no++
        }
    }
    
    if no == 1 {
        fmt.Println(" Anda belum memiliki aset kripto")
    }
    fmt.Println("==========================================================")
}

// ======================== MODUL RIWAYAT ========================
func menuRiwayat() {
    clearScreen()
    fmt.Println("=== RIWAYAT TRANSAKSI ===")
    
    if jumlahTransaksi == 0 {
        fmt.Println("\nBelum ada transaksi!")
        tungguUser()
        return
    }
    
    fmt.Println("\nDAFTAR TRANSAKSI:")
    fmt.Println("====================================================================")
    fmt.Printf("%-4s %-8s %-12s %-10s %-12s %-12s %-10s\n", "ID", "Jenis", "Nama", "Jumlah", "Harga", "Total", "Waktu")
    fmt.Println("====================================================================")
    
    for i := 0; i < jumlahTransaksi; i++ {
        fmt.Printf("%-4d %-8s %-12s %-10.4f $%-12.2f $%-12.2f %-10s\n", 
            transaksiList[i].ID, 
            transaksiList[i].Jenis, 
            transaksiList[i].Nama, 
            transaksiList[i].Jumlah, 
            transaksiList[i].Harga, 
            transaksiList[i].Total,
            transaksiList[i].Waktu)
    }
    fmt.Println("====================================================================")
    tungguUser()
}

// ======================== MODUL PENCARIAN ========================
func menuCariKripto() {
    for {
        clearScreen()
        fmt.Println("=== CARI KRIPTO ===")
        fmt.Println("1. Sequential Search")
        fmt.Println("2. Binary Search")
        fmt.Println("0. Kembali")
        fmt.Print("Pilih: ")
        
        var pilihan int
        fmt.Scanln(&pilihan)
        
        switch pilihan {
        case 1: sequentialSearch()
        case 2: binarySearch()
        case 0: return
        default:
            fmt.Println("Input salah!")
            tungguUser()
        }
    }
}

func sequentialSearch() {
    clearScreen()
    fmt.Print("Masukkan nama kripto: ")
    var nama string
    fmt.Scanln(&nama)
    
    found := false
    for i := 0; i < jumlahKripto; i++ {
        if kriptoList[i].Nama == nama {
            tampilkanDetailKripto(i)
            found = true
        }
    }
    
    if !found {
        fmt.Println("Tidak ditemukan!")
    }
    tungguUser()
}

func binarySearch() {
    clearScreen()
    insertionSortByName(true) // Urutkan dulu
    
    fmt.Print("Masukkan nama kripto: ")
    var nama string
    fmt.Scanln(&nama)
    
    low := 0
    high := jumlahKripto - 1
    found := -1
    
    for low <= high {
        mid := (low + high) / 2
        if kriptoList[mid].Nama == nama {
            found = mid
            break
        } else if kriptoList[mid].Nama < nama {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    
    if found != -1 {
        tampilkanDetailKripto(found)
    } else {
        fmt.Println("Tidak ditemukan!")
    }
    tungguUser()
}

func tampilkanDetailKripto(index int) {
    fmt.Println("\nDETAIL KRIPTO:")
    fmt.Println("====================================")
    fmt.Printf("Nama: %s\n", kriptoList[index].Nama)
    fmt.Printf("Simbol: %s\n", kriptoList[index].Simbol)
    fmt.Printf("Harga: $%.2f\n", kriptoList[index].Harga)
    fmt.Printf("Kapitalisasi Pasar: $%.2f\n", kriptoList[index].Kapitalisasi)
    fmt.Printf("Jumlah Dimiliki: %.4f\n", kriptoList[index].Jumlah)
    fmt.Println("====================================")
}

// ======================== MODUL PENGURUTAN ========================
func menuUrutkanKripto() {
    for {
        clearScreen()
        fmt.Println("=== URUTKAN KRIPTO ===")
        fmt.Println("1. Selection Sort (Harga)")
        fmt.Println("2. Insertion Sort (Nama - Ascending)")
        fmt.Println("3. Insertion Sort (Nama - Descending)")
        fmt.Println("4. Selection Sort (Kapitalisasi)")
        fmt.Println("0. Kembali")
        fmt.Print("Pilih: ")
        
        var pilihan int
        fmt.Scanln(&pilihan)
        
        switch pilihan {
        case 1: 
            selectionSortByHarga()
            tampilkanDaftarKripto(false)
            tungguUser()
        case 2: 
            insertionSortByName(true)
            tampilkanDaftarKripto(false)
            tungguUser()
        case 3: 
            insertionSortByName(false)
            tampilkanDaftarKripto(false)
            tungguUser()
        case 4: 
            selectionSortByMarketCap()
            tampilkanDaftarKripto(false)
            tungguUser()
        case 0: 
            return
        default:
            fmt.Println("Input salah!")
            tungguUser()
        }
    }
}

func selectionSortByHarga() {
    for i := 0; i < jumlahKripto-1; i++ {
        minIndex := i
        for j := i + 1; j < jumlahKripto; j++ {
            if kriptoList[j].Harga < kriptoList[minIndex].Harga {
                minIndex = j
            }
        }
        // Tukar posisi
        kriptoList[i], kriptoList[minIndex] = kriptoList[minIndex], kriptoList[i]
    }
    fmt.Println("Data telah diurutkan berdasarkan harga (ascending)!")
}

func insertionSortByName(ascending bool) {
    for i := 1; i < jumlahKripto; i++ {
        key := kriptoList[i]
        j := i - 1
        
        if ascending {
            for j >= 0 && kriptoList[j].Nama > key.Nama {
                kriptoList[j+1] = kriptoList[j]
                j--
            }
        } else {
            for j >= 0 && kriptoList[j].Nama < key.Nama {
                kriptoList[j+1] = kriptoList[j]
                j--
            }
        }
        kriptoList[j+1] = key
    }
    
    if ascending {
        fmt.Println("Data telah diurutkan berdasarkan nama (ascending)!")
    } else {
        fmt.Println("Data telah diurutkan berdasarkan nama (descending)!")
    }
}

func selectionSortByMarketCap() {
    for i := 0; i < jumlahKripto-1; i++ {
        maxIndex := i
        for j := i + 1; j < jumlahKripto; j++ {
            if kriptoList[j].Kapitalisasi > kriptoList[maxIndex].Kapitalisasi {
                maxIndex = j
            }
        }
        // Tukar posisi
        kriptoList[i], kriptoList[maxIndex] = kriptoList[maxIndex], kriptoList[i]
    }
    fmt.Println("Data telah diurutkan berdasarkan kapitalisasi pasar (descending)!")
}

// ======================== TAMPILAN DATA ========================
func tampilkanDaftarKripto(onlyOwned bool) {
    fmt.Println("\nDAFTAR KRIPTO:")
    fmt.Println("================================================")
    fmt.Printf("%-4s %-10s %-6s %-12s %-15s %-10s\n", "No", "Nama", "Simbol", "Harga", "Kapitalisasi", "Dimiliki")
    fmt.Println("================================================")
    
    for i := 0; i < jumlahKripto; i++ {
        if !onlyOwned || kriptoList[i].Jumlah > 0 {
            fmt.Printf("%-4d %-10s %-6s $%-12.2f $%-15.2f %-10.4f\n", 
                i+1, 
                kriptoList[i].Nama, 
                kriptoList[i].Simbol, 
                kriptoList[i].Harga, 
                kriptoList[i].Kapitalisasi,
                kriptoList[i].Jumlah)
        }
    }
    fmt.Println("================================================")
}