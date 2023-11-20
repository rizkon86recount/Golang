package main

import "fmt"

func main() {
	umur := 15
	if umur >= 18 {
		fmt.Println("Anda Dewasa")
	} else if umur >= 13 {
		// kode yang dijalankan jika kondisi ini benar
		fmt.Println("Anda Remaja")
	} else {
		fmt.Println("Anda Anak-Anak")
	}
}
