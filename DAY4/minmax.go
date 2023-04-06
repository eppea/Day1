package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]
	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Print("Enter a1: ")
	fmt.Scan(&a1)
	fmt.Print("Enter a2: ")
	fmt.Scan(&a2)
	fmt.Print("Enter a3: ")
	fmt.Scan(&a3)
	fmt.Print("Enter a4: ")
	fmt.Scan(&a4)
	fmt.Print("Enter a5: ")
	fmt.Scan(&a5)
	fmt.Print("Enter a6: ")
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Minimum number is:", min)
	fmt.Println("Maximum number is:", max)

	minPtr := &min
	maxPtr := &max
	fmt.Println("Using pointers:")
	fmt.Println("Minimum number is:", *minPtr)
	fmt.Println("Maximum number is:", *maxPtr)

	*minPtr = 99
	*maxPtr = 100
	fmt.Println("After modifying with pointers:")
	fmt.Println("Minimum number is:", min)
	fmt.Println("Maximum number is:", max)
}
