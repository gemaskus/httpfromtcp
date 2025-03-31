package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName string = "messages.txt"

	file, er := os.Open(fileName)
	if er != nil {
		log.Fatalf("Could not open the file: %s", fileName)
	}

	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}

	// var line string
	// buffer := make([]byte, 8)
	//
	//	for {
	//		n, err := file.Read(buffer)
	//
	//		// Print read in byte string
	//		if n > 0 {
	//			parts := strings.Split(string(buffer[:n]), "\n")
	//
	//			for i := 0; i < len(parts)-1; i++ {
	//				fmt.Printf("read: %s%s\n", line, parts[i])
	//				line = ""
	//			}
	//
	//			line += parts[len(parts)-1]
	//		}
	//
	//		if err == io.EOF {
	//			break //we hit the end of the file
	//		}
	//
	//		if err != nil {
	//			log.Fatalf("Failed to read from the file: %v", err)
	//		}
	//	}
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
	}()

	return fileChannel
}
