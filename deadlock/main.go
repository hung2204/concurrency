package main

func main() {
	chanNumber := make(chan int)
	//chanNumber <- 1
	<-chanNumber
}
