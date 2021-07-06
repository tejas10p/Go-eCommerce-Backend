package kafka

import (
	k "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func NewConsumer(group string) *k.Consumer {
	consumer, err := k.NewConsumer(&k.ConfigMap{
		"bootstrap.servers":     "localhost",
		"group.id":              group,
		"broker.address.family": "v4",
		"session.timeout.ms":    8000})
	if err != nil {
		log.Printf("Failed to create consumer: %s\n", err.Error())
	}

	return consumer
}

func ConsumeMessage(consumer *k.Consumer, topic string) string {
	err := consumer.Subscribe(topic, nil)
	if err != nil {
		log.Printf("Failed to subscribe to topic - %s\n", err.Error())
		return ""
	}

	wait := true

	for wait == true {
		event := consumer.Poll(0)
		switch e := event.(type) {
		case *k.Message:
			return (string)(e.Value)
		case k.Error:
			log.Printf("Encoutered error - %v\n", e.Code())
			if e.Code() == k.ErrAllBrokersDown {
				wait = false
			}
		default:
			log.Printf("Ignored - %v\n", e)
		}
	}

	return ""
}

func CloseConsumer(consumer *k.Consumer) {
	err := consumer.Close()
	if err != nil {
		log.Printf("Failed to close consumer - %s\n", err.Error())
	}
}
