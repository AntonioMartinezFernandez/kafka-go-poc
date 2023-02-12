package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	custom_types "github.com/AntonioMartinezFernandez/kafka-go-poc/producer/custom-types"
	kafka_client "github.com/AntonioMartinezFernandez/kafka-go-poc/producer/kafka-client"
	"github.com/Shopify/sarama"
)

func main() {
	events := []custom_types.Event{
		{
			ID:        0,
			Device:    "device01",
			Timestamp: "2023-02-11T11:52:23Z",
			Payload:   custom_types.Payload{Message: "{'foo': 'bar'}"},
		},
		{
			ID:        1,
			Device:    "device02",
			Timestamp: "2023-02-11T11:52:23Z",
			Payload:   custom_types.Payload{Message: "{'foo': 'bar2'}"},
		},
	}

	brokersUrl := []string{"localhost:9092", "localhost:9094", "localhost:9096"}
	producer, err := kafka_client.GetNewProducer(brokersUrl)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	defer producer.Close()

	eventTopic := "kafka-go-poc"

	for t := range time.NewTicker(500 * time.Millisecond).C {
		fmt.Printf("executing... %s", t)
		run(events, producer, eventTopic)
	}

}

func run(events []custom_types.Event, producer sarama.SyncProducer, eventTopic string) {
	for _, event := range events {
		marshalledEvent, _ := json.Marshal(event)
		kafka_client.Publish(producer, eventTopic, marshalledEvent)
	}
}
