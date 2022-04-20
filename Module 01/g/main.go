package main

import "fmt"

func main() {
	var a, b, c, d, e, f, x int	// a, b, c, d, e, f, x merupakan variabel yang bertipe integer

	fmt.Scan(&a, &b, &c)		// mengambil inputan dari user
	fmt.Scan(&d, &e, &f)		// mengambil inputan dari user
	fmt.Scan(&x)				// mengambil inputan dari user

	if x == 0 || x == 360 {		// jika x sama dengan 0 atau 360
		fmt.Println(a, b, c)	// maka mencetak a, b, c
		fmt.Println(d, e, f)	// maka mencetak d, e, f
	} else if x == 90 {			// jika x sama dengan 90
		fmt.Println(d, a)		// maka mencetak d, a
		fmt.Println(e, b)		// maka mencetak e, b
		fmt.Println(f, c)		// maka mencetak f, c
	} else if x == 180 {		// jika x sama dengan 180
		fmt.Println(f, e, d)	// maka mencetak f, e, d
		fmt.Println(c, b, a)	// maka mencetak c, b, a
	} else if x == 270 {		// jika x sama dengan 270
		fmt.Println(c, f)		// maka mencetak c, f
		fmt.Println(b, e)		// maka mencetak b, e
		fmt.Println(a, d)		// maka mencetak a, d
	}
}
