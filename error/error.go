/*
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Masukkann sebuah nilai numerik: ")
	fmt.Scanln(&input)

	// Mengkonversu input ke tipe data float64
	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Input bukan nilai numerik:", err)
	} else {
		fmt.Printf("Input adalah nilai numerik: %.2f\n", num)
	}
}  */

/*
// Membuat custom error
package main

import (
	"errors"
	"fmt"
)

// Membuat custom error dengan tipe data ErrorNegativeNumber
type ErrorNegativeNumber int

// Impelemntasi metode Error() untuk custom error
func (e ErrorNegativeNumber) Error() string {
	return fmt.Sprintf("Nilai negatif tidak diperbolehkan: %d", int(e))
}

// Fungsi yang menggunakan custom error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagian oleh nol tidak diperbolehkan")
	}
	if a < 0 || b < 0 {
		return 0, ErrorNegativeNumber(a)
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil pembagian:", result)
	}

	result, err = divide(-5, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil pembagian:", result)
	}
} */

/*
// penggunaan panic
package main

import (
	"fmt"
)

func main() {
	divideByZero()
	fmt.Println("Program berjalan setelah panic")
}

func divideByZero() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic terjadi:", r)
		}
	}()

	num1, num2 := 10, 0
	result := num1 / num2 // Operasi ini akan menyebabkan panic
	fmt.Println("Hasil pembagian:", result)
}   */

/*
// Penggunaan recover
package main

import (
	"fmt"
)

func main() {
	defer handlePanic() // Menjadwalkan eksekusi handlepanic() saat panic erjadi
	divideByZero()
	fmt.Println("Program berjalan setelah panic")
}

func divideByZero() {
	num1, num2 := 10, 0
	result := num1 / num2 // Operasi ini akan menyebabkan panic
	fmt.Println("Hasil pembagian:", result)
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Panic terjadi:", r)
	}
}  */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program dimulai")

	// IIFE untuk menangani panic secara lokal
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic terjadi:", r)
			}
		}()

		// kode yang dapat menyebabkan panic
		num1, num2 := 10, 0
		result := num1 / num2
		fmt.Println("Hasil pembagian:", result)
	}()

	fmt.Println("Program selesai")
}
