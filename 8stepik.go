package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num == 0 {
		fmt.Println("Ноль")
	} else if num > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}
}