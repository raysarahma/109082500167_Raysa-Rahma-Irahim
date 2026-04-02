# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">[Raysa Rahma Irahim] - [109082500167]</p>

## Unguided 

### 1. [Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap 𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑. Catatan: permutasi (P) dan kombinasi (C) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut! 𝑃(𝑛, 𝑟) = 𝑛!/(𝑛−𝑟)! , sedangkan 𝐶(𝑛, 𝑟) = 𝑛!/𝑟!(𝑛−𝑟)!]
#### assesment01.go

```go
package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul04/output/output-assesment01.png) 
[Program ini dibuat buat menghitung permutasi dan kombinasi dari dua pasangan bilangan sesuai soal. Inputnya ada empat bilangan, yaitu a, b, c, dan d, dengan syarat a ≥ c dan b ≥ d. Nah, pasangan (a, c) dipakai buat perhitungan pertama, sedangkan (b, d) buat perhitungan kedua. Hasilnya nanti ditampilkan dalam dua baris, masing-masing berisi nilai permutasi dan kombinasi.Di dalam program, ada fungsi factorial yang dipakai buat menghitung faktorial dengan perulangan. Terus hasilnya dipakai di fungsi permutation dan combination sesuai rumus yang ada. Setelah input dimasukin, program langsung ngitung dan nampilin hasilnya. Jadi, program ini mengubah rumus permutasi dan kombinasi jadi bentuk program biar bisa dihitung.]


### 2. [Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal.Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.]
#### assesment02.go

```go
package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int

	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama string
	var soal, skor int

	var pemenang string
	var maxSoal, minSkor int

	maxSoal = 0
	minSkor = 1000000

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > maxSoal {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		} else if soal == maxSoal {
			if skor < minSkor {
				minSkor = skor
				pemenang = nama
			}
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul04/output/output-assesment02.png)
[Program ini dibuat buat menentukan pemenang kompetisi dari beberapa peserta berdasarkan jumlah soal yang diselesaikan dan total waktu. Setiap peserta ngerjain 8 soal, dan yang dihitung hanya yang waktunya ≤ 300 menit. Kalau lebih, dianggap tidak selesai. Program memakai prosedur hitungSkor buat menghitung jumlah soal yang berhasil dan total waktunya. Di main, nama peserta dibaca satu per satu sampai “Selesai”, lalu hasil tiap peserta dibandingin buat cari yang terbaik. Pemenangnya adalah yang menyelesaikan soal paling banyak, dan kalau sama, yang waktunya paling kecil. Program menampilkan nama pemenang, jumlah soal, dan total waktunya.]


### 3. [Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah: 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 Untuk   suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1. Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal. prosedure cetakDeret(in n : integer ) Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000. Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi]
#### assesment03.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n, " ")

		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var n int

	fmt.Scan(&n)

	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul04 /output/output-assesment03.png)
[Program ini dibuat buat nampilin deret bilangan sesuai aturan yang ada di soal. Program mulai dari satu angka awal n yang diinput. Angka ini nanti diproses terus sampai akhirnya jadi 1, dan setiap nilainya bakal ditampilkan dalam satu baris dipisahin spasi. Prosesnya ada di prosedur cetakDeret. Di dalamnya, program nge-loop terus, setiap langkah nilai n langsung dicetak dulu. Setelah itu dicek, kalau n genap maka dibagi 2, kalau ganjil diubah jadi 3n + 1. Hasilnya dipakai lagi buat langkah berikutnya. Proses ini diulang terus sampai n jadi 1, baru berhenti. Di bagian main, program cuma baca satu input angka, lalu manggil prosedur cetakDeret buat ngejalanin dan nampilin deretnya. Jadi, program ini ngikutin aturan deret dari angka awal sampai 1 dan nampilin semua langkahnya.]  