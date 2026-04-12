package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		n = n -1 
	}
	ganjil(n)
}
func ganjil(n int) {
	if n < 1 {
		return 
	}

	ganjil(n - 2)
	fmt.Print(n, " ")
}