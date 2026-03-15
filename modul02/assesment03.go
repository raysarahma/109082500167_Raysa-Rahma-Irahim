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
