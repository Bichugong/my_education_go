package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	array := [3]int {num % 10, num / 10 % 10, num / 100}
	for i := 0; i < 3; i++{
		fmt.Printf("%d", array[i])
	}
}