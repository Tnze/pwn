package pwn

import (
	"io"
	"log"
	"os"
)

// Program is a remote or local program.
type Program struct {
	Conn io.ReadWriteCloser
}

// Interactive interact directly with the application.
func (p Program) Interactive() {
	ch := make(chan error, 2)

	go func() {
		_, err := io.Copy(p.Conn, os.Stdin)
		ch <- err
	}()
	go func() {
		_, err := io.Copy(os.Stdout, p.Conn)
		ch <- err
	}()

	for i := 0; i < 2; i++ {
		err := <-ch
		if err != nil {
			log.Fatalf("Interactive error: %v", err)
		}
	}
}

// Write data into Program's stdin.
func (p Program) Write(data []byte) {
	_, err := p.Conn.Write(data)
	if err != nil {
		log.Fatalf("Write data error: %v", err)
	}
}
