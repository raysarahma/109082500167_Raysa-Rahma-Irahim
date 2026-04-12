# <h1 align="center"> Laporan Praktikum Modul 5 </h1>
<p align="center">[Raysa Rahma Irahim] - [109082500167]</p>

## Unguided 

### 1. [Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.]
#### assesment01.go

```go
package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul05/output/output-assesment01.png) 
[Program ini dipakai buat nampilin deret Fibonacci sampai angka ke-n. Kita cukup masukin satu angka, misalnya 5 atau 10. Setelah itu, program akan ngulang dari 0 sampai n, dan tiap angka itu dihitung pakai fungsi fibonacci lalu langsung ditampilin hasilnya. Fungsi fibonacci ini kerjanya rekursif, jadi dia manggil dirinya sendiri. Kalau angkanya 0 hasilnya 0, kalau 1 hasilnya 1. Nah kalau lebih dari itu, dia ambil dua angka sebelumnya lalu dijumlahin. Jadi misalnya 2 itu dari 1+1, 3 dari 1+2, dan seterusnya. Intinya, program ini otomatis bikin deret Fibonacci dari 0 sampai n sesuai angka yang kita masukin.]


### 2. [Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.]
#### assesment02.go

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	bintang(n, 1)
}
func bintang(n int, i int) {
	if i > n {
		return
	}
	for j := 0; j < i; j++{
		fmt.Print("*")
	}
	fmt.Println()
	
	bintang(n, i+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul05/output/output-assesment02.png)
[Program ini digunakan untuk menampilkan pola bintang secara bertahap ke bawah. Kita cukup memasukkan satu angka n, misalnya 5, yang artinya nanti akan ada 5 baris bintang. Program mulai dari baris pertama (i = 1), terus tiap baris dia cetak bintang sesuai nilai i, jadi awalnya satu bintang, lalu dua, tiga, sampai n. Setelah cetak satu baris, fungsi bintang manggil dirinya sendiri dengan i ditambah 1 supaya lanjut ke baris berikutnya. Proses ini terus diulang sampai i lebih dari n, baru berhenti. Jadi intinya i itu buat nentuin jumlah bintang tiap baris, dan rekursif dipakai buat pindah ke baris selanjutnya.]


### 3. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).]
#### assesment03.go

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	faktor(n,1)
}
func faktor (n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul05/output/output-assesment03.png)
[Program ini dipakai buat nampilin faktor dari suatu angka. Kita cukup masukin satu angka, misalnya 12. Setelah itu, program bakal ngecek angka dari 1 sampai n, dan tiap angka itu dicek apakah bisa membagi n atau nggak. Kalau bisa, berarti itu termasuk faktor, jadi langsung ditampilin. Fungsi faktor ini pakai rekursif, jadi dia manggil dirinya sendiri terus sambil nambah nilai i. Awalnya i = 1, lalu dicek satu per satu sampai i lebih dari n. Di dalamnya ada pengecekan n % i == 0, artinya kalau n habis dibagi i, maka i itu faktor dari n. Proses ini terus berulang sampai semua angka dari 1 sampai n selesai dicek. Jadi intinya, program ini otomatis nyari dan nampilin semua faktor dari angka yang kita masukin.]  

### 4. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.]
#### assesment04.go

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	baris(n)
}
func baris(n int) {
	if n == 1 {
		fmt.Print(1, " ")
	} else {
		fmt.Print(n, " ")
		baris(n - 1)
		fmt.Print(n, " ")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul05/output/output-assesment04.png) 
[Program ini dipakai buat nampilin barisan angka dari n turun ke 1, lalu naik lagi ke n. Kita cukup masukin satu angka, misalnya 5. Setelah itu, program bakal langsung jalanin fungsi baris buat nampilin urutannya. Fungsi baris ini pakai rekursif, jadi dia manggil dirinya sendiri. Awalnya dia bakal cetak angka n dulu, misalnya 5, terus manggil lagi dengan nilai n-1 jadi 4, lalu 3, sampai akhirnya nyampe 1. Nah pas udah di 1, dia berhenti turun. Setelah itu, pas balik dari rekursif, dia lanjut cetak lagi angkanya, jadi naik lagi dari 2, 3, sampai ke n. Makanya outputnya jadi kayak: 5 4 3 2 1 2 3 4 5. Intinya, program ini turun dulu sampai 1 pakai rekursif, terus pas balik dia naik lagi ke atas.]


### 5. [Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.]
#### assesment05.go

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		n = n -1 
	}
	ganjil(n)
}
func ganjil(n int) {
	if n < 1 {
		return 
	}

	ganjil(n - 2)
	fmt.Print(n, " ")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul05/output/output-assesment05.png)
[Program ini dipakai buat nampilin bilangan ganjil dari 1 sampai n. Kita cukup masukin satu angka, misalnya 10. Kalau angka yang dimasukin itu genap, program bakal ngurangin 1 dulu supaya jadi ganjil, jadi misalnya 10 jadi 9. Setelah itu, program manggil fungsi ganjil buat nampilin angkanya. Fungsi ini pakai rekursif, jadi dia manggil dirinya sendiri sambil ngurangin nilai n sebanyak 2, jadi loncat antar bilangan ganjil aja (misalnya 9, 7, 5, dan seterusnya). Intinya, program ini otomatis nampilin semua bilangan ganjil dari 1 sampai n dengan bantuan rekursif.]


### 6. [Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan. Masukan terdiri dari bilangan bulat x dan y. Keluaran terdiri dari hasil x dipangkatkan y. Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan import "math".]
#### assesment06.go

```go
package main
import "fmt"
func main() {
	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(pangkat(x, y))
}
func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pangkat(x, y-1)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul05/output/output-assesment06.png)
[Program ini buat ngitung hasil pangkat. Kita masukin dua angka, misalnya 2 dan 3, artinya mau hitung 2 pangkat 3. Setelah itu, program pakai fungsi rekursif buat ngali angka x terus menerus sampai y habis. Kalau y sudah 0, hasilnya langsung 1. Jadi misalnya 2³ itu dihitung jadi 2 × 2 × 2. Intinya, program ini ngulang perkalian x sebanyak y kali pakai rekursif.]  