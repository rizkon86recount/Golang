/*
// Funngsi variadik
package main

import "fmt"

func calculateSum(numbers ...float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
func main() {
	total := calculateSum(1.0, 2.0, 3.0, 4.0, 5.0)
	fmt.Println(total)
} */

//tidak ada parameter variadic
package main

import "fmt"

func sayHello(names ...string) {
	for _, name := range names {
		fmt.Println("Hello,", name)
	}
}

func main() {
	// Memanggil fungsi sayHello dengan beberapa nama
	sayHello("Alice", "Bob", "Charlie")
}
