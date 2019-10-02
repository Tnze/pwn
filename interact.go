package pwn

import (
	"bufio"
	"io"
	"log"
	"os"
)

// Program is a remote or local program.
type Program struct {
	io.Reader
	io.Writer

	BufReader *bufio.Reader
}

// NewProgram initial an Program with io.Reader and io.Writer
func NewProgram(r io.Reader, w io.Writer) *Program {
	return &Program{
		Reader:    r,
		Writer:    w,
		BufReader: bufio.NewReader(r),
	}
}

// Interactive interact directly with the application.
func (p Program) Interactive() {
	ch := make(chan error, 2)

	go func() {
		_, err := io.Copy(p.Writer, os.Stdin)
		ch <- err
	}()
	go func() {
		_, err := io.Copy(os.Stdout, p.BufReader)
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
func (p *Program) Write(data []byte) {
	_, err := p.Writer.Write(data)
	if err != nil {
		log.Fatalf("Write data error: %v", err)
	}
}

// ReadLine read data until the first occurrence of "\n"
func (p *Program) ReadLine() []byte {
	line, err := p.BufReader.ReadBytes('\n')
	if err != nil {
		log.Fatalf("ReadLine error: %v", err)
	}

	return line
}

// ReadByte reads and returns a single byte
func (p *Program) ReadByte() byte {
	b, err := p.BufReader.ReadByte()
	if err != nil {
		log.Fatalf("ReadByte error: %v", err)
	}

	return b
}
