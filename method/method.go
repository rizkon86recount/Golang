/*
// penerapan method
package main

import "fmt"

// Definisikan tipe data struct 'mobil'
type Mobil struct {
	Merk      string
	Tahun     int
	Kecepatan int
}

// Method 'TambahKecepatan' terkait dengan tipe data 'Mobil'
func (m *Mobil) TambahKecepatan(tambahan int) {
	m.Kecepatan += tambahan
}

// Method 'kurangiKecepatan' terkait dengan tipe data 'Mobil'
func (m *Mobil) kurangiKecepatan(pengurangan int) {
	m.Kecepatan -= pengurangan
}

func main() {
	// Membuat instance 'mobilSaya' dari tipe data 'Mobil'
	mobilSaya := Mobil{
		Merk:      "Toyota",
		Tahun:     2020,
		Kecepatan: 0,
	}

	// Memanggil method 'tambahKecepatan' untuk menambah kecepatan mobil
	mobilSaya.TambahKecepatan(20)
	fmt.Println("Kecepatan Mobil:", mobilSaya.Kecepatan)

	// Memanggil method 'kurangiKecepatan' untuk mengurangi kecepatan mobil
	mobilSaya.kurangiKecepatan(10)
	fmt.Println("Kecepatan Mobil:", mobilSaya.Kecepatan)
} */

// Method pointer
package main

import "fmt"

type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

// Method dengan receiver pointer untuk menghitung luas
func (p *PersegiPanjang) HitungLuas() float64 {
	return p.Panjang * p.Lebar
}

// Method dengan receiver pointer untuk mengubah panjang
func (p *PersegiPanjang) UbahPanjang(panjangBaru float64) {
	p.Panjang = panjangBaru
}

func main() {
	// Membuat instance 'persegiPanjang' dari tipe data 'PersegiPanjang'
	PersegiPanjang := PersegiPanjang{
		Panjang: 5.0,
		Lebar:   3.0,
	}

	// Memanggil method 'HitungLuas' dengan receiver pointer
	luas := PersegiPanjang.HitungLuas()
	fmt.Println("Luas Persegi Panjang:", luas)

	// Memanggil method 'UbahPanjang' dengan receiver pointer
	PersegiPanjang.UbahPanjang(7.0)
	fmt.Println("Panjang persegi Panjang:", PersegiPanjang.Panjang)
}
