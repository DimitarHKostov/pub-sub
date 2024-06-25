package pubsub

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minPrice          = 30
	maxPrice          = 100
	timeBetweenOffers = 10 * time.Second
)

type Publisher struct {
	broker *Broker
}

var topics = []Topic{Sofia, Prague, Barcelona, Rome}

func NewPublisher(broker *Broker) *Publisher {
	return &Publisher{broker: broker}
}

func (p *Publisher) Publish() {
	i := 0

	for {
		topic := topics[i%len(topics)].String()
		msg := fmt.Sprintf("%d", rand.Intn(maxPrice-minPrice+1)+minPrice)
		fmt.Printf("\n-*- New offer: {%s$} for {%s}: \n", msg, topic)

		go p.broker.Publish(topic, msg)

		i++
		time.Sleep(timeBetweenOffers)
	}
}
