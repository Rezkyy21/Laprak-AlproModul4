package main

import "fmt"

func main() {
	var c1, c2 rune

	fmt.Scanf("%c%c", &c1, &c2)

	khusus1 := !((c1 >= 'A' && c1 <= 'Z') || (c1 >= 'a' && c1 <= 'z') || (c1 >= '0' && c1 <= '9'))
	khusus2 := !((c2 >= 'A' && c2 <= 'Z') || (c2 >= 'a' && c2 <= 'z') || (c2 >= '0' && c2 <= '9'))

	fmt.Println(khusus1, khusus2)
}