package main

import "fmt"

type User struct {
	ID       int
	Name     string
	Email    string
	Age      int
	IsActive bool
}

func main() {
	user := User{
		ID:       1,
		Name:     "farid",
		Email:    "farid@gmail.com",
		Age:      32,
		IsActive: false,
	}

	fmt.Println(user)
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Age:", user.Age)
	fmt.Println("IS Active:", user.IsActive)

	var newUser User

	fmt.Println("Name:", newUser.Name)
	fmt.Println("Age:", newUser.Age)
	fmt.Println("IS Active:", newUser.IsActive)

	nUser := User{}

	fmt.Println(nUser.Age)

	cUser := user

	cUser.Name = "Chaged"
	fmt.Println(cUser.Name)
	fmt.Println(user.Name)

}
