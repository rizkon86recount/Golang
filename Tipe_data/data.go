package main

import "fmt"

/*
// Tipe Data Numerik Non-Desimal
func main() {
	var a int = 42
	var b uint = 24
	var c int8 = -8

	fmt.Println("Bilangan bulat (int):", a)
	fmt.Println("Bilangan bulat non-negatif (uint):", b)
	fmt.Println("Bilangan bulat 8-bit (int8):", c)
} */

/*
// Tipe Data Numerik Desimal
func main() {
	// Contoh penggunaan tipe data float 32
	var a float32 = 3.14
	var b float32 = -0.5

	fmt.Println("Nilai a (float32):", a)
	fmt.Println("Nilai b (float32):", b)

	// Contoh penggunaan tipe data float64
	var x float64 = 123.456789
	var y float64 = -9876.54321

	fmt.Println("Nilai x (float64):", x)
	fmt.Println("Nilai y (float64):", y)
} */

/*
// Tipe Data Boolean
func main() {
	var benar bool = true
	var salah bool = false

	fmt.Println("Benar:", benar)
	fmt.Println("Salah:", salah)

	// Contoh penggunaan tipe data bool dalam kondisional
	umur := 18
	bisaVote := umur >= 18

	fmt.Println("Bisa Memilih:", bisaVote)
} */

func main() {
	// Contoh penggunaan tipe data string
	nama := "Rizkon"
	pesan := "Halo, selamat datang di Haltev Academy"

	fmt.Println("Nama:", nama)
	fmt.Println("Pesan:", pesan)

	// Menggabungkan string
	sapaan := "Halo, "
	namaLengkap := "Rizkon"
	salam := sapaan + namaLengkap

	fmt.Println(salam)

	// Mengakses karakter dalam string
	hurufPertama := nama[0]
	fmt.Println("Huruf pertama:", string(hurufPertama))
}
