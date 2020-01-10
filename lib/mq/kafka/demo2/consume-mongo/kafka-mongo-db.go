package main

import (
	"context"
	"fmt"
	"github.com/gs-mblock/mbgo/lib/utils"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
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
	collectionName := "example_coll2"//os.Getenv("collectionName")
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
		db := new(MqMessageModel)
		db.Offset = msg.Offset
		db.Content = msg.Value
		timeNow := time.Now().Unix()
		db.CreatedTime = timeNow
		db.ModifiedTime =timeNow
		db.ID = utils.BytesToInt64(msg.Key) //binary.BigEndian.Uint64( msg.Key)
		db.Topic = msg.Topic
		//insertResult, err := collection.InsertOne(context.Background(), msg)
		insertResult, err := collection.InsertOne(context.Background(), db)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID,db.ID)
	}
}

type MqMessageModel struct {
	ID  int64 `json:"id" `
	CreatedTime  int64 `json:"createdTime" `
	ModifiedTime int64 `json:"modifiedTime" `
	Topic string `json:"topic" `
	Offset int64 `json:"offset" `
	Content []byte `json:"content" `
}

func CheckData(checkID  int64, s *mongo.Collection)  {
	v:= s.FindOne(context.Background(), bson.D{{"id", 54116309914382037}}, )
	log.Printf("%+v\n",v)
}

// go run lib/mq/kafka/demo2/consume-mongo/kafka-mongo-db.go