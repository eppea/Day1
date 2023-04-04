package main

import "fmt"

func main() {
	var num int
	fmt.Print("Masukan BilangAn: ")
	fmt.Scan(&num)

	if num%7 == 0 {
		fmt.Println(num, "adalah bilangan kelipatan 7")
	} else {
		fmt.Println(num, "bukan bilangan kelipatan 7")
	}
}
