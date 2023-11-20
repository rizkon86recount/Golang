package main

import "fmt"

/*
func getInfo() (string, int) {
	return "Rizkon", 21
}

func main() {
	nama, _ := getInfo()
	fmt.Println("Nama:", nama)
	// _ digunakan untuk mengabaikan nilai kedua yang dikembalikan oleh getInfo()
} */

// looping
func main() {
	data := []int{1, 2, 3, 4, 5}
	total := 0

	for _, nilai := range data {
		total += nilai
	}

	fmt.Println("Total:", total)
}
