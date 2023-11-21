// package main

// import "fmt"

/*
// Fungsi yang megembalikan hasil penjumlahan dua bilangan bulat
func tambah(a, b int) int {
	return a + b
}

func main() {
	// memanggil fungsi tambah dan menyimpan hasilnya dalam variabel total
	total := tambah(5, 3)

	// menggunakan nilai kembalian fungsi dalam ekspresi atau mencetaknya
	fmt.Println("Hasil penjumlahan:", total)
} */

//Penggunaan fungsi rand.new()
/*package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Insialisasi generator angka acak dengan sebuah seed
	// Menggunakan waktu saat ini sebagai seed agar hasil selalu berubah
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	// Menghasilkan beberapa angka acak
	for i := 0; i < 5; i++ {
		randomNumber := randomGenerator.Intn(100) // Menghasilkan angka acak antara 0 dan 99
		fmt.Println("Angka acak:", randomNumber)
	}
} */

// Import banyak package
/*package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"strings"
) */

// Deklarasi Parameter Bertipe Data Sama
/* package main

import "fmt"

// Fungsi yang menerima dua bilangan bulat dan mengembalikan hasil penjumlahannya
func tambah(a, b int) int {
	return a + b
}

func main() {
	hasil := tambah(5, 3)
	fmt.Println("Hasil penjumlahan:", hasil) // Output: Hasil penjumlahan: 8
} */

/*
// Deklarasi Parameter Bertipe Data Sama
package main

import "fmt"

// Fungsi yang menerima dua string dan mengembalikan hasil penggabungan keduanya
func gabung(kata1, kata2 string) string {
	return kata1 + " " + kata2
}

func main() {
	// Mendefinisikan variabel kata1 dan kata2
	kata1 := "Halo"
	kata2 := "Dunia"

	// Memanggil fungsi gabung
	hasil := gabung(kata1, kata2)

	// Menampilkan hasil penggabungan
	fmt.Println("Hasil penggabungan:", hasil)
} */

package main

import "fmt"

// fungsi untuk menghitung kuadrat suatu angka
func kuadrat(angka int) int {
	hasil := angka * angka
	return hasil // Menghentikan eksekusi dan mengembalikan hasil
}

func main() {
	x := 5
	hasilKuadrat := kuadrat(x)
	fmt.Println("Kuadrat dari", x, "adalah", hasilKuadrat)
}
