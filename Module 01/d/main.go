package main

import "fmt"

func main() {
	var a, b, c, d int						// a, b, c, dan d merupakan variabel yang bertipe integer

	fmt.Scan(&a, &b, &c, &d) 				// mengambil inputan dari user

	if d > a && d > b && d > c {			// mengecek apakah d lebih besar dari a, b, dan c
		fmt.Println("Ada rekor baru")		// jika iya, maka mencetak Ada rekor baru
	} else {								// jika tidak
		fmt.Println("Tidak ada rekor baru")	// maka mencetak Tidak ada rekor baru
	}
}
