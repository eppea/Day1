package main

import "fmt"

func main() {
	var num int
	fmt.Print("Masukan bilangan prima: ")
	fmt.Scan(&num)

	if num <= 1 {
		fmt.Print("Bukan bilangan prima")
		return
	}

	isPrime := true
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(num, "adalah bilangan prima")
	} else {
		fmt.Println(num, "bukan bilangan prima")
	}
}
