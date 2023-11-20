package main

import "fmt"

/*
func main() {
	age := 21
	isStudent := true

	if age >= 18 {
		fmt.Println("Anda adalah seorang Dewasa")

		if isStudent {
			fmt.Println("Anda adalah seorang mahasiswa")
		} else {
			fmt.Println("Anda adalah seorang dewa yang bukan mahasiswa")
		}
	} else {
		fmt.Println("Anda adalah seorang anak-anak atau remaja")
	}
} */

func main() {
	day := "Monday"
	meal := "Lunch"

	switch day {
	case "Monday":
		fmt.Println("Hari Senin")
		switch meal {
		case "Breakfast":
			fmt.Println("Anda sedang saraoan pada hari senin")
		case "lunch":
			fmt.Println("Anda sedang makan siang pada hari Senin")
		default:
			fmt.Println("Anda sedang makan pada hari senin")
		}
	case "Tuesday":
		fmt.Println("Hari Selasa")
	default:
		fmt.Println("Hari lainnya")
	}
}
