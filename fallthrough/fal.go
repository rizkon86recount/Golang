package main

import "fmt"

func main() {
	num := 2

	switch num {
	case 1:
		fmt.Println("Ini adalah angka 1")
	case 2:
		fmt.Println("Ini adalah angka 2")
		fallthrough // mengizinkan eksekusi melanjutkan ke kasus berikutnya
	case 3:
		fmt.Println("Ini adalah angka 3")
	default:
		fmt.Println("Angka tidak dikenali")
	}
}
