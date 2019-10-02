package pwn

import (
	"io"
	"log"
	"net"
	"os"
)

// Remote dial a TCP connection to remote host.
func Remote(addr string) io.ReadWriteCloser {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Connect remote error: %v", err)
	}
	return conn
}

// Interactive interact directly with the application.
func Interactive(p io.ReadWriteCloser) {
	ch := make(chan error, 2)

	go func() {
		_, err := io.Copy(p, os.Stdin)
		ch <- err
	}()
	go func() {
		_, err := io.Copy(os.Stdout, p)
		ch <- err
	}()

	for i := 0; i < 2; i++ {
		err := <-ch
		if err != nil {
			log.Fatalf("Interactive error: %v", err)
		}
	}
}
