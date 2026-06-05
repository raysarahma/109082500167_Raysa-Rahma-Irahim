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