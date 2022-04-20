package main

import "fmt"

func main() {
	var nama, namamin, namamax string
	var p int
	var faktor, jumlah, min, max, rata2 float64

	namamin = ""
	namamax = ""
	fmt.Scan(&nama)
	for nama != "STOP" {
		fmt.Scan(&p)
		jumlah = 0
		for i := 0; i < 3*p; i++ {
			fmt.Scan(&faktor)
			jumlah += faktor
		}
		rata2 = jumlah / float64(3*p)
		if namamin == "" || namamin != "" && min > rata2 {
			min = rata2
			namamin = nama
		}
		if namamax == "" || namamax != "" && max < rata2 {
			max = rata2
			namamax = nama
		}
		fmt.Scan(&nama)
	}
	fmt.Printf("%s %.2f\n", namamax, max)
	fmt.Printf("%s %.2f\n", namamin, min)
}
