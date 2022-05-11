package pubsub
import (
    "crypto/rand"
    "fmt"
    "log"
    "sync"
)

type Subscriber struct {
    id string // id of subscriber
    messages chan* Message // messages channel
    topicID int // topics it is subscribed to.
    active bool // if given subscriber is active
    mutex sync.RWMutex // lock
}
func CreateNewSubscription() (string, *Subscriber) {
    b := make([]byte, 8)
    _, err := rand.Read(b)
    if err != nil {
        log.Fatal(err)
    }
    id := fmt.Sprintf("%X-%X", b[0:4], b[4:8])
    return id, &Subscriber{
        id: id,
        messages: make(chan *Message),
        active: true,
    }
}
func CreateNewSubscriptionwithID(id int) (string, *Subscriber) {
   
    return  id,&Subscriber{
        id: id,
        messages: make(chan *Message),
        active: true,
    }
}
// add topic to the subscriber
func (s * Subscriber)AddTopic(topicID int)(){
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    s.topicID = topicID
}
// remove topic to the subscriber
func (s * Subscriber)RemoveTopic(topicID int)(){ 
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    delete(s.topics, topic)
}

// Gets the message from the channel
func (s *Subscriber)Signal(msg *Message) () {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    if s.active{
        s.messages <- msg
    }
}
// Listens to the message channel, prints once received.
func (s *Subscriber)Listen() {
    for {
        if msg, ok := <- s.messages; ok {
            fmt.Printf("Subscriber %s, received: %s from topic: %s\n", s.id, msg.GetMessageBody(), msg.GetTopic())
        }
    }
}