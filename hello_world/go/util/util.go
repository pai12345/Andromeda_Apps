package util

import "fmt"

func PrintText(greeting string) string {
	greeting_text := fmt.Sprintf("Welcome to %v", greeting)
	fmt.Println(greeting_text)
	return greeting_text
}
