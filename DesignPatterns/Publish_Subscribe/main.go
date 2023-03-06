package main

import (
	"fmt"
	"sync"
)

// PubSub manages the communication between publishers and subscribers
type PubSub struct {
	mu   sync.Mutex    // ensures that access to subs is synchronized
	subs []chan string // subs is a slice of channels subscribed to the PubSub
}

// Subscribe adds a new channel to the list of subscribers of the PubSub
func (ps *PubSub) Subscribe(c chan string) {
	ps.mu.Lock() // lock the mutex to protect the shared subs slice
	ps.subs = append(ps.subs, c)
	ps.mu.Unlock() // unlock the mutex
}

// Publish sends a message to all the subscribers of the PubSub.
func (ps *PubSub) Publish(msg string) {
	ps.mu.Lock()         // protect the shared subs slice
	defer ps.mu.Unlock() // Unlock the mutex at the end (when the function returns)

	// send a message to all subscribers
	for _, sub := range ps.subs {
		select {
		case sub <- msg: // send the message
		default:
			// if the channel is full, continue with the next subscriber
		}
	}
}

func main() {
	// create a new PubSub
	ps := &PubSub{
		subs: make([]chan string, 0),
	}

	// subscriber channel with a buffer of 10
	c1 := make(chan string, 10)
	// Subscribe the channel to the PubSub
	ps.Subscribe(c1)

	// second subscriber channel with a buffer of 10
	c2 := make(chan string, 10)
	// sub the channel to the PubSub
	ps.Subscribe(c2)

	// publish two messages to all subscribers
	ps.Publish("Hello world")
	ps.Publish("Hello again")

	// receive messages from subscribers using loops and a select statement
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("subscriber 1 received:", msg1)
		case msg2 := <-c2:
			fmt.Println("subscriber 2 received:", msg2)
		default:
			// if there are no more messages, exit the loop
			return
		}
	}
}
