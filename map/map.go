package main

import "fmt"

/*
func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 10
} */

/*
func main() {
	// Inisialisasi map dengan nilai awal
	myMap := map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println("myMap:", myMap)
} */

/*
func main() {
	// Insialisasi map dengan make
	myMap := make(map[string]int)

	// Menambahkan nilai ke dalam map
	myMap["januari"] = 50
	myMap["februari"] = 40

	fmt.Println("myMap:", myMap)
} */

/*
// menggunakkan nilai default
func main() {
	var myMap map[string]int // Map belum diinisialisasi

	// Nilai default
	fmt.Println(myMap["januari"]) // output: 0
} */

/*
//Iterasi Item Map Menggunakan for - range
func main() {
	// Contoh map
	chicken := map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    30,
	}

	// Iterasi item dalam map menggunakan for - range
	for key, value := range chicken {
		fmt.Printf("Bulan: %s, Jumlah Ayam: %d\n", key, value)
	}
} */

/*
//Menghapus item map
func main() {
	// contoh map
	chicken := map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    30,
	}

	// Menampilkan map sebelum penghapusan
	fmt.Println("Map sebelum penghapusan:")
	fmt.Println(chicken)

	// Menghapus item dari map
	delete(chicken, "februari")

	// Menampilkan map setelah penghapusan
	fmt.Println("\nMap setelah penghapusan:")
	fmt.Println(chicken)
} */

/*
//Deteksi Keberadaan Item Dengan Key Tertentu
func main() {
	// Contoh map
	chicken := map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    30,
	}

	// Kunci yang ingin anda cek
	keyToCheck := "februari"

	// Mendeteksi keberadaan item dengan kunci tertentu
	_, exist := chicken[keyToCheck]

	// mengecek apakah kunci tersebut ada dalam map
	if exist {
		fmt.Printf("%s ada dalam map, jumlah ayam: %d\n", keyToCheck, chicken[keyToCheck])
	} else {
		fmt.Printf("%s tidak ada dalam map\n", keyToCheck)
	}
} */

func main() {
	// membuuat slice untuk menyimpan nama-nama buah
	fruits := []string{"sprl", "pisang", "ceri"}
	// membuat map untuk menghubungkan nama buah dengan harga
	fruitPrices := map[string]int{
		"apel":   5,
		"pisang": 3,
		"ceri":   10,
	}

	// Menampilkan daftar buah beserta harganya
	fmt.Println("Daftar Buah dan harga:")
	for _, fruit := range fruits {
		price, exists := fruitPrices[fruit]
		if exists {
			fmt.Printf("%s: Rp%d per kg\n", fruit, price)
		} else {
			fmt.Printf("%s: Harga tidak tersedia\n", fruit)
		}
	}
}
