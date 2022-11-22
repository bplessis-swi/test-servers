package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"io"
	"time"
)

var (
	port   = flag.Int("port",       80, "IP Port to listen to")
	host   = flag.String("host",    "", "ip/host to listen")
)

func main() {
	flag.Parse()

	// Listen for incoming connections.
	l, err := net.Listen("tcp", *host+":"+strconv.Itoa(*port) )
	
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	
	fmt.Println("Listening on " + *host + ":" + strconv.Itoa(*port) + "(tcp)" )
	
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	fmt.Println ( "Opened ", conn.RemoteAddr().Network(), " connection from ", conn.RemoteAddr().String() )
	buf := make([]byte, 1024, 4096)
	// Read the incoming connection into the buffer.

	for {
		reqLen, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		} else {
			//fmt.Println("got", n, "bytes.")
			if reqLen > 0 {
				s := string(buf[:reqLen])
				fmt.Printf( "Got %d bytes [%s]\n", reqLen, s )

				// Send a response back to person contacting us.
				hour, min, sec := time.Now().Clock()
				conn.Write( []byte (
					fmt.Sprintf (
						"Message received at %d:%02d:%02d containing %s\n",
						hour,
						min,
						sec,
						s)))
			}
		}
	}
	
	// Close the connection when you're done with it.
	conn.Close()
}
