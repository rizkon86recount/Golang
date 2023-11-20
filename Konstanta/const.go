package main

import "fmt"

/*
func main() {
	const nama = "Rizkon"
	const pesanSalam = "Halo, "

	fmt.Print(pesanSalam + nama)
} */

// Deklarasi Multi Konstanta
func main() {
	const (
		namaDepan string  = "Rizkon"
		usia      int     = 21
		tinggi    float64 = 170.5
		mahasiswa bool    = true
	)

	fmt.Println("Nama Depan:", namaDepan)
	fmt.Println("Usia:", usia)
	fmt.Println("Tinggi Badan:", tinggi)
	fmt.Println("Mahasiswa:", mahasiswa)
}
