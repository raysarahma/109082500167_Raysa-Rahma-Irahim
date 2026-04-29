package main

import "fmt"

func main() {
	var A, B string
	fmt.Print("Klub A: ")
	fmt.Scan(&A)
	fmt.Print("Klub B: ")
	fmt.Scan(&B)

	var a, b int
	var hasil [100]string
	var n int = 0
	i := 1

	for {
		fmt.Print("Pertandingan ", i, ": ")
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			break
		}

		if a > b {
			hasil[n] = A
		} else if b > a {
			hasil[n] = B
		} else {
			hasil[n] = "Draw"
		}

		n++
		i++
	}

	for j := 0; j < n; j++ {
		fmt.Println("Hasil", j+1, ":", hasil[j])
	}
	fmt.Println("Pertandingan selesai")

}
