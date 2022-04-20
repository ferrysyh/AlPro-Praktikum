package main

import "fmt"

func main() {
	var r int 																		// r merupakan variabel yang bertipe integer
	var luas_lingkaran float64 														// luas_lingkaran merupakan variabel yang bertipe float64

	fmt.Scan(&r) 																	// mengambil inputan dari user

	luas_lingkaran = (float64(22) / float64(7)) * float64(r) 						// menghitung luas lingkaran
	
	fmt.Println("Luas Lingkaran dengan jari-jari =", r, "adalah", luas_lingkaran)	// mencetak luas lingkaran
}
