package main

import "fmt"

func main() {
	var n, t1, t2, t3, jumlah int	// n, t1, t2, t3, jumlah merupakan variabel yang bertipe integer

	fmt.Scan(&n)					// mengambil inputan dari user

	for i := 0; i < n; i++ {		// mengecek apakah i lebih kecil dari n
		fmt.Scan(&t1, &t2, &t3)		// mengambil inputan dari user
		if t1+t2+t3 >= 2 {			// jika t1+t2+t3 lebih besar dari 2
			jumlah++				// jumlah ditambah 1
		}
	}

	fmt.Println(jumlah)				// mencetak jumlah
}
