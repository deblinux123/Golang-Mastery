package main

import "fmt"

func main() {
	var user map[string]string

	fmt.Println("user:", user)

	// user["farid"] = "199"

	fmt.Println(user)

	score := make(map[string]int)

	score["farid"] = 100

	fmt.Println("Score:", score)

	users := map[string]any{
		"name": "farid",
		"age":  32,
		"role": "ai engineer",
	}

	fmt.Println("user:", users)

	// todo: create a map
	student := make(map[string]int)

	// todo: add three student
	student["ali"] = 80
	student["farid"] = 90
	student["asma"] = 100

	// todo: print the first student grade
	fmt.Println(student["ali"])

	// todo: print the number of student
	fmt.Println(len(student))
}
