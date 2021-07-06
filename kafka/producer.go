package kafka

import (
	k "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func NewProducer() *k.Producer {
	producer, err := k.NewProducer(&k.ConfigMap{"bootstrap.servers": "localhost"})

	if err != nil {
		log.Printf("Failed to create producer: %s\n", err.Error())
	}

	return producer
}

func SendMessage(producer *k.Producer, message string, topic string) {
	deliveryChannel := make(chan k.Event)

	err := producer.Produce(&k.Message{
		TopicPartition: k.TopicPartition{Topic: &topic, Partition: k.PartitionAny},
		Value:          []byte(message),
	}, deliveryChannel)

	if err != nil {
		log.Printf("Could not send message %s\n", err.Error())
	}

	e := <-deliveryChannel
	m := e.(*k.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		log.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChannel)
}

func CloseProducer(producer *k.Producer) {
	producer.Close()
}
