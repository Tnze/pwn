package pwn

import (
	"io"
	"log"
	"net"
	"os/exec"
)

// Remote dial a TCP connection to remote host.
func Remote(addr string) *Program {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Connect remote error: %v", err)
	}
	return NewProgram(conn, conn)
}

// Local run an local command.
func Local(cmd *exec.Cmd) *Program {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Open stdout error: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Open stderr error: %v", err)
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("Open stdin error: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Start cmd error: %v", err)
	}

	return NewProgram(io.MultiReader(stdout, stderr), stdin)
}
