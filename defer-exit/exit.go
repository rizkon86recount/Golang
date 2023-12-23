/*
package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Ini tidak akan dieksekusi")
	fmt.Println("Halo")
	fmt.Println("Selamat datang")
	os.Exit(1)
}  */

/*
// Kombinasi defer dan IIFE
package main

import "fmt"

func main() {
	fmt.Println("Main function started")

	// IIFE (Immediately Invokad Function Espression) yang menggunakan defer
	func() {
		fmt.Println("IFEE started")
		defer fmt.Println("IIFE defered")
		fmt.Println("IIFE ended")
	}()
	fmt.Println("Main function ended")
}  */

// os.Exit()
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program disini")

	// Contoh: Menghentikan program dengan kode keluar 0 (sukses)
	os.Exit(1)

	// Baris kode di bawah ini tidak akan dieksekusi
	fmt.Println("Baris ini tidak akan dieksekusi")
}
