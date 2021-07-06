package kafka

import (
	"context"
	k "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

func CreateTopic(topic string, partitionCount int) {
	admin, err := k.NewAdminClient(&k.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Printf("Failed to create Admin client: %s\n", err.Error())
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	duration := 60 * time.Second

	_, err = admin.CreateTopics(
		ctx,
		[]k.TopicSpecification{
			{
				Topic:             topic,
				NumPartitions:     partitionCount,
				ReplicationFactor: 0,
			},
		}, k.SetAdminOperationTimeout(duration))

	if err != nil {
		log.Printf("Failed to create topic: %s\n", err.Error())
		return
	}

	admin.Close()
}
