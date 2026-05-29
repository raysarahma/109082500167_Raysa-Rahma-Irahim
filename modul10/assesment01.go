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