package main

import "fmt"

func main() {
	var n, t1, t2, t3, jumlah int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t1, &t2, &t3)
		if t1+t2+t3 >= 2 {
			jumlah++
		}
	}
	fmt.Println(jumlah)
}
