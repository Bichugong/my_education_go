package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	first_dig := num / 100
	second_dig := num / 10 % 10
	third_dig := num % 10
	if first_dig != second_dig && second_dig != third_dig && first_dig != third_dig {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}