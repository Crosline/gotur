package examples

import (
	"flag"

	"github.com/crosline/gotur"
	"github.com/crosline/gotur/socket"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	host := flag.String("host", "0.0.0.0", "Host address to bind to")
	flag.Parse()


	udp, err := socket.NewUDPSocket()
	if err != nil {
		panic(err)
	}

	if err = udp.Bind(*host, *port); err != nil {
		panic(err)
	}

	if err = udp.Listen(); err != nil {
		panic(err)
	}

	server = gotur.NewBaseServer(*udp)

	<-make(chan struct{})
}
