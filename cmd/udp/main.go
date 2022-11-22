package main

import (
	"flag"
	"fmt"
	"net"
	"net/netip"
	"os"
	"strconv"
	"time"
)

var (
	port   = flag.Int("port",       80, "IP Port to listen to")
	host   = flag.String("host",    "", "ip/host to listen")
)

func main() {
	flag.Parse()

	// Listen for incoming connections.
	addr, err := net.ResolveUDPAddr ( "udp", *host+":"+strconv.Itoa(*port) )
	if err != nil {
		fmt.Println("Cannot parse address:", err.Error())
		os.Exit(1)
	}
	
	sock, err := net.ListenUDP("udp", addr )
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	
	// Close the listener when the application closes.
	defer sock.Close()
	
	fmt.Println("Listening on " + *host + ":" + strconv.Itoa(*port) + "(udp)" )

	buf := make([]byte, 1024, 4096)
	for {
		rlen, addr, err := sock.ReadFromUDPAddrPort(buf)
		if err != nil {
			fmt.Println("Error reading from socket:", err.Error() )
		}

		// Handle connections in a new goroutine.
		go handleRequest(sock, rlen, addr, buf[0:rlen])
	}
}

// Handles incoming requests.
func handleRequest(sock *net.UDPConn, reqLen int, addr netip.AddrPort, msg []byte) {
	s := string(msg)
	
	fmt.Printf( "Got %d bytes [%s]\n", reqLen, s )

	// Send a response back to person contacting us.
	hour, min, sec := time.Now().Clock()
	_, err := sock.WriteToUDPAddrPort( []byte (
		fmt.Sprintf (
			"Message received at %d:%02d:%02d containing %s\n",
			hour,
			min,
			sec,
			s)),
		addr)
	if err != nil {
		fmt.Println("Error replying:", err.Error() )
	}
}
