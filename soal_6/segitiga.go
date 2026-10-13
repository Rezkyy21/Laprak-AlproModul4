package main

import (
	"fmt"
	"math"
)

func main() {
	var xA, yA, xB, yB, xC, yC float64
	var ab, bc, ac, terpanjang float64

	fmt.Scan(&xA, &yA)
	fmt.Scan(&xB, &yB)
	fmt.Scan(&xC, &yC)

	ab = math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
	bc = math.Sqrt(math.Pow(xC-xB, 2) + math.Pow(yC-yB, 2))
	ac = math.Sqrt(math.Pow(xC-xA, 2) + math.Pow(yC-yA, 2))

	terpanjang = math.Max(ab, math.Max(bc, ac))

	fmt.Println(math.Round(terpanjang*100) / 100)
}