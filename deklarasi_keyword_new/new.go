package main

import "fmt"

func main() {
	// Deklarasi variabel pointer ke int
	var umur *int

	// Menggunakan 'new' untuk mengalokasikan memori
	umur = new(int)

	// Mengisi nilai variabel melalui pointer
	*umur = 21

	fmt.Println("Umur:", *umur) // Mencetak nilai melalui pointer
}
