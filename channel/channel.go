/*
// Penggunaan Channel untuk Menghitung Nilai Rata-Rata
package main

import (
	"fmt"
)

func calculateAverage(numbers []int, result chan float64) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))
	result <- average
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	resultChannel := make(chan float64)

	go calculateAverage(numbers, resultChannel)

	average := <-resultChannel

	fmt.Printf("Rata-rata: %.2f\n", average)
} */

/*
// Penggunaan Channel untuk Menyinkronkan Dua Goroutine
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Pekerja %d memproses pekerjaan %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Hasil: %d\n", result)
	}
} */

/*
// Channel Sebagai Tipe Data Parameter
package main

import (
	"fmt"
	"sync"
)

func sendData(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i // Mengirim data ke channel
	}
	close(ch) // Menutup channel setelah selesai mengirim data
}

func receiveData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Menerima data: %d\n", data)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go sendData(ch, &wg) // Menjalankan sendData sebagai goroutine

	wg.Add(1)
	go receiveData(ch, &wg) //Menjalankan receiveData sebagai goroutine

	wg.Wait()
} */

/*
// Channel Select
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Pesan dari goroutine 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}  */

/*
// menerima data dari multiple channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Pesan dari goroutine 1"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
} */

/*
// Timeout dengan 'select'
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Pesan dari goroutine"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Tidak ada pesan dalam waktu yang diharapkan.")
	}
} */

// Non-blocking Operations
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)

	// Penerima non-blocking
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("Tidak ada pesan yang tersedia.")
	}

	// Mengirim data dari channel
	ch <- "Pesan dari goroutine"

	// Menerima pesan yang dikirimkan
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("Tidak ada pesan yang tersedia.")
	}
}
