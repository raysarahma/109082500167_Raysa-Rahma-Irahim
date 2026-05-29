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