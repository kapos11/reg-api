package main

import (
	"fmt"
	"main/api"
	input "main/inputs"
)

func main() {
	fmt.Println("Registration Form :")

	user := input.GetUserInput()
	response := api.RegisterUser(user)

	fmt.Println("Server response:", response)
}
