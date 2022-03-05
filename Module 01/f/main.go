package main

import "fmt"

func main() {
	var nilai int
	var jumlah int = 0
	var n int = 0
	var rata2 float64

	fmt.Scan(&nilai)

	for nilai != -1 {
		n++
		jumlah += nilai
		fmt.Scan(&nilai)
	}
	if n == 0 {
		rata2 = 0.0
	} else {
		rata2 = float64(jumlah) / float64(n)
	}
	fmt.Println(rata2)
}
