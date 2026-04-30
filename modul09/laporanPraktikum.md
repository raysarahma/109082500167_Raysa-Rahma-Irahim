# <h1 align="center"> Laporan Praktikum Modul 9 </h1>
<p align="center">[Raysa Rahma Irahim] - [109082500167]</p>

## Unguided 

### 1. [Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".]
#### assesment01.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func didalam(cx, cy, r, x, y float64) bool {
	if jarak(cx, cy, x, y) <= r {
		return true
	}
	return false
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul09/output/assesment01.png) 
[Program ini untuk mengecek sebuah titik apakah berada di dalam atau di luar dua lingkaran. Program menerima input yang kita masukkan berupa titik pusat dan jari-jari dari dua lingkaran, serta satu titik yang mau dicek posisinya. Setelah itu, program bakal hitung jarak antara titik tersebut dengan pusat masing-masing lingkaran memakai fungsi jarak. Lalu jarak digunakan pada fungsi didalam untuk mengecek apakah titik tersebut masih di dalam lingkaran atau tidak dengan membandingkan jarak dan jari-jari lingkaran. Dari hasil itu, program bakal menampilkan apakah titik berada di dalam lingkaran 1, di dalam lingkaran 2, di dalam kedua lingkaran, atau di luar kedua lingkaran.]


### 2. [Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut: a. Menampilkan keseluruhan isi dari array. b. Menampilkan elemen-elemen array dengan indeks ganjil saja. c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap). d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna. e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil f. Menampilkan rata-rata dari bilangan yang ada di dalam array. g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut. h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.]
#### assesment02.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	var arr [100]int

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Isi array: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Indeks kelipatan x: ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)

	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Print("Array setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var sum int
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	rata := float64(sum) / float64(n)
	fmt.Println("Rata-rata:", rata)

	var total float64
	for i := 0; i < n; i++ {
		total += math.Pow(float64(arr[i])-rata, 2)
	}
	std := math.Sqrt(total / float64(n))
	fmt.Printf("Standar deviasi: %.2f\n", std)

	var cari, freq int
	fmt.Print("Masukkan bilangan yang dicari: ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul09/output/assesment02.png)
[Program ini dibuat untuk mengolah sekumpulan bilangan bulat yang disimpan dalam sebuah array. Pertama, kita memasukkan jumlah data, kemudian nilai-nilai tersebut dimasukkan ke dalam array. Program lalu menampilkan seluruh isi array, serta menampilkan elemen berdasarkan indeks ganjil dan genap (dengan asumsi indeks ke-0 adalah genap). Selanjutnya, kita memasukkan nilai x untuk menampilkan elemen pada indeks kelipatan x.

Program juga menyediakan fitur untuk menghapus elemen pada indeks tertentu dengan cara menggeser elemen setelahnya, lalu jumlah data dikurangi satu agar data yang dihapus tidak ikut ditampilkan. Setelah itu, program menghitung rata-rata dari seluruh elemen, serta menghitung standar deviasi dengan membandingkan setiap nilai terhadap rata-rata. Terakhir, kita memasukkan suatu bilangan untuk mengetahui frekuensi kemunculannya di dalam array. Hasil akhir yang ditampilkan mencakup isi array, pengelompokan indeks, hasil setelah penghapusan, rata-rata, standar deviasi, dan frekuensi bilangan tersebut.]


### 3. [Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.]
#### assesment03.go

```go
package main

import "fmt"

func main() {
	var A, B string
	fmt.Print("Klub A: ")
	fmt.Scan(&A)
	fmt.Print("Klub B: ")
	fmt.Scan(&B)

	var a, b int
	var hasil [100]string
	var n int = 0
	i := 1

	for {
		fmt.Print("Pertandingan ", i, ": ")
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			break
		}

		if a > b {
			hasil[n] = A
		} else if b > a {
			hasil[n] = B
		} else {
			hasil[n] = "Draw"
		}

		n++
		i++
	}

	for j := 0; j < n; j++ {
		fmt.Println("Hasil", j+1, ":", hasil[j])
	}
	fmt.Println("Pertandingan selesai")

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul09/output/assesment03.png)
[Program ini digunakan untuk menentukan hasil dari beberapa pertandingan antara dua klub. Pada awalnya, kita memasukkan nama dua klub yang akan bertanding. Selanjutnya, program meminta kita memasukkan skor dari setiap pertandingan secara berulang. Proses input akan berhenti jika salah satu skor yang dimasukkan bernilai negatif. Setiap hasil pertandingan kemudian disimpan ke dalam array, yaitu nama klub yang menang, atau “Draw” jika kedua skor sama.

Setelah semua data pertandingan dimasukkan, program akan menampilkan hasil setiap pertandingan secara berurutan berdasarkan data yang telah disimpan sebelumnya. Terakhir, program menampilkan pesan bahwa seluruh pertandingan telah selesai.]  

### 4. [Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama. *Palindrom adalah teks yang dibaca dari awal atau akhir adalah sama, contoh: KATAK, APA, KASUR_RUSAK.]
#### assesment04.go

```go
package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		if ch != ' ' && ch != '\n' {
			t[*n] = ch
			*n++
		}
	}
}
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}
func main() {
	var tab tabel
	var n int

	fmt.Print("Teks: ")
	isiArray(&tab, &n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}

	fmt.Print("Reverse teks: ")
	balikanArray(&tab, n)
	cetakArray(tab, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul09/output/assesment04.png) 
[Program ini digunakan untuk mengolah sebuah teks yang dimasukkan oleh kita, kemudian mengecek apakah teks tersebut merupakan palindrom atau tidak, serta menampilkan hasil pembalikan teks tersebut. Teks dimasukkan karakter per karakter ke dalam array bertipe rune hingga ditemukan tanda titik (.) sebagai penanda akhir input atau kapasitas array telah terpenuhi. Selama proses input, spasi dan karakter newline tidak ikut disimpan ke dalam array agar hanya karakter penting saja yang diproses.

Setelah data tersimpan, program akan mengecek apakah teks tersebut merupakan palindrom dengan cara membandingkan karakter dari depan dan belakang secara berpasangan. Jika semua pasangan karakter sama, maka teks dinyatakan sebagai palindrom. Selanjutnya, program membalik urutan karakter dalam array menggunakan metode pertukaran antara elemen depan dan belakang. Hasil akhir yang ditampilkan adalah status palindrom serta teks yang sudah dibalik urutannya.]


