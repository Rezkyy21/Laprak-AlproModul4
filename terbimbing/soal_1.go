package main

import "fmt"

func main() {
	var h, btm float64
	const p = 1025.0
	const g = 9.8

	fmt.Scan(&h)
	fmt.Scan(&btm)
	Tekanan := p * g * h
	fmt.Print(Tekanan)
	fmt.Println(Tekanan <= btm)
}