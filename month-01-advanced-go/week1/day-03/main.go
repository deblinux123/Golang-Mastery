package main

import "fmt"

func main() {

	var numbers [3]int

	numbers[0] = 12
	numbers[1] = 22
	numbers[2] = 33

	fmt.Println(numbers)

	names := [3]string{"farid", "ahmad", "ali"}

	fmt.Println(names)

	scores := [...]int{10, 20, 30, 40}

	fmt.Println(scores)

	copyScores := scores

	fmt.Println(copyScores)
	copyScores[2] = 10000
	fmt.Println(copyScores)

	// slice
	ports := []int{80, 88, 443, 8080, 2345}

	for _, port := range ports {
		fmt.Println(port)
	}

	fmt.Printf("number of ports %d\n", len(ports))
	fmt.Printf("Cap: %d\n", cap(ports))

	ports = append(ports, 6379)

	fmt.Println(ports)

	original := []int{10, 20, 30, 40, 50}
	part := original[:3]

	part[2] = 99

	fmt.Println("original:", original)
	fmt.Println("part:", part)
}
