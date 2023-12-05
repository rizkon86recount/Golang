/*
// Funngsi variadik
package main

import "fmt"

func calculateSum(numbers ...float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
func main() {
	total := calculateSum(1.0, 2.0, 3.0, 4.0, 5.0)
	fmt.Println(total)
} */

/*
//tidak ada parameter variadic
package main

import "fmt"

func sayHello(names ...string) {
	for _, name := range names {
		fmt.Println("Hello,", name)
	}
}

func main() {
	// Memanggil fungsi sayHello dengan beberapa nama
	sayHello("Alice", "Bob", "Charlie")
} */

/*
// Penerapan Fungsi Variadic Menghitung Rata rata
package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	avg := average(2.0, 3.0, 4.0, 5.0)
	fmt.Println("Rata-rata:", avg)
} */

/*
// Meyambut beberapa orang
package main

import "fmt"

func sayHello(names ...string) {
	for _, name := range names {
		fmt.Println("Hello,", name)
	}
}

func main() {
	sayHello("Rizkon", "Faris", "Fajri")
} */

/*
// Menggabungkan Slice dengan Variadic Function
package main

import "fmt"

func concatenateStrings(separator string, strings ...string) string {
	result := ""
	for i, str := range strings {
		if i > 0 {
			result += separator
		}
		result += str
	}
	return result
}

func main() {
	combined := concatenateStrings(", ", "apple", "banana", "cherry")
	fmt.Println("Hasil penggabungan:", combined)
} */

/*
// Penggunaan Fungsi fmt.Sprintf()
package main

import (
	"fmt"
)

func main() {
	name := "Alice"
	age := 30
	height := 175.5

	// Menggunakann fmt.sprintf() untuk membuat dan memformat string
	formattedString := fmt.Sprintf("Nama: %s, Usia: %d tahun, Tinggi: %.2f cm", name, age, height)

	// Mencetak hasilnya
	fmt.Println(formattedString)
} */

/*
// Pengisian Parameter Fungsi Variadic Menggunakan Data Slice
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Mendefinisikan sebuah slice
	numSlice := []int{1, 2, 3, 4, 5}

	// Menggunakan operator ... untuk mengirim data slice ke fungsi variadic
	result := sum(numSlice...)

	fmt.Printf("Hasil penjumlahan: %d\n", result)
} */

/*
//Fungsi Dengan Parameter Biasa:
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	result := add(3, 5)
	fmt.Println("Hasil penjumlahan:", result)
} */

/*
//Fungsi Dengan Parameter Variadic:
package main

import (
	"fmt"
)

func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func main() {
	avg := average(2.0, 3.0, 4.0, 5.0)
	fmt.Println("Rata-rata:", avg)
} */

package main

import "fmt"

func example(a int, b ...int) {
	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %v\n", b)
}
func main() {
	example(1, 2, 3, 4, 5)
}
