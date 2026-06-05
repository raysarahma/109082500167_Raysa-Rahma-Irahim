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