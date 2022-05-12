package internal

import "fmt"

type Topic struct {
	TopicID int
	Name    string
}
type Message struct {
	topicID int
	body    string
}

func NewMessage(msg string, topicID int) *Message {
	// Returns the message object
	return &Message{
		topicID: topicID,
		body:    msg,
	}
}
func CreateTopic(topicID int) *Topic {
	topic := Topic{
		TopicID: topicID,
		Name:    "New Topic",
	}
	return &topic
}

func DeleteTopic(topicID int) {
	topic := Topic{
		TopicID: topicID,
	}
	fmt.Println(topic)
	// delete(topic)
}
