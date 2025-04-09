package main

import (
	"fmt"
	"github.com/gemaskus/httpfromtcp/internal/request"
	"log"
	"net"
)

func main() {
	//	var fileName string = "messages.txt"

	listener, err := net.Listen("tcp", ":42069")
	defer listener.Close()
	if err != nil {
		log.Fatalf("Could not create TCP listener: %v", err)
	}

	for {
		conn, err := listener.Accept()

		req, err := request.RequestfromReader(conn)
		if err != nil {
			listener.Close()
			log.Fatalf("Could not accept the connection: %v\n", err)
		}

		fmt.Println("Request line:")
		fmt.Printf("- Method: %s\n", req.RequestLine.Method)
		fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)
	}
}
