package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
)

var topicName = "rock-songs"
var subId = "first-fan"
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
		topic, err = client.CreateTopic(ctx, topicName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("\x1b[34m topic created \n\ttype: %T\n\tvalue:\n\t\t%+v\x1b[0m\n", topic, topic)
	}
	sub := client.Subscription(subId)
	if ok, err := sub.Exists(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else if !ok {
		fmt.Println("\x1b[32m", "Creating subscription...", "\x1b[0m")
		sub, err = client.CreateSubscription(ctx, subId, pubsub.SubscriptionConfig{Topic: topic})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("\x1b[33m", "Subscription exists", "\x1b[0m")
	}
	fmt.Println("\x1b[34m", "Waiting for message...", "\x1b[0m")
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		m.Ack()
		d := m.Data
		fmt.Println(string(d))
		if string(d) == "stop" {
			cancel()
		}
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\x1b[32m", "bye bye", "\x1b[0m")
}
