package main

import "fmt"

func main() {
	var r, t float64						// r, t merupakan variabel yang bertipe float64

	fmt.Scan(&r, &t)						// mengambil inputan dari user

	fmt.Println(hitungVolume(r, t))			// mencetak hasil dari fungsi hitungVolume
}

func hitungVolume(r, t float64) float64 {	// fungsi hitungVolume dengan parameter r dan t yang bertipe float64
	var pi float64 = 3.14					// pi merupakan variabel yang bertipe float64 dan diinisialisasi dengan nilai 3.14

	return pi * r * r * t					// menghitung volume
}
