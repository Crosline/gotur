package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"

	s "github.com/crosline/gotur/socket"
)

func main() {
	// Create a new UDP socket for the client
	socket, err := s.NewUDPSocket()
	if err != nil {
		log.Fatalf("Failed to create UDP socket: %v", err)
	}
	defer socket.Close()

	// Bind to any available local port
	if err := socket.Bind("0.0.0.0", 0); err != nil {
		log.Fatalf("Failed to bind socket: %v", err)
	}

	// Set up the server address
	serverAddr, err := s.ParseIPv4("127.0.0.1")
	if err != nil {
		log.Fatalf("Invalid server address: %v", err)
	}

	// Update the socket address to point to the server
	socket.SetRemoteAddress(syscall.SockaddrInet4{
		Port: 8000,
		Addr: serverAddr,
	})

	fmt.Println("UDP Client started. Type messages and press Enter to send.")
	fmt.Println("Type 'exit' to quit.")

	// Read input from the console
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		message := scanner.Text()
		if message == "exit" {
			break
		}

		// Send the message to server
		if err := socket.Send([]byte(message)); err != nil {
			log.Printf("Error sending message: %v", err)
			continue
		}

		// Get response
		buffer := make([]byte, 1024)
		n, err := socket.Receive(buffer)
		if err != nil {
			log.Printf("Error receiving response: %v", err)
			continue
		}

		// Print the response
		fmt.Printf("Server: %s\n", string(buffer[:n]))
	}

	fmt.Println("Client shutting down.")
}