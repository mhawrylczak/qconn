package main

import (
 "fmt"
 )
	
	
func main() {
	messages := make(chan []byte)
	port := ":10001"
	go Listen(port, messages)
	
	for{
		msg := <-messages
		fmt.Println(msg)
	}
	
	
	fmt.Println("end")
}