package internal

type Topic struct {
	topicID int
	name string
}
type Message struct {
    topic string
    body string
}
func NewMessage(msg string, topic string) (* Message) {
    // Returns the message object
    return &Message{
        topic: topic,
        body: msg,
    }
}
func CreateTopic(topicID int) (* topic) {
    topic := Topic{
		topicID: topicID
		name: "New Topic"
	}
	return topic
}

func DeleteTopic(topicID int) {
	topic := topic{
		id: topicID
	}
	delete(topic)
}