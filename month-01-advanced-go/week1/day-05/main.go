package main

import (
	"fmt"
	"sort"
)

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

	// look up the map
	grade, exists := student["reza"]

	if exists {
		fmt.Println("Grade:", grade)
	} else {
		fmt.Println("Student Not Found.")
	}

	delete(student, "reza")

	product := map[string]any{
		"name":   "phone",
		"price":  332,
		"region": "IR",
	}

	for key, value := range product {
		fmt.Println(key, ":", value)
	}

	scores := map[string]int{
		"Sara":  92,
		"Farid": 95,
		"Ali":   88,
	}

	keys := make([]string, 0, len(scores))

	for key := range scores {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key)
	}
}
