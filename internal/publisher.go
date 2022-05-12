package internal

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	topic1 = Topic{TopicID: 1, Name: "Weather"}
	topic2 = Topic{TopicID: 2, Name: "Entertainment"}
	topic3 = Topic{TopicID: 3, Name: "Music"}
	topic4 = Topic{TopicID: 4, Name: "Sports"}
	topics = []Topic{topic1, topic2, topic3, topic4}
)

func PublishMessages(ctx context.Context, wg *sync.WaitGroup, Agent *Agent) {
	fmt.Println("msg", "inside publish message")
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("msg", "exiting from PublishMessage goroutine")
			return
		default:
			randValue := topics[rand.Intn(len(topics))]
			msg := fmt.Sprintf("%f", rand.Float64())
			go Agent.Publish(randValue.TopicID, msg)
			r := rand.Intn(4)
			time.Sleep(time.Duration(r) * time.Second)
		}
	}
}
