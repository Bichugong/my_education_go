package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func Podstring(s, p string) int{
		return strings.Index(s, p)
	
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	pod, _ := reader.ReadString('\n')
	pod = strings.TrimSpace(pod)
	fmt.Println(Podstring(text, pod))
}