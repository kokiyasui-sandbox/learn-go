package main

import "fmt"

func main() {
	message := "Hi 👩 and 👱"

	message_rune := []rune(message)

	char := string(message_rune[3])

	fmt.Println(char)
}
