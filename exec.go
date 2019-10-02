package pwn

import (
	"io"
	"log"
	"os/exec"
)

type localReadWriter struct {
	out, err io.ReadCloser
	io.WriteCloser

	io.Reader // read from out & err
}

func newLocalReadWriter(out, err io.ReadCloser, in io.WriteCloser) *localReadWriter {
	return &localReadWriter{
		out:         out,
		err:         err,
		WriteCloser: in,
		Reader:      io.MultiReader(out, err),
	}
}

func (l *localReadWriter) Close() error {
	if err := l.out.Close(); err != nil {
		return err
	}
	if err := l.err.Close(); err != nil {
		return err
	}
	if err := l.WriteCloser.Close(); err != nil {
		return err
	}

	return nil
}

// Local run an local command.
func Local(cmd *exec.Cmd) io.ReadWriteCloser {
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

	return newLocalReadWriter(stdout, stderr, stdin)
}
