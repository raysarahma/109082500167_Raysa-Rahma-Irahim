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

	ketua := 1
	wakil := 2

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			wakil = i
			break
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if suara[i] > suara[wakil] {
				wakil = i
			} else if suara[i] == suara[wakil] && i < wakil {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}