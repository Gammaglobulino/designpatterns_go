package chain

import (
	"../chain_of_responsability"
	"strings"
	"testing"
)

func TestCreateDefaultChain(t *testing.T) {
	myWrite := chain.MyTestWriter{}
	writeLogger := chain.WriteLogger{
		Writer: &myWrite,
	}
	secondLogger := chain.SecondLogger{NextChain: &writeLogger}
	chain := chain.FirstLogger{NextChain: &secondLogger}
	t.Run("3 loggers, 2 writes to console if the word 'Hello' found", func(t *testing.T) {
		chain.Next("message that breaks the chain\n")
		if myWrite.ReceivedMessage != nil {
			t.Error("Last link should not receive any message")
		}
		chain.Next("Hello\n")
		if !strings.Contains(*myWrite.ReceivedMessage, "Hello") || *myWrite.ReceivedMessage == "" {
			t.Fatal("Last link didn't received expected message")
		}
	})
}
