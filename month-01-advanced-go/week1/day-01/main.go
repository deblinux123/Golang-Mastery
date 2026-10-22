package main

import "fmt"

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
}
