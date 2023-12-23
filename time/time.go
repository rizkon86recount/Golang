/*
// Time (Waktu)
package main

import (
	"fmt"
	"time"
)

func main() {
	// Mendapatkan waktu saat ini
	currentTime := time.Now()

	fmt.Println("Waktu saat ini:", currentTime)
} */

/*
// Parsing Time (Menguraikan Waktu):
package main

import (
	"fmt"
	"time"
)

func main() {
	// String yang berisi informasi waktu
	timeStr := "2023-09-21T14:30:00Z"

	// Parsing string ke tipe time.Time
	parsedTime, err := time.Parse(time.RFC3339, timeStr)

	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	fmt.Println("Waktu yang diuraikan:", parsedTime)
}  */

// Format Time (Mengatur Format Waktu)
package main

import (
	"fmt"
	"time"
)

func main() {
	// Mendapatkan waktu saat ini
	currentTime := time.Now()

	// Mengonversi waktu menjadi string dengan format tertentu
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Println("Waktu yang diformat:", formattedTime)
}
