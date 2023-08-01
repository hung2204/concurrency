package main

import (
	"fmt"
	"log"
	"sync"
)

func printNumber(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printChar(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	// b1: tạo biến có kiểu sync.WaitGroup: wg
	var wg sync.WaitGroup

	// b2: add số lượng goroutine cần phải đợi: wg.Add(2)
	wg.Add(2)

	// b3: thông báo khi goroutine chạy xong: wg.Done()

	// b4: gọi method ở hàm cần chờ cho các goroutine chạy xong: wg.Wait()

	go printNumber(&wg)
	go printChar(&wg)
	//time.Sleep(3 * time.Second)
	wg.Wait()

	fmt.Println()
	log.Println("Done")
}
