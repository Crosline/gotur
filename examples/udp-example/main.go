package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	gotur "github.com/crosline/gotur/server"
	s "github.com/crosline/gotur/socket"
)

func main() {
	// Create a new UDP server
	server, err := gotur.NewUDPServer()
	if err != nil {
		log.Fatalf("Failed to create UDP server: %v", err)
	}

	// Set up handler for incoming UDP messages
	server.Handle(func(socket s.Socket) {
		// Buffer to hold received data
		buffer := make([]byte, 1024)

		// Receive data
		n, err := socket.Receive(buffer)
		if err != nil {
			log.Printf("Error receiving data: %v", err)
			return
		}

		// Log received message
		message := string(buffer[:n])
		log.Printf("Received message: %s", message)

		// Prepare response
		response := fmt.Sprintf("Echo: %s (received at %s)",
			message, time.Now().Format(time.RFC3339))

		// Send response
		if err := socket.Send([]byte(response)); err != nil {
			log.Printf("Error sending response: %v", err)
		}
	})

	// Start the server
	if err := server.Start("127.0.0.1", 8000); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("UDP Server started on 127.0.0.1:8000")
	log.Println("Press Ctrl+C to stop the server")

	// Set up graceful shutdown on SIGINT or SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server...")
	if err := server.Stop(); err != nil {
		log.Printf("Error stopping server: %v", err)
	}
}
