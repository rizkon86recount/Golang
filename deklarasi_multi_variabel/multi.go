package main

import "fmt"

/* func main() {
	var x, y, z int // Deklarasi tiga variabel bertipe int
	x = 1
	y = 2
	z = 3

	fmt.Println("Nilai x:", x)
	fmt.Println("Nilai y:", y)
	fmt.Println("Nilai z:", z)
} */

/*
// Deklarasi Multi Variabel dengan Tipe Data yang Berbeda
func main() {
	var nama string
	var umur int
	var tinggi float64
	nama, umur, tinggi = "Rizkon", 21, 170.5

	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("Tinggi:", tinggi)
} */

/*
// Deklarasi dan Inisialisasi Multi Variabel dengan Tipe Data yang Sama:
func main() {
	var a, b, c int = 1, 2, 3 // Deklaraso dan insialisasi tiga variabel bertipe int
	fmt.Println("Nilai a:", a)
	fmt.Println("Nilai b:", b)
	fmt.Println("Nilai c:", c)
} */

func main() {
	nama, umur, tinggi := "Rizkon", 30, 170.5 // Deklarasi dan inisialisasi tiga variabel tanpa meyebutkan tipe data
	fmt.Println("Umur:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("Tinggi:", tinggi)
}
