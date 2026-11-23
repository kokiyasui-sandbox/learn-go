package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "привет"}

	greetings_sub1 := greetings[:2]
	greetings_sub2 := greetings[1:4]
	greetings_sub3 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(greetings_sub1)
	fmt.Println(greetings_sub2)
	fmt.Println(greetings_sub3)
}
