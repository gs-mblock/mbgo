package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	kafka "github.com/segmentio/kafka-go"
)

func producerHandler(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("body %+v\n",string(body))
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("address-%s", req.RemoteAddr)),
			Value: body,
		}
		err = kafkaWriter.WriteMessages(req.Context(), msg)

		if err != nil {
			wrt.Write([]byte(err.Error()))
			log.Fatalln(err)
		}
	})
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func main() {
	fmt.Println("start")
	// get kafka writer using environment variables.
	kafkaURL := "localhost:9092" //os.Getenv("kafkaURL")
	topic := "my-topic3"         //os.Getenv("topic")
	kafkaWriter := getKafkaWriter(kafkaURL, topic)

	defer kafkaWriter.Close()

	// Add handle func for producer.: post data {json}
	http.HandleFunc("/", producerHandler(kafkaWriter))

	// Run the web server.
	fmt.Println("start producer-api ... !!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// go run lib/mq/kafka/demo2/produce/produce_api.go
