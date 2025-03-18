package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var fileName string = "messages.txt"

	file, er := os.Open(fileName)
	if er != nil {
		log.Fatalf("Could not open the file: %s", fileName)
	}
	buffer := make([]byte, 8)
	for {
		n, err := file.Read(buffer)

		// Print read in byte string
		if n > 0 {
			fmt.Printf("read: %s\n", buffer[:n])
		}

		if err == io.EOF {
			break //we hit the end of the file
		}

		if err != nil {
			log.Fatalf("Failed to read from the file: %v", err)
		}
	}
}
