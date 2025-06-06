package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello, World! from a goroutine")
	}()

	fmt.Println("Hello, World! from the main function")

	time.Sleep(1 * time.Second)
}
