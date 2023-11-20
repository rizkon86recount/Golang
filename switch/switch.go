package main

import "fmt"

func main() {
	hari := "Rabu"
	switch hari {
	case "Senin":
		fmt.Println("Hari Senin")
	case "Sabtu":
		fmt.Println("Hari Sabtu")
	default:
		// kode yang dijalankan jika tidak ada kasus yang cocok
		fmt.Println("Hari lainnya")
	}
}
