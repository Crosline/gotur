package examples

import (
	"flag"

	"fmt"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	host := flag.String("host", "0.0.0.0", "Host address to bind to")
	flag.Parse()
	
	cfg := config.New(*host, *port)

	httpServer := server.New(cfg)
	httpServer.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	fmt.Printf("Server started on %s:%d\n", cfg.Host(), cfg.Port())
	
	// Wait for shutdown signal
	<-sigChan
	fmt.Println("Shutting down server...")
	httpServer.Stop()
}
