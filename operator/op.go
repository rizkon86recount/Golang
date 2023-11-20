package main

import "fmt"

/*
func main() {
	a := 10
	b := 5

	// Operasi Aritmatika
	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	modulus := a % b

	fmt.Println("Tambah:", tambah)
	fmt.Println("Kurang:", kurang)
	fmt.Println("kali:", kali)
	fmt.Println("Bagi:", bagi)
	fmt.Println("Modulus:", modulus)
} */

/*
func main() {
	a := 10
	b := 5

	// Operasi perbandingan
	samaDengan := a == b
	tidakSamaDengan := a != b
	kurangDari := a < b
	lebihDari := a > b
	kurangDariSamaDengan := a <= b
	lebihDariSamaDengan := a >= b

	fmt.Println("Sama dengan:", samaDengan)
	fmt.Println("Tidak sama dengan:", tidakSamaDengan)
	fmt.Println("Kurang dari:", kurangDari)
	fmt.Println("Lebih dari:", lebihDari)
	fmt.Println("Kurang dari atau sama dengan:", kurangDariSamaDengan)
	fmt.Println("Lebih dari atau sama dengan:", lebihDariSamaDengan)
} */

/*
func main() {
	x := true
	y := false

	// Operasi logika
	logikaAnd := x && y
	logikaOr := x || y
	nagasi := !x

	fmt.Println("Logika AND:", logikaAnd)
	fmt.Println("Logika OR:", logikaOr)
	fmt.Println("Nagasi:", nagasi)
} */

/*
func main() {
	a := 10
	b := 5

	// Operator penugasan
	a += b // a = a + b
	fmt.Println("a setelah ditambah b:", a)

	a -= b // a = a - b
	fmt.Println("a setelah dikurangi b:", a)

	a *= b // a = a * b
	fmt.Println("a setelah dikali b:", a)

	a /= b // a = a / b
	fmt.Println("a setelah dibagi b:", a)

	a %= b // a = a % b
	fmt.Println("Sisa pembagian a oleh b:", a)
} */

/*
func main() {
	a := 10

	// Operator inkrementasi dan dekrementasi
	a++ // Inkremen (a = a + 1)
	fmt.Println("Setelah inkremen:", a)

	a-- // Dekremen (a = a - 1)
	fmt.Println("Setelah dekremen:", a)
} */

func main() {
	a := 5
	b := 3

	// Operator bitwise
	danBitwise := a & b
	atauBitwise := a | b
	xorBitwise := a ^ b
	shiftKiri := a << 1  // Menggeser 1 bit ke kiri
	shiftKanan := a >> 1 // Menggeser 1 bit ke kanan

	fmt.Printf("AND Bitwise: %08b\n", danBitwise)
	fmt.Printf("OR Bitwise: %08b\n", atauBitwise)
	fmt.Printf("XOR Bitwise: %08b\n", xorBitwise)
	fmt.Println("Shift Kiri:", shiftKiri)
	fmt.Println("Shift Kanan:", shiftKanan)
}
