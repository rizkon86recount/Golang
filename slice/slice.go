package main

import (
	"fmt"
)

/*
func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	fmt.Println(fruits[0])
} */

/*
// Hubungan Slice Dengan Array & Operasi Slice
func main() {
	// membuat slice dari array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // slice sekarang akan menjadi [2 3 4]

	// Menggunakan len dan cap untuk mengukur panjang dan kapasitas
	fmt.Println(len(slice)) // output: 4 (jumlah elemen dalam slice)
	fmt.Println(cap(slice)) // Output: 4 (kapasitas maksimum slice)

	// Menggunakan copy untuk menyalin elemen
	slice2 := make([]int, len(slice))
	copy(slice2, slice) // Menyalin elemen dari slice ke slice 2

	fmt.Println(slice2) // Output: [2 3 4 6]

	// Slicing lanjutan
	subSlice := slice[1:3] // Membuat slice baru dari slice yang sudah ada
	fmt.Println(subSlice)  // Output: [3 4]
} */

/*
// Slice Merupakan Tipe Data Reference
func main() {
	original := []int{1, 2, 3, 4, 5}
	slice := original[1:4]

	fmt.Println(slice) // Output: [2 3 4]

	// Mengubah elemen dalam slice
	slice[0] = 10
	fmt.Println(original) // Output: [1 10 3 4 5]
} */

/*
// len()array
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	length := len(arr) // length akan menjadi 5
	fmt.Println("arr", arr)
	fmt.Println("length", length)
} */

/*
// len()slice
func main() {
	slice := []int{1, 2, 3, 4, 5}
	length := len(slice) // length akan menjadi 5
	fmt.Println("slice", slice)
	fmt.Println("length", length)
} */

/*
// len()string
func main() {
	str := "Hello, World!"
	length := len(str) // length akan menjadi 13
	fmt.Println("str", str)
	fmt.Println("length", length)
} */

/*
// len()map
func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	length := len(m) // length akan menjadi 3
	fmt.Println("m", m)
	fmt.Println("length", length)
} */

/*
// len()channel
func main() {
	ch := make(chan int, 3) // Channel buffered dengan kapasitas 3
	length := len(ch)       // length akan menjadi 0 karena belum ada data dalam channel
	fmt.Println("ch", ch)
	fmt.Println("length", length)
} */

/*
//fungsi cap() slice
func main() {
	slice := make([]int, 3, 5) // Slice dengan panjang 3 dan kapasitas 5
	capacity := cap(slice)     // capacity akan menjadi 5
	fmt.Println("slice", slice)
	fmt.Println("capacity", capacity)
} */

/*
// fungsi cap() channel
func main() {
	ch := make(chan int, 3)	// Channel buffered dengan kapasitas 3
	capacity := cap(ch)		// capacity akan menjadi 3
	fmt.Println("ch", ch)
	fmt.Println("capacity", capacity)
} */

/*
//Contoh penggunaan kapasitas pada slice:
func main() {
	slice := make([]int, 3, 5)
	fmt.Println("Panjang:", len(slice))
	fmt.Println("Kapasitas:", cap(slice))

	// Menambahkan elemen ke dalam slice tanpa melebihi kapasitas
	slice = append(slice, 4, 5)
	fmt.Println("Panjang setelah append:", len(slice))
	fmt.Println("Kapasitas setelah append:", cap(slice))
} */

/*
// fungsi append() menambah elemen sice
func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	// Sekarang, slice menjadi [1 2 3 4 5]
	fmt.Println("slice", slice)
} */

/*
// fungsi append() menggabungkan slice
func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice1 = append(slice1, slice2...)
	// Sekarang, slice1 menjadi {1 2 3 4 5 6}
	fmt.Println("slice1", slice1)
} */

/*
// fungsi append() menghindari kapasitas
func main() {
	slice := make([]int, 3, 5)
	slice = append(slice, 4)
	// Sekarang, slice memiliki panjang 4 dan kapasitas tetap 5
	fmt.Println("slice", slice)
} */

/*
// fungsi append() mengalokasikan kapasitas tambahan
func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6, 7)
	// Saat elemen 4, 5, 6 dan 7 ditambahkan, Go mungkin akan mengalokasikan kapasitas tambahan
	fmt.Println("slice", slice)
} */

/*
// fungsi copy()
func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))

	// Menyalin elemen elemen dari src ke dst
	copied := copy(dst, src)

	fmt.Println("Elemen yang berhasil disalin:", copied)
	fmt.Println("dst:", dst)
} */

/*
// fungsi copy()
func main() {
	slice1 := []int{1, 2}
	slice2 := []int{3, 4}
	slice3 := []int{5, 6}
	dst := make([]int, len(slice1)+len(slice2)+len(slice3))

	offset := copy(dst, slice1)
	offset += copy(dst[offset:], slice2)
	offset += copy(dst[offset:], slice3)

	fmt.Println(dst)
} */

func main() {
	original := []int{1, 2, 3, 4, 5, 6}

	// Menggunakan 3 indeks untuk slicing
	slice := original[1:4:4]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
