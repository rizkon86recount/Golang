package main

import "fmt"

/*
func main() {
	// Mendefinisikan sebuah array bilangan bulat dengan panjang 5
	var angka [5]int

	// Menginisialisasi
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	angka[3] = 40
	angka[4] = 50

	// Mengakses elemen-elemen array dan mencetaknya
	fmt.Println(angka[0])
	fmt.Println(angka[2])

	// Panjang array bisa didapatkan dengan menggunakan fungsi len()
	panjang := len(angka)
	fmt.Println("Panjang array angka adalah:", panjang)
} */

func main() {
	var angka [5]int // Mendeklarasikan array angka dengan panjang
	angka = [5]int{10, 20, 30, 40, 50}
	fmt.Println("angka", angka)
}
