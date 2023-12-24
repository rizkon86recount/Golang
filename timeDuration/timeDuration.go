/*
// Pembuatan durasi
import "time"

// Membuat durasi 2 detik
dur := 2 * time.Second

// Membuat durasi 30 menit
dur := 30 * time.Minute

*/

/*
   // Operasi Aritmatika
   dur1 := 2 * time.Second
   dur2 := 3 * time.Second

   // Penjumlahan
   sum := dur2 - dur1

   // Pengurangan
   diff := dur2 - dur1

   // Perkalian
   result := dur1 * 5

   // Pembagian
   division := dur2 / 4

*/

/*
   //konversi
   dur := 2 * time.Second

   // Konversi ke float64
   durInSeconds := dur.Second()

   // Konversi ke int64 (nanosekon)
   durInNanoseconds := dur.Nanoseconds()
*/

/*
// Mengukur Waktu
startTime := time.Now()
// Lakukan sesuatu yang memakan waktu
endTime := time.Now()

// Menghitung berapa lama waktu yang telah berlalu
elapsedTime := endTime.Sub(startTime)

*/

package main

import (
	"fmt"
	"time"
)

// Fungsi untuk menghitung faktorial dari sebuah bilangan
func factorial(n int) uint64 {
	result := uint64(1)
	for i := 1; i <= n; i++ {
		result *= uint64(i)
	}
	return result
}

func main() {
	// Catat waktu awal
	startTime := time.Now()

	// Hitung faktorial dari bilangan 20
	result := factorial(20)

	// Catat waktu akhir
	endTime := time.Now()

	// Hitung berapa lama waktu yang dibutuhkan
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("Hasil faktorial: %d\n", result)
	fmt.Printf("Waktu yang dibutuhkan: %s\n", elapsedTime)
}
