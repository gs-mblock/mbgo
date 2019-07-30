package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/mgo.v2"
)

const (
	hosts      = "localhost:27017"
	database   = "testDB"
	username   = "admin"
	password   = "1234qwer"
	collection = "jobskafka"
	kafkaServer = "localhost:32776"
	kafkaTopic = "topic-test"
)
// Job xx
type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Salary      string `json:"salary"`
}

// MongoStore xx
type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func main() {
	println("start server ")
	//Create MongoDB session
	session := initialiseMongo()
	mongoStore.session = session
	receiveFromKafka()
}

func initialiseMongo() (session *mgo.Session) {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		println(err)
	}
	return
}

func receiveFromKafka() {
	fmt.Println("Start receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{kafkaTopic}, nil)
	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			job := string(msg.Value)
			saveJobToMongo(job)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
	c.Close()
}

func saveJobToMongo(jobString string) {
	fmt.Println("Save to MongoDB")
	col := mongoStore.session.DB(database).C(collection)
	//Save data into Job struct
	var _job Job
	b := []byte(jobString)
	err := json.Unmarshal(b, &_job)
	if err != nil {
		panic(err)
	}
	//Insert job into MongoDB
	errMongo := col.Insert(_job)
	if errMongo != nil {
		panic(errMongo)
	}
	fmt.Printf("Saved to MongoDB : %s", jobString)
}