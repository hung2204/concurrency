package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func findText(filePath string, key string, chanCount chan int) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	count := strings.Count(string(fileContent), key)
	chanCount <- count

	defer close(chanCount)
}

func main() {
	//file1 := "file1.txt.golang.1998.go_build_Hung_hung_Concurrency_deadlock.exe.golang.data"
	//file2 := "file2.txt.22041998.golang.hung.dn.:\\Users\\Admin\\AppData\\Local\\JetBrains\\GoLand2023.1"
	key := "golang"
	chanCount1 := make(chan int)
	chanCount2 := make(chan int)
	go findText("./file1.txt", key, chanCount1)
	go findText("./file2.txt", key, chanCount2)
	log.Printf("Từ khóa %v xuất hiện tổng %d lần trong các file", key, <-chanCount1+<-chanCount2)
}
