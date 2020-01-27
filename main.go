package main

import (
	"fmt"
	"sync"
)

var (
	mutex             sync.Mutex
	total_books_count int
)

func init() {
	total_books_count = 100
}

func borrow_book(value int, wg *sync.WaitGroup) {
	mutex.Lock()

	if total_books_count < value {
		fmt.Printf("Insufficient Books asked for Borrowing. Please Try Again Later as the Library cannnot accomodate a request for %d books", value)
	} else {
		fmt.Printf("Borrowing %d Books from the Library which currently has %d books\n", value, total_books_count)
		total_books_count -= value
	}

	mutex.Unlock()
	wg.Done()
}

func return_book(value int, wg *sync.WaitGroup) {
	mutex.Lock()

	fmt.Printf("Returning %d Books from the Library which currently has %d books\n", value, total_books_count)
	total_books_count += value

	mutex.Unlock()
	wg.Done()

}

func main() {
	fmt.Printf("---MAIN PROGRAM---\n")

	var wg sync.WaitGroup
	wg.Add(2)

	go borrow_book(40, &wg)
	go return_book(20, &wg)

	wg.Wait()

	fmt.Printf("Current Book Count: %d\n\n", total_books_count)
}
