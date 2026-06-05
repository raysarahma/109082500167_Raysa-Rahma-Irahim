# <h1 align="center"> Laporan Praktikum Modul 12 </h1>
<p align="center">[Raysa Rahma Irahim] - [109082500167]</p>

## Unguided 

### 1. [Pada pemilihan ketua RT yang baru saja berlangsung, terdapat 20 calon ketua yang bertanding memperebutkan suara warga. Perhitungan suara dapat segera dilakukan karena warga cukup mengisi formulir dengan nomor dari calon ketua RT yang dipilihnya. Seperti biasa, selalu ada pengisian yang tidak tepat atau dengan nomor pilihan di luar yang tersedia, sehingga data juga harus divalidasi. Tugas Anda untuk membuat program mencari siapa yang memenangkan pemilihan ketua RT. Buatlah program pilkart yang akan membaca, memvalidasi, dan menghitung suara yang diberikan dalam pemilihan ketua RT tersebut. Masukan hanya satu baris data saja, berisi bilangan bulat valid yang kadang tersisipi dengan data tidak valid. Data valid adalah integer dengan nilai di antara 1 s.d. 20 (inklusif). Data berakhir jika ditemukan sebuah bilangan dengan nilai 0. Keluaran dimulai dengan baris berisi jumlah data suara yang terbaca, diikuti baris yang berisi berapa banyak suara yang valid. Kemudian sejumlah baris yang mencetak data para calon apa saja yang mendapatkan suara.]
#### assesment01.go

```go
package main

import "fmt"

func main() {
	var suara [21]int
	var x int
	var masuk, sah int

	fmt.Scan(&x)

	for x != 0 {
		masuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			sah++
		}

		fmt.Scan(&x)
	}

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul12/output/output-assesment01.png) 
[Program ini dibuat untuk menghitung hasil suara pada pemilihan ketua RT. Data suara dimasukkan satu per satu dan akan berhenti ketika memasukkan angka 0 sebagai tanda akhir input. Setiap data yang masuk akan dihitung sebagai suara masuk, kemudian program melakukan pengecekan apakah nomor yang dimasukkan termasuk nomor calon yang valid, yaitu antara 1 sampai 20. Jika valid, suara tersebut akan dihitung sebagai suara sah dan jumlah suaranya akan disimpan sesuai nomor calon yang dipilih. Setelah semua data selesai dibaca, program menampilkan jumlah seluruh suara yang masuk serta jumlah suara sah. Selanjutnya program menampilkan daftar calon yang memperoleh suara beserta banyaknya suara yang diterima oleh masing-masing calon.]


### 2. [Berdasarkan program sebelumnya, buat program pilkart yang mencari siapa pemenang pemilihan ketua RT. Sekaligus juga ditentukan bahwa wakil ketua RT adalah calon yang mendapatkan suara terbanyak kedua. Jika beberapa calon mendapatkan suara terbanyak yang sama, ketua terpilih adalah dengan nomor peserta yang paling kecil dan wakilnya dengan nomor peserta terkecil berikutnya. Masukan hanya satu baris data saja, berisi bilangan bulat valid yang kadang tersisipi dengan data tidak valid. Data valid adalah bilangan bulat dengan nilai di antara 1 s.d. 20 (inklusif). Data berakhir jika ditemukan sebuah bilangan dengan nilai 0. Keluaran dimulai dengan baris berisi jumlah data suara yang terbaca, diikuti baris yang berisi berapa banyak suara yang valid. Kemudian tercetak calon nomor berapa saja yang menjadi pasangan ketua RT dan wakil ketua RT yang baru.]
#### assesment02.go

```go
package main

import "fmt"

func main() {
	var suara [21]int
	var x int
	var masuk, sah int

	fmt.Scan(&x)

	for x != 0 {
		masuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			sah++
		}

		fmt.Scan(&x)
	}

	ketua := 1
	wakil := 2

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			wakil = i
			break
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if suara[i] > suara[wakil] {
				wakil = i
			} else if suara[i] == suara[wakil] && i < wakil {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul12/output/output-assesment02.png)
[Program ini digunakan untuk menentukan ketua RT dan wakil ketua RT berdasarkan hasil pemungutan suara. Program membaca data suara, memvalidasi suara yang bernilai 1 sampai 20, lalu menghitung jumlah suara yang diperoleh setiap calon. Setelah seluruh data diproses, program mencari calon dengan suara terbanyak sebagai ketua RT dan calon dengan suara terbanyak kedua sebagai wakil ketua RT. Hasil akhirnya berupa jumlah suara masuk, jumlah suara sah, serta nomor calon yang terpilih sebagai ketua dan wakil ketua RT.]


### 3. [Diberikan n data integer positif dalam keadaan terurut membesar dan sebuah integer lain k, apakah bilangan k tersebut ada dalam daftar bilangan yang diberikan? Jika ya, berikan indeksnya, jika tidak sebutkan "TIDAK ADA". Masukan terdiri dari dua baris. Baris pertama berisi dua buah integer positif, yaitu n dan k. n menyatakan banyaknya data, dimana 1 < n <= 1000000. k adalah bilangan yang ingin dicari. Baris kedua berisi n buah data integer positif yang sudah terurut membesar. Keluaran terdiri dari satu baris saja, yaitu sebuah bilangan yang menyatakan posisi data yang dicari (k) dalam kumpulan data yang diberikan. Posisi data dihitung dimulai dari angka 0. Atau memberikan keluaran "TIDAK ADA" jika data k tersebut tidak ditemukan dalam kumpulan. Program yang dibangun harus menggunakan subprogram dengan mengikuti kerangka yang sudah diberikan berikut ini.]
#### assesment03.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var kiri, kanan, median int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		median = (kiri + kanan) / 2

		if data[median] == k {
			return median
		} else if data[median] < k {
			kiri = median + 1
		} else {
			kanan = median - 1
		}
	}

	return -1
}

func main() {
	var n, k int
	var hasil int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil = posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_!]  (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul12/output/output-assesment03.png)
[Program ini digunakan untuk mencari suatu bilangan pada kumpulan data yang sudah terurut dengan menggunakan metode Binary Search. Program akan membaca jumlah data, nilai yang ingin dicari, serta data-data yang tersedia. Selanjutnya pencarian dilakukan dengan memeriksa elemen yang berada di posisi tengah array. Jika nilai yang dicari lebih besar atau lebih kecil dari nilai tengah, maka proses pencarian akan dilanjutkan hanya pada bagian array yang masih mungkin mengandung nilai tersebut. Jika data ditemukan, program akan menampilkan indeksnya. Namun jika data tidak ditemukan, program akan menampilkan tulisan "TIDAK ADA". Penggunaan Binary Search membuat proses pencarian menjadi lebih cepat dan efisien karena tidak perlu memeriksa seluruh data satu per satu.]  

