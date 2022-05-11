package internal
import (
    "fmt"
    "sync”
)
type Subscribers map[string]*Subscriber
type Agent struct {
    topics map[int]Subscribers
    subscribers Subscribers
    mutex sync.RWMutex 
}
//create new Agent
func NewAgent() (*Agent){
    Agent := Agent{
        subscribers: Subscribers{}
    }
    return Agent
}
//AddNewSubscriberToAgent add a new subscriber to the Agent.
func (a *Agent)AddNewSubscriberToAgent()(*Subscriber){
    a.mut.Lock()
    defer a.mut.Unlock()
    id, s := CreateNewSubscriber()
    a.subscribers[id] = s;//add subscriber id to Agent
    return s
}
// publish the message to given topic.
func (a *Agent) Publish(topicID int, msg string) {
    a.mut.RLock()
    bTopics := a.topics[topicID]
    a.mut.RUnlock()
    for _, s := range bTopics {
        m:= NewMessage(msg, topic)
        if !s.active{
            return
        }
        go (func(s *Subscriber){
            s.Signal(m)
        })(s)
    }
}
func (a *Agent) AddSubscription(topicID int, SubscriptionID int) {
    // subscribe to given topic
    b.mut.Lock()
    defer b.mut.Unlock()
    if  b.topics[topic] == nil {
        b.topics[topic] = Subscribers{}
    }
    s:= CreateNewSubscriptionwithID(SubscriptionID)
    s.AddTopic(topicID)
    b.topics[topic][s.id] = s
    fmt.Printf("%s Subscribed for topic: %s\n", s.id, topic)
}


func(a *Agent)Subscribe(SubscriptionID int,SubscriberFunc func()){
    s := a.subscriber[SubscriptionID]
    s.SubscriberFunc = SubscriberFunc
}

func(a *Agent)UnSubscribe(SubscriptionID int) {
    // unsubscribe to given topic
    b.mut.RLock()
    defer b.mut.RUnlock()
    for topicID,_ := range b.topics {
        delete(b.topics[topicID], SubscriptionID)
    }
}
