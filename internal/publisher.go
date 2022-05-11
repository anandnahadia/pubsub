package internal
topic1 := Topic{topicID:1,name:"Weather"}
topic2 := Topic{topicID:2,name:"Entertainment"}
topic3 := Topic{topicID:3,name:"Music"}
topic4 := Topic{topicID:4,name:"Sports"}

topics := []topic{topic1,topic2,topic3,topic4}
func PublishMessages(Agent *pubsub.Broker)(){
    for {
        randValue := topics[rand.Intn(len(topic))]
        msg:= fmt.Sprintf("%f", rand.Float64())
        go Agent.Publish(randValue.topicID, msg)
        r := rand.Intn(4)
        time.Sleep(time.Duration(r) * time.Second) 
    }
}