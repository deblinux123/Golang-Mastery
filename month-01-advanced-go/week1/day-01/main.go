package main

import "fmt"

func zeroValue() {
	var age int
	var name string
	var active bool
	var scores []float64
	var metadata map[string]string
	var ptr *int

	fmt.Printf("age: %d\n", age)
	fmt.Printf("name: %q\n", name)
	fmt.Printf("active: %t\n", active)
	fmt.Printf("scores: %#v\n", scores)
	fmt.Printf("metadata: %#v\n", metadata)
	fmt.Printf("ptr: %v\n", ptr)
}

func main() {

	// explicit type
	var age int = 33

	// tpe interface
	var language = "golang"

	// short variable declaration
	name := "farid"

	// declare first, assign later
	var score int
	score = 100

	fmt.Printf("Name: %q\nAge: %d\nLanguage: %q\nScore: %d\n", name, age, language, score)

	var latter rune = 'a'
	var b byte = 'a'

	fmt.Println(latter+1, string(b))

	var first uint32 = 12
	second := 32

	fmt.Println(first + uint32(second))

	fmt.Println(int(first) + second)

	fmt.Println("Zero value:")
	zeroValue()

	var user map[string]int

	fmt.Println(user["farid"])

	user = make(map[string]int)

	user["farid"] = 33

	fmt.Println(user)
}
