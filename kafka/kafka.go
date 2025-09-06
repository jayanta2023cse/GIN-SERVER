// Package kafka handles application configuration, including loading environment variables.
package kafka

import (
	"app/config"
	"context"
	"log"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer
var consumerGroup sarama.ConsumerGroup
var consumerCtx context.Context
var consumerCancel context.CancelFunc

func InitKafka() {
	broker := config.AppConfig.KafkaBroker
	if broker == "" {
		broker = "localhost:9092" // Matches the Kafka container's host-mapped port
	}
	topic := config.AppConfig.KafkaTopic
	if topic == "" {
		topic = "test-topic"
	}
	groupID := config.AppConfig.KafkaGroupID
	if groupID == "" {
		groupID = "test-group"
	}

	// Initialize context before starting consumer
	consumerCtx, consumerCancel = context.WithCancel(context.Background())

	err := InitProducer(broker, topic)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}

	go StartConsumer(broker, topic, groupID)

	log.Printf("Kafka initialized with broker: %s, topic: %s, groupID: %s", broker, topic, groupID)
}

func InitProducer(broker, topic string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	var err error
	producer, err = sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Test message from producer"),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send test message: %v", err)
	} else {
		log.Printf("Test message sent to partition %d, offset %d", partition, offset)
	}
	return nil
}

func StartConsumer(broker, topic, groupID string) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	var err error
	consumerGroup, err = sarama.NewConsumerGroup([]string{broker}, groupID, config)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka consumer group: %v", err)
	}

	consumer := &Consumer{
		ready: make(chan bool),
	}

	for {
		if consumerCtx.Err() != nil {
			return
		}
		err := consumerGroup.Consume(consumerCtx, []string{topic}, consumer)
		if err != nil {
			log.Printf("Error from consumer: %v", err)
		}
	}
}

func CloseProducer() {
	if producer != nil {
		if err := producer.Close(); err != nil {
			log.Printf("Failed to close Kafka producer: %v", err)
		} else {
			log.Println("Kafka producer closed")
		}
		producer = nil
	}
}

func CloseConsumer() {
	if consumerGroup != nil {
		// Cancel the consumer context to stop consumption
		if consumerCancel != nil {
			consumerCancel()
		}
		// Close the consumer group
		if err := consumerGroup.Close(); err != nil {
			log.Printf("Failed to close Kafka consumer group: %v", err)
		} else {
			log.Println("Kafka consumer group closed")
		}
		consumerGroup = nil
	}
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message received: topic=%s, partition=%d, offset=%d, value=%s",
			message.Topic, message.Partition, message.Offset, string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}
