// penerapan interface
/*
package main

import (
	"fmt"
)

// Buat sebuah interface bernama 'shape' dengan satu metode 'Area() float64'
type Shape interface {
	Area() float64
}

// Definisikan struct 'Circle' dengan radius
type Circle struct {
	Radius float64
}

// Implementasikan metode 'Area()' untuk tipe data 'Circle'
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Definisikan struct 'Rectangle' dengan panjang dan lebar
type Rectangle struct {
	Length float64
	width  float64
}

// Impelementasikan metode 'Area()' untuk tipe data 'Rectangle'
func (r Rectangle) Area() float64 {
	return r.Length * r.width
}

func main() {
	// Buat objek 'Circle' dan 'Rectangle'
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Length: 4, width: 6}

	// Gunakan interface 'Shape' untuk menghitung luas objek-objek tersebut
	shapes := []Shape{circle, rectangle}

	for _, shape := range shapes {
		fmt.Printf("Luas: %f\n", shape.Area())
	}
}  */

// embeded interface
package main

import (
	"fmt"
)

// Definisikan sebuah interface 'writer' dengan metode 'Write()'
type Writer interface {
	Write(data []byte) (int, error)
}

// Definisikan interface 'Closer' dengan metode 'Close()'
type Closer interface {
	Close() error
}

// Definsikan tipe data 'File' yang mengimplementasikan 'Writer' dan 'Closer'
type File struct{}

// Implementasikan metode 'Write()' untuk 'File'
func (F *File) Write(data []byte) (int, error) {
	// Implementasi Write() di sini
	return len(data), nil
}

// Implementasikan metode 'Close()' untuk 'File'
func (f *File) close() error {
	// Implementasi Close() di sini
	return nil
}

func main() {
	file := &File{}
	data := []byte("Ini adalah contoh teks")

	// Menggunakan metode yang diimplementasikan oleh 'File'
	bytesWritten, _ := file.Write(data)
	fmt.Printf("Banyaknya byte yang ditulis: %d\n", bytesWritten)

	_ = file.close()
}
