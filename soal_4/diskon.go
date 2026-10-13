package main

import "fmt"

func main() {
	var totalBelanja, diskon, totalAkhir int
	fmt.Scan(&totalBelanja, &diskon)
	totalAkhir = totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println(totalAkhir)
}