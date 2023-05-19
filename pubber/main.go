package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

var topicName = "rock-songs"
var (
	endpoint  = os.Getenv("PUBSUB_EMULATOR_HOST")
	projectID = os.Getenv("PUBSUB_PROJECT_ID")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("\x1b[34m", "initializing client...", "\x1b[0m")
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()
	fmt.Println("\x1b[34m", "checking topic existence...", "\x1b[0m")
	topic := client.Topic(topicName)
	exists, err := topic.Exists(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !exists {
		fmt.Println("\x1b[32m", "Creating topic: ", topicName, "\x1b[0m")
		t, err := client.CreateTopic(ctx, topicName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("\x1b[34m topic created \n\ttype: %T\n\tvalue:\n\t\t%+v\x1b[0m\n", t, t)
	}
	fmt.Println("\x1b[34m", "topic exists", "\x1b[0m")

	msgs := []string{
		"Seisun complex",
		"Guitar, Solitude and the Blue Planet",
		"If I can be a constellation",
		"stop",
	}
	for _, m := range msgs {
		topic.Publish(ctx, &pubsub.Message{
			Data: []byte(m),
		})
		time.Sleep(time.Second)
	}
}
