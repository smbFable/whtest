package main

import (
	"fmt"
	"time"
)

func main() {
	num := 0
	for {
		time.Sleep(1000 * time.Millisecond)
		num += 1
		fmt.Println("Секунд прошло: ", num)
	}
}
