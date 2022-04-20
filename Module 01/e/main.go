package main

import "fmt"

func main() {
	var string string 			// string merupakan variabel yang bertipe string

	fmt.Scan(&string) 			// mengambil inputan dari user

	for string != "selesai" { 	// mengecek apakah string tidak sama dengan selesai
		fmt.Println(string) 	// jika iya, maka mencetak string
		fmt.Scan(&string)   	// mengambil inputan dari user
	}
}
