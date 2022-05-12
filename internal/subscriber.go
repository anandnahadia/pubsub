package internal

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"sync"
)

type Subscriber struct {
	ID             string        // id of subscriber
	messages       chan *Message // messages channel
	topicID        int           // topics it is subscribed to.
	active         bool          // if given subscriber is active
	mutex          sync.RWMutex  // lock
	SubscriberFunc func()
}

func CreateNewSubscription() (string, *Subscriber) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	id := fmt.Sprintf("%X-%X", b[0:4], b[4:8])
	return id, &Subscriber{
		ID:       id,
		messages: make(chan *Message),
		active:   true,
	}
}
func CreateNewSubscriptionwithID(id string) (string, *Subscriber) {

	return id, &Subscriber{
		ID:       id,
		messages: make(chan *Message),
		active:   true,
	}
}

// add topic to the subscriber
func (s *Subscriber) AddTopic(topicID int) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.topicID = topicID
}

// remove topic to the subscriber
func (s *Subscriber) RemoveTopic(topicID int) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.topicID = 0
}

// Gets the message from the channel
func (s *Subscriber) Signal(msg *Message) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.active {
		s.messages <- msg
	}
}

// Listens to the message channel, prints once received.
func (s *Subscriber) Listen(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("msg", "exiting from Listen goroutine")
			return
		case msg, ok := <-s.messages:
			if ok {
				fmt.Println("message received from topic", msg.topicID, msg.body)
			}

		}
	}
}
