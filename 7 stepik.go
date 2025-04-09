package main

import "fmt"

func main() {
	var num, hours, minutes int
	fmt.Scan(&num)
	hours = num / 30
	minutes = (num - 30 * hours) * 2
	fmt.Println("It is ", hours, "hours", minutes, "minutes")
}