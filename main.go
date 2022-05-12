package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/anandnahadia/pubsub/internal"
)

var (
	topic1 = internal.Topic{TopicID: 1, Name: "Weather"}
	topic2 = internal.Topic{TopicID: 2, Name: "Entertainment"}
	topic3 = internal.Topic{TopicID: 3, Name: "Music"}
	topic4 = internal.Topic{TopicID: 4, Name: "Sports"}
)

func main() {
	// create new agent
	agent := internal.NewAgent()
	// create new subscribers
	// subscribe topics.
	s1 := agent.AddSubscription("1", topic1.TopicID)

	s2 := agent.AddSubscription("2", topic2.TopicID)

	s3 := agent.AddSubscription("3", topic3.TopicID)

	s4 := agent.AddSubscription("4", topic4.TopicID)

	time.Sleep(3 * time.Second)
	wg := &sync.WaitGroup{}
	wg.Add(5)
	ctx, cancel := context.WithCancel(context.Background())
	// Concurrently publish to topics.
	go internal.PublishMessages(ctx, wg, agent)
	// Concurrently listens from topics.

	go s1.Listen(ctx, wg)

	go s2.Listen(ctx, wg)

	go s3.Listen(ctx, wg)

	go s4.Listen(ctx, wg)

	go func() {
		os.Stdin.Read(make([]byte, 1)) // wait for Enter keystroke
		cancel()                       // cancel the associated context
	}()

	wg.Wait()
	fmt.Println("Done!")
}
