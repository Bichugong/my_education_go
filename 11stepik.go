package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	first_half := num / 1000
	second_half := num % 1000
	if first_half / 100 + first_half / 10 % 10 + first_half % 10 == second_half / 100 + second_half / 10 % 10 + second_half % 10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}