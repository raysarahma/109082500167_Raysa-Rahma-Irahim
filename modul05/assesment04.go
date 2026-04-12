package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	baris(n)
}
func baris(n int) {
	if n == 1 {
		fmt.Print(1, " ")
	} else {
		fmt.Print(n, " ")
		baris(n - 1)
		fmt.Print(n, " ")
	}
}
