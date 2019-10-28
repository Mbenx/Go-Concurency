package main

import (
	"fmt"
	"time"
)

// channel
var itemsChannel = make(chan string)
var cleansItemsChannel = make(chan string)
var savedItemsChannel = make(chan string)

func main() {
	items := [7]string{"batu", "harta", "kerang", "harta", "batu", "harta", "kerang"}
	// menyelam mencari harta karun
	// terdiri dari banyak barang
	// 1. Penyelam, 2. Pembersih 3. Penyimpan
	// channel

	go penyelam(items)
	go pembersih()
	go penyimpan()

	time.Sleep(500 * time.Millisecond)
}

func penyelam(items [7]string) {
	for _, item := range items {
		if item == "harta" {
			fmt.Println("Berhasil Mendapatkan" + item)
			itemsChannel <- item
		}
	}
}

func pembersih() {
	for rawItem := range itemsChannel {
		fmt.Println("Berhasil Membersihkan" + rawItem)
		cleansItemsChannel <- rawItem
	}
}

func penyimpan() {
	for rawItem := range cleansItemsChannel {
		fmt.Println("Berhasil Menyimpan" + rawItem)
		savedItemsChannel <- rawItem
	}
}
