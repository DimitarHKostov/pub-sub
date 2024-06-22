package main

import (
	"fmt"
	"math/rand"
	"pub-sub/pkg/pubsub"
	"time"
)

const (
	minPrice          = 30
	maxPrice          = 100
	timeBetweenOffers = 10 * time.Second
)

var topics = []pubsub.Topic{pubsub.Sofia, pubsub.Prague, pubsub.Barcelona, pubsub.Rome}

func pricePublisher(broker *pubsub.Broker) {
	i := 0

	for {
		topic := topics[i%len(topics)].String()
		msg := fmt.Sprintf("%d", rand.Intn(maxPrice-minPrice+1)+minPrice)
		fmt.Printf("\n-*- New offer: {%s$} for {%s}: \n", msg, topic)

		go broker.Publish(topic, msg)

		i++
		time.Sleep(timeBetweenOffers)
	}
}

func addSubscriberForTopics(broker *pubsub.Broker, s *pubsub.Subscriber, topics []pubsub.Topic) {
	for _, topic := range topics {
		broker.Subscribe(s, topic.String())
	}

	go s.Listen()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	broker := pubsub.NewBroker()

	john := broker.AddSubscriber("John Smith")
	michael := broker.AddSubscriber("Michael Brown")
	olivia := broker.AddSubscriber("Olivia Smith")

	addSubscriberForTopics(broker, john, []pubsub.Topic{pubsub.Sofia})
	addSubscriberForTopics(broker, michael, []pubsub.Topic{pubsub.Prague, pubsub.Barcelona, pubsub.Rome})
	addSubscriberForTopics(broker, olivia, []pubsub.Topic{pubsub.Sofia, pubsub.Prague, pubsub.Barcelona, pubsub.Rome})

	go pricePublisher(broker)

	fmt.Scanln()
	fmt.Println("Done!")
}
