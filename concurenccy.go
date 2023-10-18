package main

import (
	"fmt"
	"time"
)

func main() {
	go beriSalam("Hallo !!!")
	beriSalam("Selamat Datang!!!")
}

func beriSalam(salam string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(salam)
	}
}
