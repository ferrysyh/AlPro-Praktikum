package main

import "fmt"

func main() {
	var i, pengunjung, jumlah, penurunan, sebelum int

	jumlah = 0
	sebelum = 0
	penurunan = 0
	i = 0
	for penurunan != 3 {
		i++
		fmt.Print("hari ", i, " : ")
		fmt.Scan(&pengunjung)
		for pengunjung < 0 || pengunjung > 100 {
			fmt.Print("hari ", i, " : ")
			fmt.Scan(&pengunjung)
		}
		jumlah += pengunjung
		if pengunjung < sebelum {
			penurunan++
		} else {
			penurunan = 0
		}
		sebelum = pengunjung
	}
	fmt.Println("Museum buka selama", i, "hari")
	fmt.Printf("Rata-rata pengunjung %.2f orang", float64(jumlah)/float64(i))
}
