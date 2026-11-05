package main

import (
	"fmt"
	"strings"
)

func sliceExample(numbers []int) {
	part := numbers[:2]

	fmt.Println("Original:", numbers)
	fmt.Println("Part:", part)

	fmt.Println("Len:", len(numbers))
}

func main() {

	numbers := []int{10, 20, 30, 40, 50}

	sliceExample(numbers)

	first := numbers[:2]
	second := numbers[1:3]

	second[0] = 44

	fmt.Println("Original:", numbers)
	fmt.Println("First:", first)
	fmt.Println("Second:", second)

	copyNumbers := make([]int, len(numbers))
	copy(copyNumbers, numbers)

	copyNumbers[0] = 99

	fmt.Println("Original:", numbers)
	fmt.Println("Copy:", copyNumbers)

	score := make([]int, 2, 4)

	score[0] = 10
	score[1] = 20

	partScore := score[:2]

	partScore = append(partScore, 30)

	fmt.Println("Score:", score)
	fmt.Println("partScore:", partScore)

	fmt.Println(strings.Repeat("=", 20))
	org := []int{10, 20, 30, 40, 50}
	partOrg := org[1:3:3]

	fmt.Println("Original:", org)
	fmt.Println("Length:", len(partOrg))
	fmt.Println("Capacity:", cap(partOrg))

	partOrg = append(partOrg, 999)

	fmt.Println("Original:", org)
	fmt.Println("PartOrg:", partOrg)
}
