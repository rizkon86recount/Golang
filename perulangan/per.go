package main

import (
	"fmt"
)

/*
func main() {
	// Perulangan untuk mencetak angka dari 1 hingga 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Perulangan untuk mencetak angka dari 5 hingga 1 denga langkah mundur
	for j := 5; j >= 1; j-- {
		fmt.Println(j)
	}
} */

/*
func main() {
	// Menggunakan perulangan for dengan hanya kondisi
	count := 1
	for count <= 5 {
		fmt.Println(count)
		count++
	}
} */

/*
func main() {
	// Perulangan for tanpa argumen (infinite loop)
	for {
		fmt.Println("Ini adalah perulangan tak berujung")
	}
} */

/*
func main() {
	count := 1
	for {
		fmt.Println("Ini adalah perulangan tak berujung")
		if count >= 5 {
			break // keluar dari perulangan setelah count mencapai 5
		}
		count++
	}
} */

/*
func main() {
	// Contoh perulangan menggunakan range pada slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// contoh perulangan menggunakan range pada map
	person := map[string]string{
		"Nama": "Rizkon",
		"Umur": "21",
		"Kota": "Pemalang",
	}
	for key, value := range person {
		fmt.Printf("key: %s, Value: %s\n", key, value)
	}
}  */

/*
func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // Menghentikan perulangan i = 5
		}
		fmt.Println(i)
	}
} */

/*
func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // Melompati iterasi ketika 1 = 3
		}
		fmt.Println(i)
	}
} */

/*
func main() {
	// Perulangan bersarang untuk mencetak pola bintang
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" * ")
		}
		fmt.Println()
	}
} */

func main() {
outerLoop:
	for i := 1; i <= 3; i++ {
		fmt.Printf("Iterasi luar: %d\n", i)

		for j := 1; j <= 3; j++ {
			fmt.Printf("Iterasi dalam: %d\n", j)

			if i == 2 && j == 2 {
				break outerLoop // Menghentikan perulangan luar
			}
		}
	}
}
