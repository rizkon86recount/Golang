package main

import "fmt"

func main() {
	umur := 21
	isAdult := false

	if umur >= 18 {
		isAdult = true
	} else {
		isAdult = false
	}

	if isAdult {
		fmt.Println("Anda Dewasa")
	} else {
		fmt.Println("Anda belum Dewasa")
	}
}
