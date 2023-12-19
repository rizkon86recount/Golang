/*
// Mencari Tipe Data & Value Menggunakan Reflect
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 42

	// Mencari tipe data dari variabel x
	t := reflect.TypeOf(x)
	fmt.Println("Tipe data:", t)

	// Mencari nilai dari variabel x
	v := reflect.ValueOf(x)
	fmt.Println("Nilai:", v.Int())
} */

/*
// pengaksesan nilai dalam bentuk interface{}
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 42

	// Mencari nilai dari variabel x dan mengonversinya menjadi interface{
	v := reflect.ValueOf(x)
	interfaceValue := v.Interface()

	// Menggunakan type assertion untuk mengakses nilai asli
	if intValue, ok := interfaceValue.(int); ok {
		fmt.Println("Nilai asli:", intValue)
	} else {
		fmt.Println("Gagal mengoversi ke int")
	}
} */

/*
// Pengaksesan Informasi Property Variabel Objek
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	person := Person{
		Name:  "John",
		Age:   30,
		Email: "john@xample.com",
	}

	// Menggunakan refleksi untuk mengakses nilai properti objek
	v := reflect.ValueOf(person)

	nameField := v.FieldByName("Name")
	ageField := v.FieldByName("Age")
	emailField := v.FieldByName("Email")

	if nameField.IsValid() {
		fmt.Println("Name:", nameField.Interface())
	}

	if ageField.IsValid() {
		fmt.Println("Age:", ageField.Interface())
	}

	if emailField.IsValid() {
		fmt.Println("Email:", emailField.Interface())
	}
} */

// Pengaksesan Informasi Method Variabel Objek
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p Person) GetAge() int {
	return p.Age
}

func main() {
	person := Person{
		Name: "John",
		Age:  30,
	}

	// Menggunakan refleksi untuk mengakses informasi method
	t := reflect.TypeOf(person)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println("Method:", method.Name)
		fmt.Println("Type:", method.Type)
	}
}
