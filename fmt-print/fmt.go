package main

import "fmt"

func main() {
	fmt.Print("faris fadhilah")    // Cetak tanpa tambahan karakter baris baru (\n)
	fmt.Print("faris", "fadhilah") // Sama dengan sebelumnya, tanpa spasi atau karakter tambahan

	fmt.Println("faris fadhilah")          // Cetak dengan karakter baris baru (\n) di akhir
	fmt.Println("faris", "fadhilah")       // Cetak dengan spasi antara "faris" dan "fadhilah" dan karakter baris baru di akhir
	fmt.Println("faris ", " ", "fadhilah") // Cetak dengan spasi ekstra antara "faris" dan "fadhilah" dan karakter baris baru di akhir
}
