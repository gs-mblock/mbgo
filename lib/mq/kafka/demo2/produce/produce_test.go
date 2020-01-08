package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"testing"
	"time"
)

func TestFn_product_1(t *testing.T) {
	// to produce messages
	topic := "my-topic1"
	partition := 0
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	conn.Close()
}

func TestFn_product_2(t *testing.T) {
	// get kafka writer using environment variables.
	kafkaURL := "localhost:9092" //os.Getenv("kafkaURL")
	topic := "my-topic2"         //os.Getenv("topic")
	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	log.Println("start producing ... !!")
	//
	for i := 0; i < 3; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(fmt.Sprintf("v-%v",uuid.New())),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("--::end")
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}