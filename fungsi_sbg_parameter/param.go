// package main

// import "fmt"

/*
// Callback Functions (Fungsi Panggilan Kembali):
func Hitung(a, b int, callback func(int, int) int) int {
	hasil := callback(a, b)
	return hasil
}

func penjumlahan(x, y int) int {
	return x + y
}

func Pengurangan(x, y int) int {
	return x - y
}

func main() {
	hasilTambah := Hitung(5, 3, penjumlahan)
	hasilKurang := Hitung(8, 4, Pengurangan)

	fmt.Println("Hasil Penjumlahan:", hasilTambah)
	fmt.Println("Hasil Pengurangan:", hasilKurang)
} */

/*
// Iterasi Koleksi
func IterasiSlice(slice []int, callback func(int)) {
	for _, elem := range slice {
		callback(elem)
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}

	IterasiSlice(data, func(elem int) {
		fmt.Println(elem)
	})
} */

/*
// Validasi input
package main

import (
	"fmt"
	"regexp"
)

func validasiInput(input string, validasi func(string) bool) bool {
	return validasi(input)
}

func validasiEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

func main() {
	inputEmail := "haltev@gmail.com"
	isValid := validasiInput(inputEmail, validasiEmail)
	fmt.Println("Email valid:", isValid)
} */

/*
// Alias Skeman Closure
package main

import "fmt"

func main() {
	x := 10

	// fungsi anonim (closure) yang memiliki akses ke variabel x
	closure := func() {
		fmt.Println(x)
	}

	// Memanggil closure
	closure()

	// Mengubah nilai x
	x = 20

	// Memanggil closure lagi
	closure()
} */

// Penggunaan Fungsi string.Contains()
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Halo, ini adalah contoh penggunaan string.Contains() dalam Go"
	substr1 := "contoh"
	substr2 := "Golang"

	// Memeriksa apakah str menganndung substr1
	if strings.Contains(str, substr1) {
		fmt.Printf("'%s' ditemukan dalam string\n", substr1)
	} else {
		fmt.Printf("'%s' tidak ditemukan dalam string\n", substr1)
	}

	// Memeriksa apakah str mengandung substr2
	if strings.Contains(str, substr2) {
		fmt.Printf("'%s' ditemukan dalam string\n", substr2)
	} else {
		fmt.Printf("'%s' tidak ditemukan dalam string\n", substr2)
	}
}
