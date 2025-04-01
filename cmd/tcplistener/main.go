package main

import (
	"fmt"
	"io"
	"log"
	"net"
	//	"os"
	"strings"
)

func main() {
	//	var fileName string = "messages.txt"

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("Could not create TCP listener: %v", err)
	}

	for {
		conn, err := listener.Accept()

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("%s\n", line)
		}
		if err != nil {
			listener.Close()
			log.Fatalf("Could not accept the connection: %v\n", err)
		}

	}

	listener.Close()
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	fileChannel := make(chan string)

	go func() {
		var line string
		buffer := make([]byte, 8)
		for {
			n, err := f.Read(buffer)

			// Print read in byte string
			if n > 0 {
				parts := strings.Split(string(buffer[:n]), "\n")

				for i := 0; i < len(parts)-1; i++ {
					fileChannel <- line + parts[i]
					line = ""
				}

				line += parts[len(parts)-1]
			}

			if err == io.EOF {
				if line != "" {
					fileChannel <- line
				}
				break //we hit the end of the file
			}

			if err != nil {
				log.Fatalf("Failed to read from the file: %v", err)
			}
		}
		f.Close()
		close(fileChannel)
		fmt.Printf("Channel is closed\n")
	}()

	return fileChannel
}
