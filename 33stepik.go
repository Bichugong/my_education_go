package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text) // Удаляем все пробельные символы по краям включая \n

	if len(text) == 0 {
		fmt.Println("Wrong")
		return
	}

	runes := []rune(text) // Преобразуем в руны для корректной работы с Unicode
	firstChar := runes[0]
	lastChar := runes[len(runes)-1]

	if unicode.IsUpper(firstChar) && lastChar == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}