package main

import "fmt"

func main() {
	umur := 15
	if umur >= 18 {
		fmt.Println("Anda Dewasa.")
	} else {
		// kode yang dijalankan jika kondisi salah
		fmt.Println("Anda belum Dewasa")
	}
}
