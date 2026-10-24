package main

import "fmt"

type Config struct {
	Port    int
	Debug   bool
	AppName string
}

func main() {
	age := 23
	ptr := &age

	fmt.Println("Value:", age)
	fmt.Println("Address:", &age)
	fmt.Println("Pointer:", ptr)
	fmt.Println("Value through pointer:", *ptr)

	score := 100

	// todo: store the address of score in a pointer
	sPtr := &score

	// todo: print score value
	fmt.Println("Score:", score)

	// todo: print socre memory address
	fmt.Println("Address:", &score)

	// todo print value through the pointer
	fmt.Println("Value through pointer:", *sPtr)

	// todo change score to 200 through pointer
	*sPtr = 133

	// todo: print the score again
	fmt.Println("New score:", score)
	fmt.Println("Value through pointer:", *sPtr)

	x := 10

	xPtr := &x

	*xPtr = 20

	fmt.Println("x:", x)
	fmt.Println("Value of x through pointer:", *xPtr)

	var nPtr *float64

	y := 12.11
	nPtr = &y

	fmt.Println("Pointer:", nPtr)
	fmt.Println("is nil:", nPtr == nil)
	fmt.Println("Value:", *nPtr)

	cfg := Config{
		Port:    9090,
		Debug:   true,
		AppName: "GoMastry",
	}

	cPtr := &cfg

	// todo: print the config through the pointer
	fmt.Println(*cPtr)

	// todo: change port to 8080 through the pointer
	cPtr.Port = 8080

	// todo: change debug to flase through pointer
	cPtr.Debug = false

	fmt.Println(*cPtr)

	var mAge int = 44
	var agePtr *int = &mAge

	*agePtr = 32
	fmt.Println(mAge)

	z := 12
	zPtr := &z

	change(zPtr)

	fmt.Println(z)
	fmt.Println(*zPtr)

	changev2(zPtr)
	fmt.Println(z)
	fmt.Println(*zPtr)

}

func change(ptr *int) {
	value := 1000
	ptr = &value
}

func changev2(ptr *int) {
	*ptr = 1000
}
