package main

import "fmt"

func tambah(a, b int) int {
	return a + b
}

func main() {
	total := tambah(5, 3)

	fmt.Println("Hasil penjumlahan:", total)
}
