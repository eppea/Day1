package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {

	merged := append(arrayA, arrayB...)

	uniqueMap := make(map[string]bool)

	mergedUnique := []string{}

	for _, val := range merged {

		if _, ok := uniqueMap[val]; !ok {
			mergedUnique = append(mergedUnique, val)
			uniqueMap[val] = true
		}
	}

	return mergedUnique
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alias", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alias", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
