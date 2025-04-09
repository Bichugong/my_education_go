package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(num / 100 + num / 10 % 10 + num % 10)
}