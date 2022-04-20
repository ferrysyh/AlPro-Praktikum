package main

import "fmt"

func main() {
	var kata string               		// kata merupakan variabel yang bertipe string
	var angka1, angka2, hasil int 		// angka1, angka2, dan hasil merupakan variabel yang bertipe integer

	fmt.Scanln(&kata, &angka1, &angka2)	// mengambil inputan dari user

	hasil = angka1 + angka2 			// menjumlahkan angka1 dan angka2

	fmt.Println("Kata =", kata)    		// mencetak kata
	fmt.Println("Jumlah =", hasil) 		// mencetak hasil
}
