package main

import (
	"fmt"
	"strconv"
)

func munculsekali(angka string) []int {
	count := make(map[int]int)
	for _, digit := range angka {
		num, _ := strconv.Atoi(string(digit))
		count[num]++
	}

	unique := []int{}
	for num, cnt := range count {
		if cnt == 1 {
			unique = append(unique, num)
		}
	}

	return unique
}

func main() {
	fmt.Println(munculsekali("1234123"))
	fmt.Println(munculsekali("76523752"))
	fmt.Println(munculsekali("12345"))
	fmt.Println(munculsekali("1122334455"))
	fmt.Println(munculsekali("8872504"))
}
