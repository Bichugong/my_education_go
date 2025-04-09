package main

import "fmt"

func main() {
	var num string
	fmt.Scan(&num)
	first_rune := num[0]
	fmt.Printf("%c\n", first_rune)
}