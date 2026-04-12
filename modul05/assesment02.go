package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	bintang(n, 1)
}
func bintang(n int, i int) {
	if i > n {
		return
	}
	for j := 0; j < i; j++{
		fmt.Print("*")
	}
	fmt.Println()
	
	bintang(n, i+1)
}