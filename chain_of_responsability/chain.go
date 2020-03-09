package chain

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (c *FirstLogger) Next(s string) {
	fmt.Printf("First logger:%s\n", s)
	if c.NextChain != nil {
		c.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (c *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Sprintf("Second logger: %s\n", s)
		c.NextChain.Next(s)
	}
	return
}

type WriteLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriteLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte(s))
	}
	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}

type MyTestWriter struct {
	ReceivedMessage *string
}

func (w *MyTestWriter) Write(p []byte) (int, error) {
	if w.ReceivedMessage == nil {
		w.ReceivedMessage = new(string)
	}
	tmpMsg := fmt.Sprintf("%s%s", *w.ReceivedMessage, p)
	w.ReceivedMessage = &tmpMsg
	return len(p), nil
}

func (w *MyTestWriter) Next(s string) {
	w.Write([]byte(s))
}
