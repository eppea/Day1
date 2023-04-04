package main

import "fmt"

func main() {
	var sisiAtas, sisiBawah, tinggi float64
	fmt.Print("Masukkan sisi atas: ")
	fmt.Scan(&sisiAtas)
	fmt.Print("Masukkan sisi bawah: ")
	fmt.Scan(&sisiBawah)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	luas := (sisiAtas + sisiBawah) / 2 * tinggi
	fmt.Println("Luas trapesium:", luas)
}
