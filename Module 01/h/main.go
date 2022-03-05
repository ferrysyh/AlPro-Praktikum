package main

import "fmt"

func main() {
	var r, t float64
	fmt.Scan(&r, &t)
	fmt.Println(hitungVolume(r, t))
}

func hitungVolume(r, t float64) float64 {
	var pi float64 = 3.14
	return pi * r * r * t
}
