package main

import (
	"fmt"
	"time"
)

func main() {
	// Menjadwalkan eksekusi fungsi foo setiap 5 detik
	schedular := time.NewTicker(5 * time.Second)

	go func() {
		for range schedular.C {
			foo()
		}
	}()
	<-time.After(30 * time.Second)
}

func foo() {
	fmt.Println("Eksekusi fungsi foo pada", time.Now())
}
