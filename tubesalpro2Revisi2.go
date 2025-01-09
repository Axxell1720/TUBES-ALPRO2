package main
// Gusti Caesar Yuliawan
// Muhammad Harits
import (
	"fmt"
	"sort"
)

type Mobil struct {
	Nama        string
	TahunKeluar int
	Pabrikan    string
	Penjualan   int
}

type Pabrikan struct {
	Nama         string
	Mobil        [100]Mobil
	JumlahMobil  int
	TotalPenjualan int
}

var (
	pabrikanList [100]Pabrikan
	jumlahPabrikan int
)

func tambahPabrikan(nama string) {
	if jumlahPabrikan < len(pabrikanList) {
		pabrikanList[jumlahPabrikan] = Pabrikan{Nama: nama, JumlahMobil: 0}
		jumlahPabrikan++
		fmt.Println("Pabrikan berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas pabrikan penuh!")
	}
}

func tambahMobil(namaPabrikan, namaMobil string, tahun int, penjualan int) {
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == namaPabrikan {
			if pabrikanList[i].JumlahMobil < len(pabrikanList[i].Mobil) {
				pabrikanList[i].Mobil[pabrikanList[i].JumlahMobil] = Mobil{
					Nama:        namaMobil,
					TahunKeluar: tahun,
					Pabrikan:    namaPabrikan,
					Penjualan:   penjualan,
				}
				pabrikanList[i].JumlahMobil++
				pabrikanList[i].TotalPenjualan += penjualan
				fmt.Println("Mobil berhasil ditambahkan.")
				return
			} else {
				fmt.Println("Kapasitas mobil penuh untuk pabrikan ini!")
				return
			}
		}
	}
	fmt.Println("Pabrikan tidak ditemukan!")
}

func isiBanyakStokMobil(namaPabrikan string, mobilList []Mobil) {
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == namaPabrikan {
			for _, mobil := range mobilList {
				tambahMobil(namaPabrikan, mobil.Nama, mobil.TahunKeluar, mobil.Penjualan)
			}
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan!")
}

func daftarPabrikanTerurut() {
	sort.Slice(pabrikanList[:jumlahPabrikan], func(i, j int) bool {
		return pabrikanList[i].JumlahMobil > pabrikanList[j].JumlahMobil
	})
	fmt.Println("Daftar Pabrikan Terurut Berdasarkan Jumlah Mobil:")
	for _, p := range pabrikanList[:jumlahPabrikan] {
		fmt.Printf("%s: %d mobil\n", p.Nama, p.JumlahMobil)
	}
}

func tigaMobilTerlaris() {
	var semuaMobil []Mobil
	for i := 0; i < jumlahPabrikan; i++ {
		semuaMobil = append(semuaMobil, pabrikanList[i].Mobil[:pabrikanList[i].JumlahMobil]...)
	}
	sort.Slice(semuaMobil, func(i, j int) bool {
		return semuaMobil[i].Penjualan > semuaMobil[j].Penjualan
	})
	fmt.Println("3 Mobil dengan Penjualan Tertinggi:")
	for i := 0; i < 3 && i < len(semuaMobil); i++ {
		fmt.Printf("%d. %s oleh %s (%d penjualan)\n", i+1, semuaMobil[i].Nama, semuaMobil[i].Pabrikan, semuaMobil[i].Penjualan)
	}
}

func editMobil(namaPabrikan, namaMobil string, mobilBaru Mobil) {
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == namaPabrikan {
			for j := 0; j < pabrikanList[i].JumlahMobil; j++ {
				if pabrikanList[i].Mobil[j].Nama == namaMobil {
					pabrikanList[i].Mobil[j] = mobilBaru
					fmt.Println("Mobil berhasil diubah.")
					return
				}
			}
		}
	}
	fmt.Println("Mobil tidak ditemukan!")
}

