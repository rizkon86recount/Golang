package main

import "fmt"

func main() {
	defer fmt.Println("Ini akan dieksekusi terakhir")
	fmt.Println("Hallo")
	fmt.Println("Selamat datang")
}
