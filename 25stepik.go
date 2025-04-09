package main

import "fmt"

func main() {
    // Объявляем массив с минимально возможным беззнаковым целым типом
    var workArray [10]uint8

    // Читаем 10 чисел и записываем их в массив
    for i := 0; i < 10; i++ {
        fmt.Scan(&workArray[i])
    }

    // Меняем местами элементы по трем парам индексов
    for j := 0; j < 3; j++ {
        var a, b uint8
        fmt.Scan(&a, &b)
        
        // Меняем местами элементы
        workArray[a], workArray[b] = workArray[b], workArray[a]
    }

    // Выводим результат
    for _, num := range workArray {
        fmt.Print(num, " ")
    }
}