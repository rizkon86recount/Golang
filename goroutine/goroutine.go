/*
// Penggunaan Goroutine Sederhana:
package main

import (
	"fmt"
	"time"
)

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar:", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	go foo() // Membuat goroutine untuk fungsi foo
	go bar() // Membuat goroutine untuk fungsi bar

	// Tunggu sebentar untuk memungkinkan goroutine berjalan
	time.Sleep(time.Second)
} */

/*
// penggunaan goroutine dengan channel
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Membuat lima goroutine pekerja
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	// Mengirim pekerjaan ke goroutine
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// Menutup channel pekerjaan setelah selesai
	close(jobs)

	// Mengumpulkan hasil
	for a := 1; a <= numJobs; a++ {
		<-results
	}
} */

/*
// Web Scraper Concurrent
package main

import (
	"fmt"
	"sync"
	"time"
)

func scrapeWebsite(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	// logika web scraping disini
	fmt.Println("Scraping", url)
	time.Sleep(time.Second)
}

func main() {
	urls := []string{"http://example.com", "https://google.com", "https://github.com"}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go scrapeWebsite(url, &wg)
	}

	wg.Wait()
	fmt.Println("Scrapping selesai.")
} */

/*
//Penggunaan Fungsi runtime.GOMAXPROCS()
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Mendapatkan jumlah CPU yang tersedia
	numCPU := runtime.NumCPU()
	fmt.Println("Number of CPUs:", numCPU)

	// Mendapatkan dan mencetak jumlah maksimum CPU yang diizinkan untuk goroutine
	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Println("Max Procs (default):", maxProcs)

	// Mengatur jumlah maksimum CPU yang diizinkan untuk goroutine
	newMaxProcs := 2
	oldMaxProcs := runtime.GOMAXPROCS(newMaxProcs)
	fmt.Println("Max Procs (changed):", oldMaxProcs)

	// Goroutine contoh
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is running\n", id)
		}(i)
	}
	wg.Wait()
} */

// Penggunaan Fungsi fmt.Scanln()
package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Masukan nama Anda:")
	fmt.Scanln(&name)

	fmt.Print("Masukan usia Anda:")
	fmt.Scanln(&age)

	fmt.Printf("Halo, %s! Usia Anda adalah %d tahun.\n", name, age)
}
