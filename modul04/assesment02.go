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
