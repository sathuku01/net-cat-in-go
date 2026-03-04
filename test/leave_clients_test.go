package test

import (
	"testing"
	"time"

	ac "net-cat/internal"
)

func TestClientLeave(t *testing.T) {
	server := ac.NewServer(2)
	go server.Run()

	alice := &ac.Client{Name: "Alice", Messages: make(chan string, 10)}
	server.Join <- alice
	time.Sleep(20 * time.Millisecond)

	server.Leave <- alice
	time.Sleep(20 * time.Millisecond)

	if _, exists := server.Clients["Alice"]; exists {
		t.Error("Alice should have been removed from Clients map")
	}

	select {
	case _, ok := <-alice.Messages:
		if ok {
			t.Error("Alice's Messages channel should be closed")
		}
	default:
		t.Error("Alice's Messages channel should be closed")
	}
}