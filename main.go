package main
import (
    "fmt"
    "math/rand"
    "time"
    "github.com/anandnahadia/internal"
)

func main(){
    // create new agent
    agent := internal.NewAgent()
    // create new subscriber
    s1 := agent.AddNewSubscriberToAgent()
    // subscribe Weather and Entertainment to s1.
    agent.AddSubscription(s1, topic1.topicID)
    // create new subscriber
    s2 := agent.AddNewSubscriberToAgent()
    // subscribe Music and Sports to s2.
    broker.Subscribe(s2, topic2.topicID)

	s3 := agent.AddNewSubscriberToAgent()
    // subscribe Weather and Entertainment to s1.
    agent.AddSubscription(s3, topic3.topicID)
    // create new subscriber
    s4 := agent.AddNewSubscriberToAgent()
    // subscribe Music and Sports to s2.
    broker.Subscribe(s4, topic4.topicID)

    go (func(){
        time.Sleep(3*time.Second)
        agent.Subscribe(s2, topic2.TopicID)
    })()
    go (func(){
        time.Sleep(5*time.Second)
        agent.Unsubscribe(s2, topic3.TopicID)
    })()

    go (func(){
        time.Sleep(10*time.Second)
        agent.RemoveSubscriber(s2)
    })()
    // Concurrently publish the values.
    go internal.PublishMessages(broker)
    // Concurrently listens from s1.
    go s1.Listen()
    // Concurrently listens from s2.
    go s2.Listen()
	go s3.Listen()
	go s4.Listen()
    // to prevent terminate
    fmt.Scanln()
    fmt.Println("Done!")
}