package main

import "fmt"

func main() {
	var kata string
	var angka1, angka2, hasil int

	fmt.Scanln(&kata, &angka1, &angka2)

	hasil = angka1 + angka2
	fmt.Println("Kata =", kata)
	fmt.Println("Jumlah =", hasil)
}
