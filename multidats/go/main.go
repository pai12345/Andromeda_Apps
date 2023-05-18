package main

import (
	"fmt"

	"github.com/pai12345/Andromeda_Apps/multidats/go/userdats"
)

func main() {
	fmt.Println("Welcome to multidats")
	userdats.GenerateUserdata("Rahul", "Pai", "01-01-2000", true, "+1234")
}
