/*
// Penerapan Fungsi Multiple Return
package main

import (
	"fmt"
)

// Fungsi ini mengembalikan dua nilai: hasil penjumlahan dan hasil perkalian dua bilangan
func tambahDanaKali(a, b int) (int, int) {
	jumlah := a + b
	kali := a * b
	return jumlah, kali
}

func main() {
	x, y := 5, 3
	hasilPenjumlahan, hasilPerkalian := tambahDanaKali(x, y)
	fmt.Printf("Hasil penjumlahan: %d\n", hasilPenjumlahan)
	fmt.Printf("Hasil perkalian: %d\n", hasilPerkalian)
} */

// Fungsi Dengan Predefined Return Value
package main

import "fmt"

// Fungsi dengan predefined return value
func tambah(a, b int) int {
	hasil := a + b
	return hasil
}

func main() {
	// Memanggil fungsi tambah dan menyimpan hasilnya dalam variabel
	hasilPenambahan := tambah(5, 3)

	// Mencetak hasil penambahan
	fmt.Println("Hasil penambahan:", hasilPenambahan)
}
