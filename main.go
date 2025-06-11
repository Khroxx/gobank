package main

import "fmt"


func main() {
	server := newAPIServer(":8080")
	server.Run()
	fmt.Println("Server running on ", server)
}
