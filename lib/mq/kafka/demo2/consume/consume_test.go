package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func TestFn_consume(t *testing.T) {
	// to consume messages
	topic := "my-topic2"
	partition := 0
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println("--get::",string(b))
		//fmt.Println(string(b))
	}
	batch.Close()
	conn.Close()
	fmt.Println("--get::end")
}