func cariMobilBerdasarkanPabrikan(namaPabrikan string) {
	var ditemukan bool
	for i := 0; i < jumlahPabrikan; i++ {
		if pabrikanList[i].Nama == namaPabrikan {
			ditemukan = true
			fmt.Printf("Mobil-mobil dari Pabrikan %s:\n", namaPabrikan)
			for _, mobil := range pabrikanList[i].Mobil[:pabrikanList[i].JumlahMobil] {
				fmt.Printf("Nama Mobil: %s, Tahun Keluar: %d, Penjualan: %d\n", mobil.Nama, mobil.TahunKeluar, mobil.Penjualan)
			}
		}
	}
	if !ditemukan {
		fmt.Println("Pabrikan tidak ditemukan!")
	}
}

func main() {
	for {
		fmt.Println("\nAplikasi Dealer Mobil")
		fmt.Println("1. Tambah Pabrikan")
		fmt.Println("2. Tambah Mobil")
		fmt.Println("3. Isi Banyak Stok Mobil")
		fmt.Println("4. Daftar Pabrikan Terurut")
		fmt.Println("5. 3 Mobil dengan Penjualan Tertinggi")
		fmt.Println("6. Edit Mobil")
		fmt.Println("7. Cari Mobil Berdasarkan Pabrikan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan Anda: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var nama string
			fmt.Print("Masukkan nama pabrikan: ")
			fmt.Scan(&nama)
			tambahPabrikan(nama)
		case 2:
			var namaPabrikan, namaMobil string
			var tahun, penjualan int
			fmt.Print("Masukkan nama pabrikan: ")
			fmt.Scan(&namaPabrikan)
			fmt.Print("Masukkan nama mobil: ")
			fmt.Scan(&namaMobil)
			fmt.Print("Masukkan tahun keluar: ")
			fmt.Scan(&tahun)
			fmt.Print("Masukkan jumlah penjualan: ")
			fmt.Scan(&penjualan)
			tambahMobil(namaPabrikan, namaMobil, tahun, penjualan)
		case 3:
			var namaPabrikan string
			fmt.Print("Masukkan nama pabrikan: ")
			fmt.Scan(&namaPabrikan)
			var n int
			fmt.Print("Masukkan jumlah mobil yang ingin ditambahkan: ")
			fmt.Scan(&n)
			mobilList := make([]Mobil, n)
			for i := 0; i < n; i++ {
				fmt.Printf("Mobil ke-%d:\n", i+1)
				fmt.Print("Nama: ")
				fmt.Scan(&mobilList[i].Nama)
				fmt.Print("Tahun Keluar: ")
				fmt.Scan(&mobilList[i].TahunKeluar)
				fmt.Print("Jumlah Penjualan: ")
				fmt.Scan(&mobilList[i].Penjualan)
				mobilList[i].Pabrikan = namaPabrikan
			}
			isiBanyakStokMobil(namaPabrikan, mobilList)
		case 4:
			daftarPabrikanTerurut()
		case 5:
			tigaMobilTerlaris()
		case 6:
			var namaPabrikan, namaMobil, namaBaru string
			var tahunBaru, penjualanBaru int
			fmt.Print("Masukkan nama pabrikan: ")
			fmt.Scan(&namaPabrikan)
			fmt.Print("Masukkan nama mobil yang akan diedit: ")
			fmt.Scan(&namaMobil)
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&namaBaru)
			fmt.Print("Masukkan tahun keluar baru: ")
			fmt.Scan(&tahunBaru)
			fmt.Print("Masukkan jumlah penjualan baru: ")
			fmt.Scan(&penjualanBaru)
			editMobil(namaPabrikan, namaMobil, Mobil{Nama: namaBaru, TahunKeluar: tahunBaru, Pabrikan: namaPabrikan, Penjualan: penjualanBaru})
		case 7:
			var namaPabrikan string
			fmt.Print("Masukkan nama pabrikan yang dicari: ")
			fmt.Scan(&namaPabrikan)
			cariMobilBerdasarkanPabrikan(namaPabrikan) 
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
