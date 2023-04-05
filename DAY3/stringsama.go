package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
	mapping := make(map[string]int)
	for _, value := range slice {
		mapping[value]++
	}
	return mapping
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
}
