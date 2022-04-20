package main

import "fmt"

func main() {
	var n, pos, neg, num1, num2 int

	fmt.Scan(&n)
	pos = 0
	neg = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&num1, &num2)
		if num1 >= 0 {
			pos += num1
		} else {
			neg += num1
		}
		if num2 >= 0 {
			pos += num2
		} else {
			neg += num2
		}
	}
	fmt.Println("Negative:", neg)
	fmt.Println("Positive:", pos)
}
