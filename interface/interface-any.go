/*package main

import "fmt"

func main() {
	// Menggunakan interface kosong untuk menyimpan berbagai tipe data
	var data interface{}

	data = 42
	fmt.Println(data) // Output: 42

	data = "Hello, World!"
	fmt.Println(data)

	data = []int{1, 2, 3, 4, 5}
	fmt.Println(data)
} */

/*
// kosong (empty interface)
package main

import "fmt"

func main() {
	//Menggunakan interface kosong untuk menyimpan berbagai tipe data
	var data interface{}

	data = 42

	// memeriksa apakah data merupakan int
	value, ok := data.(int)
	if ok {
		fmt.Println("Nilai data adalah sebuah angka:", value)
	} else {
		fmt.Println("Data bukan angka.")
	}

} */

/*
// type assertion pada interface kosong
package main

import "fmt"

func main() {
	var data interface{}
	data = 42

	// Melakukan type assertion untuk mengubah data menjadi int
	angka, sukses := data.(int)
	if sukses {
		fmt.Printf("Nilai data adalah angka: %d\n", angka)
	} else {
		fmt.Println("Type assertion gagal.")
	}

	// Melakukan type assertion untuk mengubah data menjadi string
	str, sukses := data.(string)
	if sukses {
		fmt.Printf("Nilai data adalah string: %s\n", str)
	} else {
		fmt.Println("Type assertion gagal.")
	}
} */

// Kombinasi Slice, map, dan interface{}
package main

import (
	"fmt"
)

func main() {
	// Membuat slice dari interface{}
	data := []interface{}{}

	// Menambahkan berbagai jenis data ke dalam slice
	data = append(data, 42)
	data = append(data, "Hello, World!")
	data = append(data, map[string]int{"one": 1, "two": 2})
	data = append(data, []float64{1.1, 2.2, 3.3})

	// Mengakses data dalam slice dan melakukan type assertion
	for _, item := range data {
		switch value := item.(type) {
		case int:
			fmt.Printf("Ini adalah angka: %d\n", value)
		case string:
			fmt.Printf("Ini adalah string: %s\n", value)
		case map[string]int:
			fmt.Printf("Ini adalah peta: %+v\n", value)
		case []float64:
			fmt.Printf("Ini adalah slice float64: %+v\n", value)
		default:
			fmt.Println("Tipe data tidak dikenali.")
		}
	}
}
