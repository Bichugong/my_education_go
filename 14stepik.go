package main

import "fmt"

func main() {
	var start, end int
	k := 0
	for fmt.Scan(&start, &end); start < end + 1; start++{
		k += start
	}
	fmt.Println(k)
}