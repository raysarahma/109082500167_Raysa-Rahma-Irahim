# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Raysa Rahma Irahim] - [109082500167]</p>

## Unguided 

### 1. [Soal 1]
#### assesment01.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/modul01/blob/main/output/output-assesment01.png)
[Program ini digunakan untuk menerima 3 input string, menampilkan urutan awalnya sesuai dengan urutan input dimasukkan, kemudian ditukar posisi nilai inputnya dengan rotasi (satu-> dua, dua -> tiga, tiga -> satu) menggunakan variabel sementara (temp) lalu menampilkan urutan yang baru.]


### 2. [Soal 2]
#### assesment02.go

```go
package main

import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)

		if !(gelas1 == "merah" && gelas2 == "kuning" && gelas3 == "hijau" && gelas4 == "ungu") {
			berhasil = false
		}
	}
	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/modul01/blob/main/output/output-assesment02.png)
[Program ini untuk mengecek susunan warna pada percobaan kimia apakah sudah benar atau belum. Jadi di program ini kita diminta memasukkan warna dari 4 gelas sebanyak 5 kali percobaan. Setiap 1 kali percobaan itu kita masukkan 4 warna yang mewakili isi dari 1 sampai gelas 4. Urutan warna yang dianggap benar itu adalah : merah, kuning, hijau, ungu.
Jadi setiap kita masukkan warna, programnya langsung bandingin urutan warna yang kita masukkan dengan urutannya yang dianggap benar. Kalau urutannya sama persis, maka percobaannya dianggap benar(true). Tapi kalau ada satu saja yang berbeda atau urutannya tidak sesuai, maka percobaannya salah(false).]


### 3. [Soal 3]
#### assesment03.go

```go
package main

import "fmt"

func main() {

	var berat, kg, sisa, biaya1, biaya2, total int

	fmt.Print("Berat parsel(gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biaya1 = kg * 10000

	if kg > 10 {
		biaya2 = 0
	} else {
		if sisa >= 500 {
			biaya2 = sisa * 5
		} else {
			biaya2 = sisa * 15
		}
	}

	total = biaya1 + biaya2

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biaya1, "+ Rp.", biaya2)
	fmt.Println("Total biaya: Rp.", total)

}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_!] (https://github.com/raysarahma/109082500167_Raysa-Rahma-Irahim/modul01/blob/main/output/output-assesment03.png)
[Program ini untuk menghitung biaya kirim paket berdasarkan berat parsel. Kita diminta memasukkan berat parsel dalam satuan (gram). Setelah itu programnya bakal hitung berapa beratnya dalam kilogram dan berapa sisa gramnya. Berat dalam kilogram dapat dari pembagian berat dengan 1000, sedangkan sisa gramnya diproleh dari sisa pembagian yang tadi. 
Lalu program bakal hitung biaya pengirimannya, biaya untuk 1kg adalah Rp.10.000 jadi jumlah kilogram itu dikalikan dengan Rp.10.000 untuk biaya utama. Selanjutnya program hitung biaya dari sisa beratnya, jika sisa beratnya 500 gram atau lebih, maka biaya bertambah Rp.5 per gram. Namun kalau sisa beratnya kurang dari 500 gram, biaya tambahannya menjadi Rp15 per gram. 
Selain itu ada aturan khusus, yaitu jika total berat parsel lebih dari 10 kg, maka sisa gramnya digratiskan. Jadi yang dibayar hanya biaya kilogramnya saja.]