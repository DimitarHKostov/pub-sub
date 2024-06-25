package main

import (
	"fmt"
	"math/rand"
	"pub-sub/pkg/pubsub"
	"time"
)

func addSubscriberForTopics(broker *pubsub.Broker, s *pubsub.Subscriber, topics []pubsub.Topic) {
	for _, topic := range topics {
		broker.Subscribe(s, topic.String())
	}

	go s.Listen()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	broker := pubsub.NewBroker()

	person1 := pubsub.NewSubscriber("Person1")
	person2 := pubsub.NewSubscriber("Person2")
	person3 := pubsub.NewSubscriber("Person3")

	publisher := pubsub.NewPublisher(broker)

	addSubscriberForTopics(broker, person1, []pubsub.Topic{pubsub.Sofia})
	addSubscriberForTopics(broker, person2, []pubsub.Topic{pubsub.Prague, pubsub.Barcelona, pubsub.Rome})
	addSubscriberForTopics(broker, person3, []pubsub.Topic{pubsub.Sofia, pubsub.Prague, pubsub.Barcelona, pubsub.Rome})

	go publisher.Publish()

	fmt.Scanln()
	fmt.Println("Done!")
}
