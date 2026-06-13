package main

import "fmt"

const NMAX int = 100

type catatan struct {
	Matkul    string
	Isi       string
	Tanggal   string
	Kesulitan int
}

type tabCatatan [NMAX]catatan

func main() {
	var data tabCatatan
	var jumlah int
	var menu int

	for {
		fmt.Println("\n==== APLIKASI AI ASISTEN BELAJAR ====")
		fmt.Println("1. Tambah Catatan")
		fmt.Println("2. Ubah Catatan")
		fmt.Println("3. Hapus Catatan")
		fmt.Println("4. Buat Soal Latihan")
		fmt.Println("5. Cari Materi")
		fmt.Println("6. Urutkan Catatan")
		fmt.Println("7. Tampilkan Jadwal Belajar")
		fmt.Println("8. Exit")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		if menu == 1 {
			tambahCatatan(&data, &jumlah)
		} else if menu == 2 {
			ubahCatatan(&data, jumlah)
		} else if menu == 3 {
			hapusCatatan(&data, &jumlah)
		} else if menu == 4 {
			buatSoal(data, jumlah)
		} else if menu == 5 {
			cariMateri(&data, jumlah)
		} else if menu == 6 {
			urutkanCatatan(&data, jumlah)
		} else if menu == 7 {
			tampilkanJadwal(data, jumlah)
		} else if menu == 8 {
			fmt.Println("Program selesai.")
			break
		} else {
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func tambahCatatan(data *tabCatatan, jumlah *int) {
	var n int

	fmt.Println("\n==== TAMBAH CATATAN ====")
	fmt.Print("Banyak catatan yang ingin ditambahkan: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if *jumlah >= NMAX {
			fmt.Println("Data catatan sudah penuh.")
			return
		}

		fmt.Println("\nCatatan ke-", *jumlah+1)

		fmt.Print("Masukkan nama matkul: ")
		fmt.Scan(&data[*jumlah].Matkul)

		fmt.Print("Masukkan isi catatan: ")
		fmt.Scan(&data[*jumlah].Isi)

		fmt.Print("Masukkan tanggal belajar: ")
		fmt.Scan(&data[*jumlah].Tanggal)

		fmt.Print("Masukkan tingkat kesulitan 1-5: ")
		fmt.Scan(&data[*jumlah].Kesulitan)

		*jumlah = *jumlah + 1

		fmt.Println("Catatan berhasil ditambahkan.")
	}
}

//FUNGSI UBAH CATATAN MENGGUNAKAN SEQUENTIAL SEARCH

func ubahCatatan(data *tabCatatan, jumlah int) {
	var keyword string
	var index int

	if jumlah == 0 {
		fmt.Println("Belum ada catatan yang bisa diubah.")
		return
	}

	fmt.Println("\n==== UBAH CATATAN ====")
	fmt.Print("Masukkan nama matkul yang ingin diubah: ")
	fmt.Scan(&keyword)

	index = sequentialSearch(data, jumlah, keyword)

	if index != -1 {
		fmt.Println("\nCatatan ditemukan.")
		fmt.Println("Matkul lama:", data[index].Matkul)
		fmt.Println("Isi lama   :", data[index].Isi)

		fmt.Print("Masukkan matkul baru: ")
		fmt.Scan(&data[index].Matkul)

		fmt.Print("Masukkan isi catatan baru: ")
		fmt.Scan(&data[index].Isi)

		fmt.Print("Masukkan tanggal belajar baru: ")
		fmt.Scan(&data[index].Tanggal)

		fmt.Print("Masukkan tingkat kesulitan baru 1-5: ")
		fmt.Scan(&data[index].Kesulitan)

		fmt.Println("Catatan berhasil diubah.")
	} else {
		fmt.Println("Catatan tidak ditemukan.")
	}
}

//FUNGSI HAPUS CATATAN MENGGUNAKAN SEQUENTIAL SEARCH

func hapusCatatan(data *tabCatatan, jumlah *int) {
	var keyword string
	var index int

	if *jumlah == 0 {
		fmt.Println("Belum ada catatan yang bisa dihapus.")
		return
	}

	fmt.Println("\n==== HAPUS CATATAN ====")
	fmt.Print("Masukkan nama matkul yang ingin dihapus: ")
	fmt.Scan(&keyword)

	index = sequentialSearch(data, *jumlah, keyword)

	if index != -1 {
		for i := index; i < *jumlah-1; i++ {
			data[i] = data[i+1]
		}

		*jumlah = *jumlah - 1

		fmt.Println("Catatan berhasil dihapus.")
	} else {
		fmt.Println("Catatan tidak ditemukan.")
	}
}

//FUNGSI BUAT SOAL MENGGUNAKAN SEQUENTIAL SEARCH

func buatSoal(data tabCatatan, jumlah int) {
	var keyword string
	var index int

	if jumlah == 0 {
		fmt.Println("Belum ada catatan untuk dibuat soal.")
		return
	}

	fmt.Println("\n==== MEMBUAT SOAL LATIHAN ====")
	fmt.Print("Masukkan nama matkul yang ingin dibuat soal: ")
	fmt.Scan(&keyword)

	index = sequentialSearch(&data, jumlah, keyword)

	if index != -1 {
		fmt.Println("\n==== SOAL LATIHAN ====")
		fmt.Println("Matkul:", data[index].Matkul)
		fmt.Println("Materi:", data[index].Isi)
		fmt.Println()
		fmt.Println("1. Jelaskan pengertian dari", data[index].Matkul)
		fmt.Println("2. Sebutkan poin penting dari materi", data[index].Isi)
		fmt.Println("3. Berikan contoh penerapan dari", data[index].Matkul)
		fmt.Println("4. Mengapa materi", data[index].Matkul, "penting untuk dipelajari?")
	} else {
		fmt.Println("Matkul tidak ditemukan.")
	}
}

func sequentialSearch(data *tabCatatan, jumlah int, keyword string) int {
	for i := 0; i < jumlah; i++ {
		if data[i].Matkul == keyword {
			return i
		}
	}

	return -1
}

//FUNGSI CARI MATERI MENGGUNAKAN SELECTION SORT DAN BINARY SEARCH

func cariMateri(data *tabCatatan, jumlah int) {
	var keyword string
	var index int

	if jumlah == 0 {
		fmt.Println("Belum ada catatan yang bisa dicari.")
		return
	}

	fmt.Println("\n==== CARI MATERI ====")
	fmt.Print("Masukkan nama matkul yang dicari: ")
	fmt.Scan(&keyword)

	selectionSortMatkul(data, jumlah)

	index = binarySearch(data, jumlah, keyword)

	if index != -1 {
		fmt.Println("\n==== DATA DITEMUKAN ====")
		fmt.Println("Matkul    :", data[index].Matkul)
		fmt.Println("Isi       :", data[index].Isi)
		fmt.Println("Tanggal   :", data[index].Tanggal)
		fmt.Println("Kesulitan :", data[index].Kesulitan)
	} else {
		fmt.Println("Materi tidak ditemukan.")
	}
}

func selectionSortMatkul(data *tabCatatan, jumlah int) {
	var min int
	var temp catatan

	for i := 0; i < jumlah-1; i++ {
		min = i

		for j := i + 1; j < jumlah; j++ {
			if data[j].Matkul < data[min].Matkul {
				min = j
			}
		}

		temp = data[i]
		data[i] = data[min]
		data[min] = temp
	}
}

func binarySearch(data *tabCatatan, jumlah int, keyword string) int {
	var left, right, mid int

	left = 0
	right = jumlah - 1

	for left <= right {
		mid = (left + right) / 2

		if data[mid].Matkul == keyword {
			return mid
		} else if data[mid].Matkul < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

//FUNGSI MENGURUTKAN CATATAN MENGGUNAKAN SELECTION SORT(BASED ON TANGGAL)

func urutkanCatatan(data *tabCatatan, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada catatan untuk diurutkan.")
		return
	}

	fmt.Println("\n==== URUTKAN CATATAN ====")

	selectionSortTanggal(data, jumlah)

	fmt.Println("Catatan berhasil diurutkan berdasarkan tanggal.")

	for i := 0; i < jumlah; i++ {
		fmt.Println("\nCatatan ke-", i+1)
		fmt.Println("Matkul    :", data[i].Matkul)
		fmt.Println("Isi       :", data[i].Isi)
		fmt.Println("Tanggal   :", data[i].Tanggal)
		fmt.Println("Kesulitan :", data[i].Kesulitan)
	}
}

func selectionSortTanggal(data *tabCatatan, jumlah int) {
	var min int
	var temp catatan

	for i := 0; i < jumlah-1; i++ {
		min = i

		for j := i + 1; j < jumlah; j++ {
			if data[j].Tanggal < data[min].Tanggal {
				min = j
			}
		}

		temp = data[i]
		data[i] = data[min]
		data[min] = temp
	}
}

//FUNGSI TAMPIL JADWAL MENGGUNAKAN INSERTION SORT

func tampilkanJadwal(data tabCatatan, jumlah int) {
	var waktu int

	if jumlah == 0 {
		fmt.Println("Belum ada catatan untuk dibuat jadwal.")
		return
	}

	fmt.Println("\n==== JADWAL BELAJAR ====")
	fmt.Print("Masukkan waktu luang belajar dalam menit: ")
	fmt.Scan(&waktu)

	insertionSortKesulitan(&data, jumlah)

	fmt.Println("\nJadwal belajar yang disarankan:")

	if waktu < 30 {
		fmt.Println("Waktu singkat. Fokus belajar materi dengan kesulitan paling tinggi.")
		fmt.Println("Matkul:", data[jumlah-1].Matkul)
		fmt.Println("Materi:", data[jumlah-1].Isi)
	} else if waktu >= 30 && waktu < 60 {
		fmt.Println("30 menit: Pelajari materi utama")
		fmt.Println("10 menit: Baca ulang catatan")
		fmt.Println("10 menit: Jawab soal latihan")
		fmt.Println("Rekomendasi matkul:", data[jumlah-1].Matkul)
	} else {
		fmt.Println("30 menit: Membaca catatan")
		fmt.Println("30 menit: Merangkum materi")
		fmt.Println("30 menit: Mengerjakan soal latihan")
		fmt.Println("15 menit: Evaluasi hasil belajar")
		fmt.Println("Rekomendasi matkul:", data[jumlah-1].Matkul)
	}
}

func insertionSortKesulitan(data *tabCatatan, jumlah int) {
	var temp catatan
	var j int

	for i := 1; i < jumlah; i++ {
		temp = data[i]
		j = i - 1

		for j >= 0 && data[j].Kesulitan > temp.Kesulitan {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = temp
	}
}
