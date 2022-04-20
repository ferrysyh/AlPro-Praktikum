package main

import "fmt"

func main() {
	var nilai int								// nilai merupakan variabel yang bertipe integer
	var jumlah int = 0							// jumlah merupakan variabel yang bertipe integer dan diinisialisasi dengan nilai 0
	var n int = 0								// n merupakan variabel yang bertipe integer dan diinisialisasi dengan nilai 0
	var rata2 float64							// rata2 merupakan variabel yang bertipe float64

	fmt.Scan(&nilai)							// mengambil inputan dari user

	for nilai != -1 {							// mengecek apakah nilai tidak sama dengan -1
		n++										// n ditambah 1
		jumlah += nilai							// jumlah ditambah nilai
		fmt.Scan(&nilai)						// mengambil inputan dari user
	}

	if n == 0 {									// jika n sama dengan 0
		rata2 = 0.0								// rata2 diisi dengan 0.0
	} else {									// jika tidak
		rata2 = float64(jumlah) / float64(n)	// rata2 diisi dengan jumlah dibagi n
	}

	fmt.Println(rata2)							// mencetak rata2
}
