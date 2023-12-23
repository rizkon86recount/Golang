/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Menghasilkan float64 acak antara 0.0 dan 1.0
	randomFloat := rand.Float64()
	fmt.Printf("Angka float64 acak: %f\n", randomFloat)
} */

/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Menghasilkan int64 acak antara 0 dan 100 (inklusif)
	randomInt := rand.Int63n(101)
	fmt.Printf("Angka int64 acak: %d\n", randomInt)
} */

/*
// Angka random index tertentu
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inisialisasi generator angka acak dengan seed berdasarkan waktu saat ini
	rand.Seed(time.Now().UnixNano())

	// Buat sebuah slice aatau array data
	data := []string{"item1", "item2", "item3", "item4", "item5"}

	// Menghasilkan indeks acak
	randomIndex := rand.Intn(len(data))

	// Mendapatkan item dari slice dengan indeks yang dihasilkan secara acak
	randomItem := data[randomIndex]

	// Cetak hasil
	fmt.Printf("Indeks acak:%d\n", randomIndex)
	fmt.Printf("Item acak: %s\n", randomItem)
}  */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inisialisasi generator angka acak dengan seed berdasarkan waktu saat ini
	rand.Seed(time.Now().UnixNano())

	// Buat slice yang berisi pilihan-pilihan string
	pilihan := []string{"Apel", "Jeruk", "pisang", "Mangga", "Strawberi"}

	// Menghasilkan indeks acak untuk memilih string dari slice
	randomIndex := rand.Intn(len(pilihan))

	// Memilih string dari slice berdasarkamn indeks yang dihasilkan secara acak
	randomString := pilihan[randomIndex]

	// Cetak hasil
	fmt.Printf("string acak: %s\n", randomString)
}
