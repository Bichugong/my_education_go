package main

import "fmt"
import "math"

func main() {
	var x, p, y float64
	fmt.Scan(&x, &p, &y)
	volume := 0

	for ; x < y; {
		x *= 1 + p / 100
		x = math.Floor(x)
		volume += 1
	}
	fmt.Println(volume)
}