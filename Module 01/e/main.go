package main

import "fmt"

func main() {
	var string string
	fmt.Scan(&string)

	for string != "selesai" {
		fmt.Scan(&string)
	}
}
