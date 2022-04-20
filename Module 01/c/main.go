package main

import "fmt"

func main() {
	var x int 				// x merupakan variabel yang bertipe integer

	fmt.Scanln(&x) 			// mengambil inputan dari user

	if x%3 == 0 {			// mengecek apakah x habis dibagi 3
		fmt.Println("Fizz")	// jika iya, maka mencetak Fizz
	} 

	if x%5 == 0 { 			// mengecek apakah x habis dibagi 5
		fmt.Println("Bazz")	// jika iya, maka mencetak Bazz
	}
}
