package main

import "fmt"

func main() {
	var s string = "Это строка"

	fmt.Printf("Длина строки: %d байт\n", len(s))

	// Получим подстроку строки
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\"\n", s[7:])
}