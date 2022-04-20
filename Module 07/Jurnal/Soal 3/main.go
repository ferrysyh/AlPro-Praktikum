package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	konversi(x, &y)
	fmt.Println(y)
}

func pangkat(a, b int) int {
	var hasil int
	hasil = 1
	for i := 0; i < b; i++ {
		hasil *= a
	}
	return hasil
}

func konversi(biner int, desimal *int) {
	var j, x int
	*desimal = 0
	j = 0
	for biner > 0 {
		x = biner % 10
		*desimal += x * pangkat(2, j)
		biner /= 10
		j++
	}
}
