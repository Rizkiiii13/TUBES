package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 100

type catatan struct {
	Matkul    string
	Isi       string
	Tanggal   string
	Kesulitan int
}

var data [NMAX]catatan
var jumlah int
var reader = bufio.NewReader(os.Stdin)

func main() {
	var menu int

	for {
		fmt.Println("\n==== APLIKASI AI SISTEMBELAJAR ====")
		fmt.Println("1. TAMBAH CATATAN")
		fmt.Println("2. UBAH CATATAN")
		fmt.Println("3. HAPUS CATATAN")
		fmt.Println("4. MEMBUAT SOAL LATIHAN")
		fmt.Println("5. MENGATUR JADWAL BELAJAR")
		fmt.Println("6. CARI MATERI SEQUENTIAL SEARCH")
		fmt.Println("7. EXIT")
		fmt.Println("Pilih menu berikut : ")

		if menu == 1 {
			inputcatatan()
		} else if menu == 2 {
			ubahcatatan()
		} else if menu == 3 {
			hapuscatatan()
		} else if menu == 4 {
			buatsoal()
		} else if menu == 5 {
			aturjadwal()
		} else if menu == 6 {
			carimateri()
		} else if menu == 7 {
			fmt.Println("Program Selesai")
			break
		} else {
			fmt.Println("Menu tidak tersedia, Silahkan masukkan kembali")
		}
	}
}

func bacaString(pesan string) string {
	var input string

	fmt.Print(pesan)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}

func bacaInt(pesan string) int {
	var angka int

	fmt.Print(pesan)
	fmt.Scanln(&angka)

	return angka
}

func inputcatatan() {

	var banyak int

	fmt.Println("\n====Tambah Catatan====")
	fmt.Println("Banyaknya catatan yang ingin di input : ")
	fmt.Scan(&banyak)

	for i := 0; i < banyak; i++ {
		if jumlah >= NMAX {
			fmt.Println("Catatan Sudah Penuh")
			return
		}

		fmt.Println("\nCatatan ke-", jumlah+1)

		data[jumlah].Matkul = bacaString("Masukkan Nama Matkul: ")
		data[jumlah].Isi = bacaString("Masukkan Isi Catatan: ")
		data[jumlah].Tanggal = bacaString("Masukkan Tanggal Catatan: ")
		data[jumlah].Kesulitan = bacaInt("Masukkan Tingkat Kesulitan Catatan: ")

		jumlah++

		fmt.Println("Catatan berhasil ditambahkan.")
	}
}

func ubahcatatan() {
	var keyword string
	var index int

	if jumlah == 0 {
		fmt.Println("Tidak Ada Catatan Yang Bisa Diubah")
		return
	}

	fmt.Println("\n====UBAH CATATAN====")
	keyword = bacaString("Masukkan topik catatan yang ingin diubah: ")

	index = sequentialsearch(keyword)

	if index != -1 {
		fmt.Println("Catatan Lama Ditemukan")
		fmt.Println("Matkul : ", data[index].Matkul)
		fmt.Println("Isi Lama : ", data[index].Isi)

		data[index].Matkul = bacaString("Masukkan nama Matkul yang baru : ")
		data[index].Isi = bacaString("Masukkan Isi Catatan yang baru : ")
		data[index].Tanggal = bacaString("Masukkan tanggal Catatan baru: ")
		data[index].Kesulitan = bacaInt("Masukkan Tingkat Kesulitan baru 1 - 5 : ")

		fmt.Println("Catatan berhasil diubah.")
	} else {
		fmt.Println("Catatan tidak ditemukan")
	}

}

func hapuscatatan() {
	var keyword string
	var index int

	if jumlah == 0 {
		fmt.Println("Belum ada catatan yang bisa dihapus.")
		return
	}

	fmt.Println("\n=== HAPUS CATATAN ===")
	keyword = bacaString("Masukkan topik catatan yang ingin dihapus: ")

	index = sequentialsearch(keyword)

	if index != -1 {
		for i := index; i < jumlah-1; i++ {
			data[i] = data[i+1]
		}

		jumlah--

		fmt.Println("Catatan berhasil dihapus.")
	} else {
		fmt.Println("Catatan tidak ditemukan.")
	}
}

func sequentialsearch(keyword string) int {
	for i := 0; i < jumlah; i++ {
		if strings.ToLower(data[i].Matkul) == strings.ToLower(keyword) {
			return i
		}
	}

	return -1
}

func buatsoal() {
	var keyword string
	var index int

	if jumlah == 0 {
		fmt.Println("Belum ada catatan untuk dibuat soal.")
		return
	}

	fmt.Println("\n=== MEMBUAT SOAL LATIHAN ===")
	keyword = bacaString("Masukkan matkul yang ingin dibuat soal: ")

	index = sequentialsearch(keyword)

	if index != -1 {
		fmt.Println("\n=== SOAL LATIHAN ===")
		fmt.Println("Matkul :", data[index].Matkul)
		fmt.Println("Isi materi :", data[index].Isi)
		fmt.Println()
		fmt.Println("1. Jelaskan Pengertian dari", data[index].Matkul)
		fmt.Println("2. Sebutkan poin penting dari materi tersebut!")
		fmt.Println("3. Berikan contoh penerapan dari ", data[index].Matkul)
		fmt.Println("4. Mengapa Materi", data[index].Matkul, "Penting Untuk Di Pelajari")
	} else {
		fmt.Println("Matkul tidak ditemukan")
	}

}
