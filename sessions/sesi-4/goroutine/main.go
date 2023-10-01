package main

import "fmt"

func main() {
	fmt.Println("My first goroutine!");

	// 
	go func() {
		fmt.Println("My another goroutine")
	}();
}