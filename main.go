package main

import (
	"fmt"
	"time"
)

func main() {
	num := 0
	for {
		time.Sleep(100 * time.Millisecond)
		num += 1
		fmt.Println("Секунды: ", num)
	}
}
