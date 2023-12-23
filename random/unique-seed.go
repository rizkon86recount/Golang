package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Insialisasi generator angka acak dengan seed berdasarkan waktu saat ini
	rand.Seed(time.Now().UnixNano())

	// Menghasilkan angka acak antara 1 dan 100
	randomNumber := rand.Intn(100)

	// Cetak angka acak
	fmt.Printf("Angka acak: %d\n", randomNumber)
}
