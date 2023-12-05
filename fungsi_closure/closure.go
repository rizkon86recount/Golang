package main

import (
	"fmt"
)

/*
// Closure Disimpan Sebagai Variabel
func main() {
	// Definisikan sebuah closure dan simpan dalam variabel
	closure := func(x, y int) int {
		return x + y
	}

	// panggil closure dengan menggunakan variabel
	result := closure(10, 5)
	fmt.Println(result)
} */

/*
// Penggunaan Template String %v
func main() {
	name := "John"
	age := 30

	// Menggunakan %v untuk memasukkan nilai variabel ke dalam string
	fmt.Printf("Nama: %v, Umur: %v\n", name, age)
} */

/*
func main() {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("original number :", newNumbers)
} */

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := createCounter()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
