# pubsub
pubsub or publish-subscribe service is a messaging pattern in which publishers send message to all subscribers on specific topics.

## Scenario
In cloud based systems with microservices architecture there is a need for async processing. There are several cases where an API request for a user completes and we need to give him an ok response but the system needs to do post processing of different kinds. Most of this post processing is done in an async fashion on different services which are deployed separately(lets call them workers). A simple method to approach this problem is to call these workers asynchronously from the user facing service and wait for them to process them without making the user wait for it, but it has problems: 1) we cannot wait indefinitely for the worker to respond because the task can be a long running one and the machine which initiated the request might go down before the request completes 2) the user facing service will need a context of all different post processing that is being done.

To solve these problems we would like to build a PubSub messaging system in Golang. The PubSub system will sit between the services to facilitate async communication. It will ensure reliable delivery of messages to the subscribers(workers) and decouple the senders and receivers.


## Primary components
The two clients of the system are Publishers and Subscribers. Publishers send the messages to the system and subscribers listen to messages from the system.
PubSub system has two components: Topics and Subscriptions. There is one to many mapping between Topics and Subscriptions. Each message that a Topic receives is propagated to all of its Subscriptions.
Request flows:
Publisher will Publish a message to the Topic.
Subscriber will Subscribe to the Subscription. We will build Push based Subscriptions, hence it will be the responsibility of the Subscription to ensure that all messages are propagated to the Subscriber.

## Constraints:
There can only be one Subscriber attached to a Subscription.
 One Subscription is attached to only one Topic.

## The PubSub system should support below methods:
CreateTopic(TopicID)

DeleteTopic(TopicID)

AddSubscription(TopicID,SubscriptionID); Creates and adds Subscription with id SubscriptionID to topicID.

DeleteSubscription(SubscriptionID)

Publish(TopicID, message); publishes the message on the given Topic

Subscribe(SubscriptionID, SubscriberFunc); SubscriberFunc is the callback function given by Subscriber, which is executed for each message of the Subscription.

UnSubscribe(SubscriptionID); After this the Subscriber will no longer receive messages from the Subscription with given SubscriptionID

Ack(SubscriptionID, MessageID); Called by Subscriber to intimate the Subscription that the message has been received and processed.
