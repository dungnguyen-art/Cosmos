package main

import (
	"fmt"
)
//============================Cach 1 =====================================
// 1 gorountine
// toc do cham vi chi xu ly 1 tien trinh
// func main() {
// 	fmt.Println("Hello, blockchain lover")
// 	number := 10
// 	for i := 1; i<= number; i++ {
// 		fmt.Print(fib(i), " ")
// 	}

// }
// func fib(n int) int {
// 	if n <=1 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }
//===========================Cach 2 =======================================
// Da tien trinh
// Chay dong thoi
func main() {
	fmt.Println("Hello, blockchain lover")
	number := 10
	
	numberOfWorker := 5
	jobs := make(chan int, number)
	results := make(chan int, number)
	for i := 0; i<numberOfWorker; i++ {
		go worker(jobs, results)
	}

	for i := 1; i<= number; i++ {
		jobs <- i
	}
	close (jobs)

	for i := 0; i < number; i++ {
		fmt.Println(<- results)
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <=1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}