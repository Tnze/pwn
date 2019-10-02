package pwn

import (
	"log"
	"net"
)

// Remote dial a TCP connection to remote host.
func Remote(addr string) Program {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Connect remote error: %v", err)
	}
	return Program{conn}
}
