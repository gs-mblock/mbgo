package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/segmentio/kafka-go"
	"log"
)

func getMongoCollection(mongoURL, dbName, collectionName string) *mongo.Collection {
	client, err := mongo.Connect(context.Background(), mongoURL)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB ... !!")

	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return collection
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func main() {

	// get Mongo db Collection using environment variables.
	mongoURL := "mongodb://admin:1234qwer@localhost:27017" //os.Getenv("mongoURL")
	dbName := "example_db"//os.Getenv("dbName")
	collectionName := "example_coll"//os.Getenv("collectionName")
	collection := getMongoCollection(mongoURL, dbName, collectionName)

	// get kafka reader using environment variables.
	kafkaURL := "localhost:9092" // os.Getenv("kafkaURL")
	topic :="my-topic3"// os.Getenv("topic")
	groupID := "groupMG"//os.Getenv("groupID")
	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start consuming ... !!")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		insertResult, err := collection.InsertOne(context.Background(), msg)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}
}

// go run lib/mq/kafka/demo2/consume-mongo/kafka-mongo-db.go