package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	serverAddr := "localhost:42069"

	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		log.Fatalf("Could not create udp sender: %v", err)
	}

	con, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Error dialing UDP: %v", err)
	}
	defer con.Close()

	fmt.Printf("Sending to %s, Type your message and press Enter to send. Press CTRL+c to exit.\n", serverAddr)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Could not read input: %v", err)
		}

		_, err = con.Write([]byte(message))
		if err != nil {
			log.Fatalf("Could not send message to UDP connection: %v", err)
		}

		fmt.Printf("Message sent: %s", message)
	}
}
