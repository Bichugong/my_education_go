package main

import "fmt"

func main(){
    var a int
	fmt.Print("Введите число: ")
    fmt.Scan(&a)
    fmt.Println((a*2) + 100)
}