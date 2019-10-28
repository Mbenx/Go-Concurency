package main

import (
	"fmt"
	"time"
)

func printSalam(text string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(text)
	}
}

func main() {
	start := time.Now()

	go printSalam("Hallo!!")
	printSalam("Selamat Datang")

	fmt.Println(time.Since(start))
}
