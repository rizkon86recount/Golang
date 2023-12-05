/*
//Pointer untuk Struktur (Struct):
package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func main() {
	var person1 Person
	person1.Age = 30
	person1.Name = "John"

	var personPtr *Person
	personPtr = &person1

	fmt.Println(personPtr.Age)
} */

/*
// penerapan pointer
package main

import "fmt"

func main() {
	// Deklarasi variabel integer
	var number1 int = 42
	fmt.Println("value of number1:", number1)

	// Mendeklarasikan pointer yang menunjuk ke variabel number1
	var pointerToNumber1 *int
	pointerToNumber1 = &number1

	// Mengakses nilai yang ditunjuk oleh pointer
	fmt.Println("Value of number1 via pointer:", *pointerToNumber1)

	// Mengubah nilai melalui pointer
	*pointerToNumber1 = 100
	fmt.Println("Value of number1 after modification via pointer:", number1)
} */

/*
package main

import "fmt"

func main() {
	// Deklarasi variabel integer
	var number1 int = 42
	fmt.Println("Value of number1:", number1)

	// Mendeklarasikan pointer yang menunjuk ke variabel number1
	var pointerToNumber1 *int
	pointerToNumber1 = &number1

	// Mengubah nilai melalui pointer
	*pointerToNumber1 = 100

	// Mencetak nilai variabel asli
	fmt.Println("Value of number1 after modificationn via pointer:", number1)
} */

/*
// Inisialisasi Langsung Dengan Nilai-nilai Bidang (Field):
package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	person1 := Person{
		Name:    "John",
		Age:     30,
		Address: "123 Main St",
	}

	person2 := Person{
		Name:    "Alice",
		Age:     25,
		Address: "456 Elm St",
	}

	fmt.Println(person1)
	fmt.Println(person2)
} */

/*
//Inisialisasi Dengan Nilai-nilai Default
package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var person1 Person
	fmt.Println(person1)
} */

/*
// Inisialisasi Tanpa Nama Bidang
package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	person1 := Person{"John", 30, "123 Main St"}

	fmt.Println(person1)
} */

/*
// Penggunaan Pointer dalam Fungsi:
package main

import "fmt"

func doubleValue(x *int) {
	*x *= 2
}

func main() {
	y := 5
	doubleValue(&y)
	fmt.Println(y)
} */

/*
// Embeded Struct
package main

import (
	"fmt"
)

// struct yang akan disematkan (embedded)
type Person struct {
	Name string
	Age  int
}

// Struct yang menggunakan Person sebagai embedded struct
type Employee struct {
	Person     // Struct Person disematkan di sini
	EmployeeID int
}

func main() {
	// Membuat objek Employee
	employee := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		EmployeeID: 12345,
	}

	// Mengakses field dari struct Employee dan person
	fmt.Println("Name:", employee.Name)
	fmt.Println("Age:", employee.Age)
	fmt.Println("EmployeeID", employee.EmployeeID)
} */

/*
// Embedded Struct Dengan Nama Property Yang Sama
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Salary int
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		Salary: 50000,
	}

	fmt.Println("Nama:", emp.Name)
	fmt.Println("Usia:", emp.Age)
	fmt.Println("Gaji:", emp.Salary)
} */

/*
//Pengisian Nilai Sub-Struct
package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Person struct {
	Name    string
	Age     string
	Address Address
}

func main() {
	addr := Address{
		Street:  "123 Main St",
		City:    "Exampleville",
		ZipCode: "12345",
	}

	person := Person{
		Name:    "John Doe",
		Age:     "30",
		Address: addr, // Mengisi nilai properti Address dengan struktur address yang sudah dibuat
	}

	fmt.Println("Nama:", person.Name)
	fmt.Println("Usia:", person.Age)
	fmt.Println("Alamat:", person.Address.Street)
	fmt.Println("Kota:", person.Address.City)
	fmt.Println("Kode Pos:", person.Address.ZipCode)
} */

/*
// Anonymous struck
package main

import "fmt"

func main() {
	// Membuat objek anonymous struct dengan properti Name dan Age
	person := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	}

	fmt.Println("Nama:", person.Name)
	fmt.Println("Usia:", person.Age)
} */

/*
// Kombinasi Slice & Struct
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Membuat sebuah slice yang berisi objek-objek Person
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}

	// Menambahkan data baru ke dalam slice
	newperson := Person{Name: "Eve", Age: 28}
	people = append(people, newperson)

	// Mengakses dan mencetak data dalam slice
	for _, person := range people {
		fmt.Println("Nama:", person.Name)
		fmt.Println("Usia:", person.Age)
	}
} */

/*
// Inisialisasi Slice Anonymous Struct
package main

import "fmt"

func main() {
	// Inisialisasi slice yang berisi objek anonymous struct
	people := []struct {
		Name string
		Age  int
	}{
		{Name: "John", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}

	// Menambahkan data baru ke dalam slice
	newPerson := struct {
		Name string
		Age  int
	}{Name: "Eve", Age: 28}
	people = append(people, newPerson)

	// Mengakses dan mencetak data dalam slice
	for _, person := range people {
		fmt.Println("Nama:", person.Name)
		fmt.Println("Usia:", person.Age)
	}
} */

/*
// Deklarasi Anonymous Struct Menggunakan Keyword var
package main

import "fmt"

func main() {
	// Deklarasi variabel dangan tipe anonymous struct menggunakan var
	var person struct {
		Name string
		Age  int
	}

	// Insialisasi nilai pada variabel
	person.Name = "John"
	person.Age = 30

	// Mengakses dan mencetak nilai
	fmt.Println("Nama:", person.Name)
	fmt.Println("Usia:", person.Age)
} */

/*
// Nested struct
package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Person struct {
	Name string
	Age  int
	Home Address //properti Home adalah struktur Address
}

func main() {
	// Membuat objek Address
	homeAddress := Address{
		Street:  "123 Main St",
		City:    "Exampleville",
		ZipCode: "12345",
	}

	// Membuat objek Person dengan properti Address
	person := Person{
		Name: "John Doe",
		Age:  30,
		Home: homeAddress, // Mengisi nilai properti Home dengan objek Address
	}

	// Mengakses dan mencetak nilai-nilai dalam objek Personn dan Address
	fmt.Println("Nama:", person.Name)
	fmt.Println("Usia:", person.Age)
	fmt.Println("Alamat Rumah:", person.Home.Street)
	fmt.Println("Kota:", person.Home.City)
	fmt.Println("Kode Pos:", person.Home.ZipCode)
} */

/*
// Deklarasi Dan Inisialisasi Struct Secara Horizontal
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Deklarasi dan inisialisasi struktur secara horizontal
	john := Person{"John Doe", 30}
	alice := Person{"Alice Smith", 25}

	// Mengakses dan mencetak nilai dari objek struktur
	fmt.Println("Nama John:", john.Name)
	fmt.Println("Usia John:", john.Age)

	fmt.Println("Nama Alice:", alice.Name)
	fmt.Println("Usia Alice:", alice.Age)
} */

/*
// Tag property dalam struct
type Person struck {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
} */

// Type Alias
package main

import "fmt"

// Membuat type alias untuk tipe data yang sudah ada
type MyInt int

func main() {
	var num MyInt
	num = 42

	fmt.Println(num)
}
