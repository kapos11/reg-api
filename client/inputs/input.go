package input

import (
	"fmt"
	"main/models"
)

func GetUserInput() models.User {
	var user models.User
	fmt.Print("Enter name: ")
	fmt.Scanln(&user.Name)
	fmt.Print("Enter email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Enter password: ")
	fmt.Scanln(&user.Password)
	return user
}
