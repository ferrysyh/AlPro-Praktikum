package main

import "fmt"

func main() {
	var r int
	var luas_lingkaran float64
	
	fmt.Scan(&r)

	luas_lingkaran = (float64(22) / float64(7)) * float64(r)
	fmt.Println("Luas Lingkaran dengan jari-jari =", r, "adalah", luas_lingkaran)
}
