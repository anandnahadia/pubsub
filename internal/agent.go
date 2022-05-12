package internal

import (
	"fmt"
	"sync"
)

type Subscribers map[string]*Subscriber
type Agent struct {
	topics      map[int]Subscribers
	subscribers Subscribers
	mut         sync.RWMutex
}

//create new Agent
func NewAgent() *Agent {
	Agent := Agent{
		subscribers: Subscribers{},
		topics:      map[int]Subscribers{},
	}
	return &Agent
}

//AddNewSubscriberToAgent add a new subscriber to the Agent.
func (a *Agent) AddNewSubscriberToAgent() *Subscriber {
	a.mut.Lock()
	defer a.mut.Unlock()
	id, s := CreateNewSubscription()
	a.subscribers[id] = s //add subscriber id to Agent
	return s
}

// publish the message to given topic.
func (a *Agent) Publish(topicID int, msg string) {
	a.mut.RLock()
	subscribers := a.topics[topicID]
	a.mut.RUnlock()
	for _, s := range subscribers {
		m := NewMessage(msg, topicID)
		if !s.active {
			return
		}
		go (func(s *Subscriber) {
			s.Signal(m)
		})(s)
	}
}
func (a *Agent) AddSubscription(SubscriptionID string, topicID int) *Subscriber {
	// subscribe to given topic
	a.mut.Lock()
	defer a.mut.Unlock()
	if a.topics[topicID] == nil {
		a.topics[topicID] = Subscribers{}
	}
	_, s := CreateNewSubscriptionwithID(SubscriptionID)
	s.AddTopic(topicID)
	a.topics[topicID][SubscriptionID] = s
	fmt.Println("Subscribed for topic:", topicID, "subscriptionID", SubscriptionID)
	return s
}

func (a *Agent) Subscribe(SubscriptionID string, SubscriberFunc func()) {
	s := a.subscribers[SubscriptionID]
	s.SubscriberFunc = SubscriberFunc
}

func (a *Agent) UnSubscribe(SubscriptionID string) {
	// unsubscribe to given topic
	a.mut.RLock()
	defer a.mut.RUnlock()
	for topicID, _ := range a.topics {
		delete(a.topics[topicID], SubscriptionID)
	}
}
func (a *Agent) UnsubscribeFromTopic(SubscriptionID string, topicID int) {
	// unsubscribe to given topic
	a.mut.RLock()
	defer a.mut.RUnlock()
	delete(a.topics[topicID], SubscriptionID)
	// a.RemoveTopic(topic)
}
