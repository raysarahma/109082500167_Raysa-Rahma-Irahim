# <h1 align="center"> Laporan Praktikum Modul 3 </h1>
<p align="center">[Raysa Rahma Irahim] - [109082500167]</p>

## Unguided 

### 1. [Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n! / (n−r)! , sedangkan C(n, r) = n! / r!(n−r)!]
#### assesment01.go

```go
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}
func permutation(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}
func combination(n int, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul03/output/output-assesment1.png) 
[Program ini digunakan untuk menghitung nilai permutasi dan kombinasi dari dua pasangan bilangan yang akan kita masukkan. Program menerima empat buah bilangan, lalu angka tersebut digunakan sebagai pasangan untuk menghitung nilai permutasi dan kombinasi. Untuk perhitungan itu, program menggunakan fungsi faktorial sebagai fungsi bantuan. Program ini juga dibuat untuk menunjukkan penggunaan fungsi dalam bahasa Go, sehingga perhitungan bisa dipisahkan dalam beberapa fungsi agar program lebih rapi dan mudah dipahami. ]


### 2. [Diberikan tiga buah fungsi matematika yaitu f (x) = x2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!]
#### assesment02.go

```go
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul03/output/output-assesment2.png)
[Program menerima tiga buah angka dari yang kita masukkan, yaitu a, b, dan c, kemudian setiap angka diproses dengan tiga fungsi yang berbeda yaitu f, g, dan h. Fungsi f untuk menghitung kuadrat sebuah angka, fungsi g untuk mengurangi angka dengan 2, dan fungsi h untuk menambah angka dengan 1. Pada program ini, fungsi tersebut dipanggil secara berurutan sehingga hasil dari satu fungsi akan menjadi input untuk fungsi berikutnya. Denga cara ini, program menunjukkan proses pengolahan nilai dilakukan secara bertahap menggunakan beberapa fungsi.]


### 3. [[Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2". Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah: jarak = √(a − c)2 + (b − d)2 dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak. function jarak(a,b,c,d : real) -> real {Mengembalikan jarak antara titik (a,b) dan titik (c,d)} function didalam(cx,cy,r,x,y : real) -> boolean {Mengembalikan true apabila titik (x,y) berada di dalam lingkaran yang memiliki titik pusat (cx,cy) dan radius r} Catatan: Lihat paket math dalam lampiran untuk menggunakan fungsi math.Sqrt() untuk menghitung akar kuadrat.]
#### assesment03.go

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
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/blob/main/modul03/output/output-assesment3.png)
[Program ini untuk mengecek sebuah titik apakah berada di dalam atau di luar dua lingkaran. Program menerima input yang kita masukkan berupa titik pusat dan jari-jari dari dua lingkaran, serta satu titik yang mau dicek posisinya. Setelah itu, program bakal hitung jarak antara titik tersebut dengan pusat masing-masing lingkaran memakai fungsi jarak. Lalu jarak digunakan pada fungsi didalam untuk mengecek apakah titik tersebut masih di dalam lingkaran atau tidak dengan membandingkan jarak dan jari-jari lingkaran. Dari hasil itu, program bakal menampilkan apakah titik berada di dalam lingkaran 1, di dalam lingkaran 2, di dalam kedua lingkaran, atau di luar kedua lingkaran.]  