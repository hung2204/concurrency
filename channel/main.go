package main

import (
	"log"
)

func printNumber(numChan chan int) {
	var count int
	for i := 0; i < 100; i++ {
		//fmt.Printf("%d ", i)
		count += i
	}
	numChan <- count // locked
}

func printChar(charChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		//fmt.Printf("%c ", i)
		result += string(i)
	}
	charChan <- result
}

func main() {
	chanNumber := make(chan int)
	chanChar := make(chan string)

	go printNumber(chanNumber)
	go printChar(chanChar)

	log.Println("Kq chanNumber: ", <-chanNumber) // unlocked
	log.Println("kq chanChar: ", <-chanChar)
	log.Println("Done")
}
