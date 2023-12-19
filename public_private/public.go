// package mypackage

// // ini adalah properti public
// var publicVariable int

// // ini adalah fungsi public
// func publicFunction() {
// 	// ...
// }

// package mypackage
// // ini adalah properti private
// var privateVariable int

// // Ini adalah fungsi private
// func privateFunction() {
// 	// ...
// }

package main

import (
	"fmt"
)

func main() {
	// Ini merupakan properti public
	fmt.Println(mypackage.publicVariable)
	mypackage.publicFunction()

	// Ini akan menyebabkan kesalahan karena privateVariable dan privateFunction tidak dapat diakses dari luar mypackage
	// fmt.Println(mypackage.privateVariable)
	// mypackage.privateFunction()
}
