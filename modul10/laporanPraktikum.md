# <h1 align="center"> Laporan Praktikum Modul 10 </h1>
<p align="center">[Raysa Rahma Irahim] - [109082500167]</p>

## Unguided 

### 1. [Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.]
#### assesment01.go

```go
package main

import "fmt"

type arrKelinci [1000]float64

func main() {
	var berat arrKelinci
	var n int
	var min, max float64

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	min = berat[0]
	max = berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f kg\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul10/output/assesment01.png) 
[Program ini digunakan untuk mencari berat kelinci terkecil dan terbesar dari sejumlah data yang dimasukkan pengguna. Data berat kelinci disimpan ke dalam array, kemudian elemen pertama dijadikan nilai awal untuk berat minimum dan maksimum. Setelah itu, program melakukan perulangan untuk membandingkan setiap data dengan nilai minimum dan maksimum yang sudah ada. Jika ditemukan berat yang lebih kecil, nilai minimum akan diperbarui, sedangkan jika ditemukan berat yang lebih besar, nilai maksimum akan diperbarui. Setelah seluruh data selesai diperiksa, program menampilkan berat kelinci terkecil dan terbesar yang ditemukan.]


### 2. [Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.]
#### assesment02.go

```go
package main

import "fmt"

type arrIkan [1000]float64

func main() {
	var ikan arrIkan
	var wadah [1000]float64
	var x, y int
	var i, jumlahWadah int
	var totalSemua float64

	fmt.Scan(&x, &y)

	for i = 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah = (x + y - 1) / y

	for i = 0; i < x; i++ {
		wadah[i/y] += ikan[i]
		totalSemua += ikan[i]
	}

	for i = 0; i < jumlahWadah; i++ {
		fmt.Printf("Total berat wadah %d: %.2f kg\n", i+1, wadah[i])
	}

	fmt.Printf("\nRata-rata berat per wadah: %.2f kg\n",
		totalSemua/float64(jumlahWadah))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul10/output/assesment02.png)
[Program ini digunakan untuk menghitung total berat ikan pada setiap wadah serta rata-rata berat ikan per wadah. Data berat ikan disimpan dalam array dengan kapasitas 1000 elemen. Pengguna memasukkan jumlah ikan (x) dan kapasitas setiap wadah (y), kemudian memasukkan berat masing-masing ikan. Program menghitung jumlah wadah yang diperlukan berdasarkan banyaknya ikan dan kapasitas setiap wadah. Selanjutnya, setiap berat ikan dijumlahkan ke wadah yang sesuai berdasarkan urutan input, sekaligus menghitung total berat seluruh ikan. Setelah proses selesai, program menampilkan total berat pada setiap wadah dan menghitung rata-rata berat per wadah dengan membagi total berat seluruh ikan dengan jumlah wadah yang digunakan.]


### 3. [Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya. Buatlah program dengan spesifikasi subprogram sebagai berikut:]
#### assesment03.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var berat arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(berat, n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul10/output/assesment03.png)
[Program ini digunakan untuk mengolah data berat badan balita menggunakan bahasa Go. Data berat disimpan dalam array arrBalita dengan tipe float64 supaya bisa menyimpan angka desimal. Fungsi hitungMinMax dipakai untuk mencari berat minimum dan maksimum dengan cara membandingkan setiap data dalam array menggunakan perulangan. Nilai minimum dan maksimum dikirim menggunakan pointer agar hasil perhitungannya langsung tersimpan pada variabel di fungsi utama. Selain itu, terdapat fungsi rerata yang digunakan untuk menghitung rata-rata berat balita dengan menjumlahkan seluruh data lalu membaginya dengan jumlah data yang dimasukkan. Pada fungsi main, pengguna diminta memasukkan jumlah data dan berat masing-masing balita, kemudian program akan menampilkan berat minimum, maksimum, dan rata-rata berat balita.]  

