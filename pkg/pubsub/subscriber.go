package pubsub

import (
	"fmt"
	"sync"
)

type Subscriber struct {
	id       string
	messages chan *Message
	topics   map[string]bool
	active   bool
	mutex    sync.RWMutex
}

func CreateNewSubscriber(name string) *Subscriber {
	return &Subscriber{
		id:       name,
		messages: make(chan *Message),
		topics:   map[string]bool{},
		active:   true,
	}
}

func (s *Subscriber) AddTopic(topic string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.topics[topic] = true
}

func (s *Subscriber) RemoveTopic(topic string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	delete(s.topics, topic)
}

func (s *Subscriber) GetTopics() []string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	topics := []string{}

	for topic := range s.topics {
		topics = append(topics, topic)
	}

	return topics
}

func (s *Subscriber) Destruct() {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.active = false
	close(s.messages)
}

func (s *Subscriber) Signal(msg *Message) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.active {
		s.messages <- msg
	}
}

func (s *Subscriber) Listen() {
	for {
		if msg, ok := <-s.messages; ok {
			fmt.Printf("{%s}, received offer: {%s$} for destination: {%s}\n", s.id, msg.GetMessageBody(), msg.GetTopic())
		}
	}
}
