package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inisialisasi generator angka acak dengan seed berdasarkan waktu saat ini
	rand.Seed(time.Now().UnixNano())

	// Buat slice dengan angka dari 1 hingga 10
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Acak urutan angka dalam slice
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	// Cetak slice yang sudah diacak
	fmt.Println(numbers)
}
